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
var policyPack string

var policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Check ClusterForge safety policies",
}

var policyListCmd = &cobra.Command{
	Use:   "list",
	Short: "List built-in policy packs",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, pack := range []string{"baseline", "production", "kubernetes-security", "aws-security"} {
			fmt.Fprintln(cmd.OutOrStdout(), pack)
		}
		return nil
	},
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
			Pack:        defaultString(policyPack, "baseline"),
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
		response.Messages = append(response.Messages, policyPackMessages(response.Pack)...)
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
	policyCheckCmd.Flags().StringVar(&policyPack, "pack", "baseline", "Policy pack to check: baseline, production, kubernetes-security, aws-security")
	policyCheckCmd.Flags().BoolVar(&policyCheckJSON, "json", false, "Print policy check as JSON")
	policyCmd.AddCommand(policyListCmd)
	policyCmd.AddCommand(policyCheckCmd)
}

func policyPackMessages(pack string) []string {
	switch pack {
	case "baseline":
		return []string{"baseline pack: require reviewed plans and safe destructive-operation gates"}
	case "production":
		return []string{"production pack: remote backend, provider pinning, tagged module sources, and ingress approvals are required by convention"}
	case "kubernetes-security":
		return []string{"kubernetes-security pack: pod security labels and network policies are recommended; privileged workloads require review"}
	case "aws-security":
		return []string{"aws-security pack: public S3, unrestricted security groups, unencrypted state, and IAM wildcards require review"}
	default:
		return []string{fmt.Sprintf("unknown pack %q; only built-in baseline checks were evaluated", pack)}
	}
}

type policyCheckResponse struct {
	Environment string         `json:"environment"`
	Pack        string         `json:"pack"`
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
