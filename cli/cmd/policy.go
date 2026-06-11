package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/policy"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
	"github.com/textracta/clusterforge/cli/internal/terraform/planjson"
)

var policyPlanFile string

var policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Check ClusterForge safety policies",
}

var policyCheckCmd = &cobra.Command{
	Use:   "check <env>",
	Short: "Check safety policies for an environment",
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
			fmt.Fprintln(cmd.OutOrStdout(), "Policy: production apply requires --plan-file")
			fmt.Fprintln(cmd.OutOrStdout(), "Policy: production destroy requires --allow-destroy")
		} else {
			fmt.Fprintln(cmd.OutOrStdout(), "Policy: non-production changes should still be reviewed before apply")
		}
		if policyPlanFile == "" {
			return nil
		}

		binary, err := engineBinary(cfg)
		if err != nil {
			return err
		}
		runner := cfterraform.NewRunner(binary, env.Path, opts.Verbose)
		data, err := runner.ShowPlanJSON(cmd.Context(), policyPlanFile)
		if err != nil {
			evaluation, evalErr := policy.EvaluatePlanParseError(envName, err)
			planjson.Print(cmd.OutOrStdout(), envName, planjson.Summary{}, evaluation.Risk, evaluation.Policy)
			for _, warning := range evaluation.Warnings {
				printer.Warn(warning)
			}
			return evalErr
		}
		summary, err := planjson.Parse(data)
		if err != nil {
			evaluation, evalErr := policy.EvaluatePlanParseError(envName, err)
			planjson.Print(cmd.OutOrStdout(), envName, planjson.Summary{}, evaluation.Risk, evaluation.Policy)
			for _, warning := range evaluation.Warnings {
				printer.Warn(warning)
			}
			return evalErr
		}
		evaluation := policy.EvaluatePlan(envName, summary)
		planjson.Print(cmd.OutOrStdout(), envName, summary, evaluation.Risk, evaluation.Policy)
		for _, warning := range evaluation.Warnings {
			printer.Warn(warning)
		}
		return nil
	},
}

func init() {
	policyCheckCmd.Flags().StringVar(&policyPlanFile, "plan-file", "", "Existing plan file to summarize")
	policyCmd.AddCommand(policyCheckCmd)
}
