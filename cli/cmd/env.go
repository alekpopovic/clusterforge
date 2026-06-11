package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/config"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage ClusterForge environments",
}

var envCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create an environment entry and live directory",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		name := args[0]
		if _, exists := cfg.Environments[name]; exists {
			return fmt.Errorf("environment %q already exists", name)
		}

		env := config.Environment{
			Cloud:        cfg.Defaults.Cloud,
			Region:       cfg.Defaults.Region,
			Orchestrator: cfg.Defaults.Orchestrator,
			Path:         fmt.Sprintf("live/%s/%s-%s", name, cfg.Defaults.Cloud, cfg.Defaults.Orchestrator),
		}
		cfg.Environments[name] = env

		if err := os.MkdirAll(env.Path, 0o755); err != nil {
			return fmt.Errorf("create live directory %s: %w", env.Path, err)
		}
		if err := config.Write(opts.ConfigPath, cfg, true); err != nil {
			return err
		}

		printer.Success(fmt.Sprintf("created environment %q at %s", name, env.Path))
		return nil
	},
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured environments",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		names := make([]string, 0, len(cfg.Environments))
		for name := range cfg.Environments {
			names = append(names, name)
		}
		sort.Strings(names)

		for _, name := range names {
			env := cfg.Environments[name]
			fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%s\t%s\n", name, env.Cloud, env.Region, env.Orchestrator, env.Path)
		}
		return nil
	},
}

func init() {
	envCmd.AddCommand(envCreateCmd)
	envCmd.AddCommand(envListCmd)
}
