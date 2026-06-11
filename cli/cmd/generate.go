package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/generator"
)

var generateCmd = &cobra.Command{
	Use:   "generate <env>",
	Short: "Generate readable Terraform files for an environment",
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
		if err := generator.Generate(args[0], env); err != nil {
			return err
		}
		printer.Success(fmt.Sprintf("generated environment files in %s", env.Path))
		return nil
	},
}
