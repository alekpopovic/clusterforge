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

		outFile := planOut
		if planRiskSummary && outFile == "" {
			outFile = filepath.Join(".cf", "plans", args[0]+".tfplan")
		}
		if outFile != "" {
			if err := os.MkdirAll(filepath.Join(env.Path, filepath.Dir(outFile)), 0o755); err != nil {
				return fmt.Errorf("create plan directory: %w", err)
			}
		}

		runner := cfterraform.NewRunner(binary, env.Path, opts.Verbose)
		if err := runner.Plan(cmd.Context(), outFile, nil); err != nil {
			return err
		}
		if !planRiskSummary {
			return nil
		}
		data, err := runner.ShowPlanJSON(cmd.Context(), outFile)
		if err != nil {
			evaluation, evalErr := policy.EvaluatePlanParseError(args[0], err)
			for _, warning := range evaluation.Warnings {
				printer.Warn(warning)
			}
			if evalErr != nil {
				return evalErr
			}
			return nil
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
		return nil
	},
}

func init() {
	planCmd.Flags().StringVar(&planOut, "out", "", "Write a plan file")
	planCmd.Flags().BoolVar(&planRiskSummary, "risk-summary", false, "Show a risk summary from Terraform plan JSON")
}
