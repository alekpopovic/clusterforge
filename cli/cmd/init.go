package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
)

var initCmd = &cobra.Command{
	Use:   "init <env>",
	Short: "Run Terraform/OpenTofu init for an environment",
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
		return cfterraform.NewRunner(binary, env.Path, opts.Verbose).Init(cmd.Context())
	},
}
