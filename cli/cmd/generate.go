package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/generator"
)

var generateForce bool
var generateDryRun bool
var generateCloud string
var generateOrchestrator string

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
		result, err := generator.Generate(args[0], env, generator.Options{
			Force:        generateForce,
			DryRun:       generateDryRun,
			Cloud:        generateCloud,
			Orchestrator: generateOrchestrator,
			Project:      cfg.Project.Name,
			RootDir:      ".",
			Stdout:       os.Stdout,
		})
		if err != nil {
			return err
		}
		if generateDryRun {
			printer.Success(fmt.Sprintf("dry run complete for %s", result.Target))
			return nil
		}
		printer.Success(fmt.Sprintf("generated %s environment files in %s", result.Target, env.Path))
		return nil
	},
}

func init() {
	generateCmd.Flags().BoolVar(&generateForce, "force", false, "Overwrite generated environment files")
	generateCmd.Flags().BoolVar(&generateDryRun, "dry-run", false, "Print files that would be generated without writing them")
	generateCmd.Flags().StringVar(&generateCloud, "cloud", "", "Override environment cloud for generation")
	generateCmd.Flags().StringVar(&generateOrchestrator, "orchestrator", "", "Override environment orchestrator for generation")
}
