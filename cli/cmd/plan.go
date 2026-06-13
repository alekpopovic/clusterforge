package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/policy"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
	"github.com/textracta/clusterforge/cli/internal/terraform/planjson"
)

var planOut string
var planRiskSummary bool
var planStack string

var planCmd = &cobra.Command{
	Use:   "plan <env>",
	Short: "Run Terraform/OpenTofu plan for an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		env, ok := cfg.Environments[args[0]]
		if !ok {
			return fmt.Errorf("environment %q not found", args[0])
		}
		binary, err := engineBinary(cfg)
		if err != nil {
			return err
		}

		paths, err := resolveStackPaths(env, planStack)
		if err != nil {
			return err
		}
		if planOut != "" && len(paths) > 1 {
			return fmt.Errorf("--out requires --stack when planning multiple stacks")
		}
		outFile := planOut
		if planRiskSummary && outFile == "" {
			outFile = filepath.Join(".cf", "plans", args[0]+".tfplan")
		}
		for index, path := range paths {
			currentOut := outFile
			if planRiskSummary && planOut == "" && len(paths) > 1 {
				currentOut = filepath.Join(".cf", "plans", fmt.Sprintf("%s-%d.tfplan", args[0], index+1))
			}
			if currentOut != "" {
				if err := os.MkdirAll(filepath.Join(path, filepath.Dir(currentOut)), 0o755); err != nil {
					return fmt.Errorf("create plan directory: %w", err)
				}
			}
			if label := stackLabel(path, len(paths)); label != "" {
				fmt.Fprintln(cmd.OutOrStdout(), label)
			}
			runner := cfterraform.NewRunner(binary, path, opts.Verbose)
			if err := runner.Plan(cmd.Context(), currentOut, nil); err != nil {
				return err
			}
			if !planRiskSummary {
				continue
			}
			data, err := runner.ShowPlanJSON(cmd.Context(), currentOut)
			if err != nil {
				evaluation, evalErr := policy.EvaluatePlanParseError(args[0], err)
				for _, warning := range evaluation.Warnings {
					printer.Warn(warning)
				}
				if evalErr != nil {
					return evalErr
				}
				continue
			}
			summary, err := planjson.Parse(data)
			if err != nil {
				evaluation, evalErr := policy.EvaluatePlanParseError(args[0], err)
				planjson.Print(cmd.OutOrStdout(), args[0], planjson.Summary{}, evaluation.Risk, evaluation.Policy)
				for _, warning := range evaluation.Warnings {
					printer.Warn(warning)
				}
				return evalErr
			}
			evaluation := policy.EvaluatePlan(args[0], summary)
			planjson.Print(cmd.OutOrStdout(), args[0], summary, evaluation.Risk, evaluation.Policy)
			for _, warning := range evaluation.Warnings {
				printer.Warn(warning)
			}
		}
		return nil
	},
}

func init() {
	planCmd.Flags().StringVar(&planOut, "out", "", "Write a plan file")
	planCmd.Flags().BoolVar(&planRiskSummary, "risk-summary", false, "Show a risk summary from Terraform plan JSON")
	planCmd.Flags().StringVar(&planStack, "stack", "", "Stack to plan for stacked environments")
}
