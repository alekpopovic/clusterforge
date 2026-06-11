package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/policy"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
)

var allowDestroy bool
var destroyConfirmProd bool

var destroyCmd = &cobra.Command{
	Use:   "destroy <env>",
	Short: "Run Terraform/OpenTofu destroy for an environment",
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
		if err := policy.CheckDestroy(policy.Operation{
			Environment:  envName,
			AllowDestroy: allowDestroy,
			ConfirmProd:  destroyConfirmProd,
			Policies:     cfg.Policies,
		}); err != nil {
			return err
		}

		printer.Warn("destroy is destructive; Terraform/OpenTofu will ask for confirmation")
		binary, err := engineBinary(cfg)
		if err != nil {
			return err
		}
		return cfterraform.NewRunner(binary, env.Path, opts.Verbose).Run(cmd.Context(), "destroy")
	},
}

func init() {
	destroyCmd.Flags().BoolVar(&allowDestroy, "allow-destroy", false, "Allow destroy in protected environments")
	destroyCmd.Flags().BoolVar(&destroyConfirmProd, "confirm-prod", false, "Explicitly confirm a production destroy")
}
