package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/policy"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
	"github.com/textracta/clusterforge/cli/internal/terraform/planjson"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

var policyPlanFile string
var policyCheckJSON bool

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
		response := policyCheckResponse{
			Environment: envName,
			Policies: policySettings{
				RequirePlanFileForApply: cfg.Policies.RequirePlanFileForApply,
				BlockDestroyInProd:      cfg.Policies.BlockDestroyInProd,
			},
		}
		if policy.IsProd(envName) {
			response.Messages = append(response.Messages,
				"production apply requires --plan-file",
				"production destroy requires --allow-destroy",
			)
		} else {
			response.Messages = append(response.Messages, "non-production changes should still be reviewed before apply")
		}
		if policyPlanFile == "" {
			if policyCheckJSON {
				return ui.WriteJSON(cmd.OutOrStdout(), response)
			}
			for _, message := range response.Messages {
				fmt.Fprintf(cmd.OutOrStdout(), "Policy: %s\n", message)
			}
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
			response.Risk = evaluation.Risk
			response.Policy = evaluation.Policy
			response.Warnings = evaluation.Warnings
			response.Summary = summaryResponse(planjson.Summary{})
			if policyCheckJSON {
				if writeErr := ui.WriteJSON(cmd.OutOrStdout(), response); writeErr != nil {
					return writeErr
				}
			} else {
				planjson.Print(cmd.OutOrStdout(), envName, planjson.Summary{}, evaluation.Risk, evaluation.Policy)
			}
			for _, warning := range evaluation.Warnings {
				printer.Warn(warning)
			}
			return evalErr
		}
		summary, err := planjson.Parse(data)
		if err != nil {
			evaluation, evalErr := policy.EvaluatePlanParseError(envName, err)
			response.Risk = evaluation.Risk
			response.Policy = evaluation.Policy
			response.Warnings = evaluation.Warnings
			response.Summary = summaryResponse(planjson.Summary{})
			if policyCheckJSON {
				if writeErr := ui.WriteJSON(cmd.OutOrStdout(), response); writeErr != nil {
					return writeErr
				}
			} else {
				planjson.Print(cmd.OutOrStdout(), envName, planjson.Summary{}, evaluation.Risk, evaluation.Policy)
			}
			for _, warning := range evaluation.Warnings {
				printer.Warn(warning)
			}
			return evalErr
		}
		evaluation := policy.EvaluatePlan(envName, summary)
		response.Risk = evaluation.Risk
		response.Policy = evaluation.Policy
		response.Warnings = evaluation.Warnings
		response.Summary = summaryResponse(summary)
		if policyCheckJSON {
			return ui.WriteJSON(cmd.OutOrStdout(), response)
		}
		for _, message := range response.Messages {
			fmt.Fprintf(cmd.OutOrStdout(), "Policy: %s\n", message)
		}
		planjson.Print(cmd.OutOrStdout(), envName, summary, evaluation.Risk, evaluation.Policy)
		for _, warning := range evaluation.Warnings {
			printer.Warn(warning)
		}
		return nil
	},
}

func init() {
	policyCheckCmd.Flags().StringVar(&policyPlanFile, "plan-file", "", "Existing plan file to summarize")
	policyCheckCmd.Flags().BoolVar(&policyCheckJSON, "json", false, "Print policy check as JSON")
	policyCmd.AddCommand(policyCheckCmd)
}

type policyCheckResponse struct {
	Environment string         `json:"environment"`
	Policies    policySettings `json:"policies"`
	Messages    []string       `json:"messages"`
	Risk        string         `json:"risk,omitempty"`
	Policy      string         `json:"policy,omitempty"`
	Warnings    []string       `json:"warnings,omitempty"`
	Summary     *summaryJSON   `json:"summary,omitempty"`
}

type policySettings struct {
	RequirePlanFileForApply bool `json:"require_plan_file_for_apply"`
	BlockDestroyInProd      bool `json:"block_destroy_in_prod"`
}
