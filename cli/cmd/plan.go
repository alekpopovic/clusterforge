package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
)

var planOut string

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

		tfArgs := []string{"plan"}
		if planOut != "" {
			tfArgs = append(tfArgs, "-out", planOut)
		}
		return cfterraform.NewRunner(binary, env.Path, opts.Verbose).Run(cmd.Context(), tfArgs...)
	},
}

func init() {
	planCmd.Flags().StringVar(&planOut, "out", "", "Write a plan file")
}
