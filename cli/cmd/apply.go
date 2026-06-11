package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/policy"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
)

var applyPlanFile string
var applyConfirmProd bool

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
		return cfterraform.NewRunner(binary, env.Path, opts.Verbose).Apply(cmd.Context(), applyPlanFile, nil)
	},
}

func init() {
	applyCmd.Flags().StringVar(&applyPlanFile, "plan-file", "", "Existing plan file to apply")
	applyCmd.Flags().BoolVar(&applyConfirmProd, "confirm-prod", false, "Explicitly confirm a production apply")
}
