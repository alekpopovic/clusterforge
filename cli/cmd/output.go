package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
)

var outputStack string
var outputJSON bool

var outputCmd = &cobra.Command{
	Use:   "output <env>",
	Short: "Run Terraform/OpenTofu output for an environment or stack",
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
		paths, err := resolveStackPaths(env, outputStack)
		if err != nil {
			return err
		}
		binary, err := engineBinary(cfg)
		if err != nil {
			return err
		}
		for _, path := range paths {
			if label := stackLabel(path, len(paths)); label != "" {
				fmt.Fprintln(cmd.OutOrStdout(), label)
			}
			if err := cfterraform.NewRunner(binary, path, opts.Verbose).Output(cmd.Context(), outputJSON); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	outputCmd.Flags().StringVar(&outputStack, "stack", "", "Stack to read outputs from for stacked environments")
	outputCmd.Flags().BoolVar(&outputJSON, "json", false, "Print outputs as JSON")
}
