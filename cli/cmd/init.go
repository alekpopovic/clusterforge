package cmd

import (
	"fmt"
	"os"

	cfterraform "github.com/alekpopovic/clusterforge/cli/internal/terraform"
	"github.com/spf13/cobra"
)

var initStack string

var initCmd = &cobra.Command{
	Use:   "init [env]",
	Short: "Run Terraform/OpenTofu init for an environment",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return runScaffoldWizard(cmd, args)
		}
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
		paths, err := resolveStackPaths(env, initStack)
		if err != nil {
			return err
		}
		for _, path := range paths {
			if label := stackLabel(path, len(paths)); label != "" {
				fmt.Fprintln(os.Stdout, label)
			}
			if err := cfterraform.NewRunner(binary, path, opts.Verbose).Init(cmd.Context()); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	initCmd.Flags().StringVar(&initStack, "stack", "", "Stack to initialize for stacked environments")
	initCmd.Flags().BoolVar(&wizardDefaults, "defaults", false, "Use safe scaffolding defaults when no environment is supplied")
}
