package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/policy"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
	"github.com/textracta/clusterforge/cli/internal/terraform/planjson"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

var driftStack string
var driftJSON bool
var driftSummary bool

var driftCmd = &cobra.Command{
	Use:   "drift",
	Short: "Detect Terraform/OpenTofu drift without applying changes",
}

var driftCheckCmd = &cobra.Command{
	Use:   "check <env>",
	Short: "Run a detailed-exitcode plan to detect drift",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		envName := args[0]
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		env, ok := cfg.Environments[envName]
		if !ok {
			return fmt.Errorf("environment %q not found", envName)
		}
		if policy.IsProd(envName) {
			printer.Warn("production drift must be reviewed; ClusterForge will not remediate automatically")
		}
		binary, err := engineBinary(cfg)
		if err != nil {
			return err
		}
		paths, err := resolveStackPaths(env, driftStack)
		if err != nil {
			return err
		}
		response := driftResponse{Environment: envName}
		highestExit := 0
		for index, path := range paths {
			planFile := filepath.Join(".cf", "plans", fmt.Sprintf("%s-drift-%d.tfplan", envName, index+1))
			if len(paths) == 1 {
				planFile = filepath.Join(".cf", "plans", fmt.Sprintf("%s-drift.tfplan", envName))
			}
			if err := os.MkdirAll(filepath.Join(path, filepath.Dir(planFile)), 0o755); err != nil {
				return fmt.Errorf("create drift plan directory: %w", err)
			}
			runner := cfterraform.NewRunner(binary, path, opts.Verbose)
			code, err := runner.PlanDetailedExitCode(cmd.Context(), planFile, nil)
			if err != nil {
				return err
			}
			if code > highestExit {
				highestExit = code
			}
			item := driftStackResult{Path: path, PlanFile: planFile, ExitCode: code, Drift: code == 2}
			if driftJSON || driftSummary {
				data, err := runner.ShowPlanJSON(cmd.Context(), planFile)
				if err == nil {
					summary, parseErr := planjson.Parse(data)
					if parseErr == nil {
						evaluation := policy.EvaluatePlan(envName, summary)
						item.Risk = evaluation.Risk
						item.Policy = evaluation.Policy
						item.Summary = summaryResponse(summary)
					}
				}
			}
			response.Stacks = append(response.Stacks, item)
		}
		response.Drift = highestExit == 2
		if driftJSON {
			if err := ui.WriteJSON(cmd.OutOrStdout(), response); err != nil {
				return err
			}
		} else {
			if response.Drift {
				fmt.Fprintf(cmd.OutOrStdout(), "drift detected for %s\n", envName)
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "no drift detected for %s\n", envName)
			}
			for _, stack := range response.Stacks {
				fmt.Fprintf(cmd.OutOrStdout(), "%s: exit_code=%d plan=%s\n", stack.Path, stack.ExitCode, stack.PlanFile)
			}
		}
		if highestExit == 2 {
			os.Exit(2)
		}
		return nil
	},
}

type driftResponse struct {
	Environment string             `json:"environment"`
	Drift       bool               `json:"drift"`
	Stacks      []driftStackResult `json:"stacks"`
}

type driftStackResult struct {
	Path     string       `json:"path"`
	PlanFile string       `json:"plan_file"`
	ExitCode int          `json:"exit_code"`
	Drift    bool         `json:"drift"`
	Risk     string       `json:"risk,omitempty"`
	Policy   string       `json:"policy,omitempty"`
	Summary  *summaryJSON `json:"summary,omitempty"`
}

func init() {
	driftCheckCmd.Flags().StringVar(&driftStack, "stack", "", "Stack to check for stacked environments")
	driftCheckCmd.Flags().BoolVar(&driftJSON, "json", false, "Print drift result as JSON")
	driftCheckCmd.Flags().BoolVar(&driftSummary, "summary", false, "Parse and print drift plan summary")
	driftCmd.AddCommand(driftCheckCmd)
}
