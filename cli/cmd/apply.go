package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/policy"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
	"github.com/textracta/clusterforge/cli/internal/terraform/planjson"
)

var applyPlanFile string
var applyConfirmProd bool
var applyAllowDestroy bool
var applyStack string

var applyCmd = &cobra.Command{
	Use:   "apply <env>",
	Short: "Run Terraform/OpenTofu apply for an environment",
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
		if err := policy.CheckApply(policy.Operation{
			Environment: envName,
			PlanFile:    applyPlanFile,
			ConfirmProd: applyConfirmProd,
			Policies:    cfg.Policies,
		}); err != nil {
			return err
		}

		printer.Warn("apply may change infrastructure; review the plan before continuing")
		binary, err := engineBinary(cfg)
		if err != nil {
			return err
		}
		paths, err := resolveStackPaths(env, applyStack)
		if err != nil {
			return err
		}
		if applyPlanFile != "" && len(paths) > 1 {
			return fmt.Errorf("--plan-file requires --stack when applying multiple stacks")
		}
		runner := cfterraform.NewRunner(binary, paths[0], opts.Verbose)
		if applyPlanFile != "" {
			data, err := runner.ShowPlanJSON(cmd.Context(), applyPlanFile)
			if err != nil {
				evaluation, evalErr := policy.EvaluatePlanParseError(envName, err)
				planjson.Print(cmd.OutOrStdout(), envName, planjson.Summary{}, evaluation.Risk, evaluation.Policy)
				for _, warning := range evaluation.Warnings {
					printer.Warn(warning)
				}
				if evalErr != nil {
					return evalErr
				}
			} else {
				summary, err := planjson.Parse(data)
				if err != nil {
					evaluation, evalErr := policy.EvaluatePlanParseError(envName, err)
					planjson.Print(cmd.OutOrStdout(), envName, planjson.Summary{}, evaluation.Risk, evaluation.Policy)
					for _, warning := range evaluation.Warnings {
						printer.Warn(warning)
					}
					if evalErr != nil {
						return evalErr
					}
				} else {
					evaluation, err := policy.CheckApplyPlan(policy.Operation{
						Environment:  envName,
						AllowDestroy: applyAllowDestroy,
						Policies:     cfg.Policies,
					}, summary)
					planjson.Print(cmd.OutOrStdout(), envName, summary, evaluation.Risk, evaluation.Policy)
					for _, warning := range evaluation.Warnings {
						printer.Warn(warning)
					}
					if err != nil {
						return err
					}
				}
			}
		}
		for _, path := range paths {
			if label := stackLabel(path, len(paths)); label != "" {
				fmt.Fprintln(os.Stdout, label)
			}
			if err := cfterraform.NewRunner(binary, path, opts.Verbose).Apply(cmd.Context(), applyPlanFile, nil); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	applyCmd.Flags().StringVar(&applyPlanFile, "plan-file", "", "Existing plan file to apply")
	applyCmd.Flags().BoolVar(&applyConfirmProd, "confirm-prod", false, "Explicitly confirm a production apply")
	applyCmd.Flags().BoolVar(&applyAllowDestroy, "allow-destroy", false, "Allow applying prod plans that contain delete actions")
	applyCmd.Flags().StringVar(&applyStack, "stack", "", "Stack to apply for stacked environments")
}
