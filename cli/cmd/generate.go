package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/config"
	"github.com/textracta/clusterforge/cli/internal/generator"
)

var generateForce bool
var generateDryRun bool
var generateCloud string
var generateOrchestrator string
var generateLayout string

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
		if generateLayout != "" {
			updated, err := environmentWithLayout(env, generateLayout)
			if err != nil {
				return err
			}
			env = updated
		}
		backend := cfg.BackendFor(args[0])
		result, err := generator.Generate(args[0], env, generator.Options{
			Force:        generateForce,
			DryRun:       generateDryRun,
			Cloud:        generateCloud,
			Orchestrator: generateOrchestrator,
			Layout:       generateLayout,
			Backend:      backend,
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
		if generateLayout != "" {
			cfg.Environments[args[0]] = env
			if err := cfg.Save(opts.ConfigPath); err != nil {
				return err
			}
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
	generateCmd.Flags().StringVar(&generateLayout, "layout", "", "Environment layout to generate: simple or stacked")
}

func environmentWithLayout(env config.Environment, layout string) (config.Environment, error) {
	switch layout {
	case "simple":
		env.Layout = "simple"
		env.Stacks = nil
	case "stacked":
		env.Layout = "stacked"
		if env.Stacks == nil {
			env.Stacks = config.Stacks{}
		}
		for _, stack := range config.StackOrder() {
			current := env.Stacks[stack]
			if current.Path == "" {
				current.Path = filepath.Join(env.Path, stack)
			}
			env.Stacks[stack] = current
		}
	default:
		return env, fmt.Errorf("unsupported layout %q; expected simple or stacked", layout)
	}
	return env, nil
}
