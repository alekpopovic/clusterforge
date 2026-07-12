package cmd

import (
	"fmt"
	"os"

	"github.com/alekpopovic/clusterforge/cli/internal/policy"
	cfterraform "github.com/alekpopovic/clusterforge/cli/internal/terraform"
	"github.com/spf13/cobra"
)

var allowDestroy bool
var destroyConfirmProd bool
var destroyStack string

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
		paths, err := resolveStackPaths(env, destroyStack)
		if err != nil {
			return err
		}
		if destroyStack == "" && env.EffectiveLayout() == "stacked" {
			for i, j := 0, len(paths)-1; i < j; i, j = i+1, j-1 {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
		for _, path := range paths {
			if label := stackLabel(path, len(paths)); label != "" {
				fmt.Fprintln(os.Stdout, label)
			}
			if err := cfterraform.NewRunner(binary, path, opts.Verbose).Destroy(cmd.Context(), nil); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	destroyCmd.Flags().BoolVar(&allowDestroy, "allow-destroy", false, "Allow destroy in protected environments")
	destroyCmd.Flags().BoolVar(&destroyConfirmProd, "confirm-prod", false, "Explicitly confirm a production destroy")
	destroyCmd.Flags().StringVar(&destroyStack, "stack", "", "Stack to destroy for stacked environments")
}
