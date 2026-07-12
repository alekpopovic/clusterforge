package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"github.com/spf13/cobra"
)

var projectForce bool
var projectWizard bool

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage ClusterForge projects",
}

var projectInitCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Initialize a ClusterForge project",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if projectWizard {
			return runScaffoldWizard(cmd, args)
		}
		name, err := requireValueWithPrompt(optionalArg(args, 0), "project name", newPromptSession(cmd))
		if err != nil {
			return err
		}
		cfgPath := opts.ConfigPath
		if !projectForce {
			if _, err := os.Stat(cfgPath); err == nil {
				return fmt.Errorf("%s already exists; use --force to overwrite it", cfgPath)
			}
		}

		printer.Info(fmt.Sprintf("project: %s", name))
		printer.Info(fmt.Sprintf("config: %s", filepath.Clean(cfgPath)))
		printer.Info("directories: apps, live, .cf")

		for _, dir := range []string{"apps", "live", ".cf"} {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return fmt.Errorf("create %s: %w", dir, err)
			}
		}

		cfg := config.DefaultConfig(name)
		if err := config.Write(cfgPath, cfg, projectForce); err != nil {
			return err
		}

		printer.Success(fmt.Sprintf("initialized ClusterForge project %q", name))
		printer.Info(fmt.Sprintf("config: %s", filepath.Clean(cfgPath)))
		return nil
	},
}

func init() {
	projectInitCmd.Flags().BoolVar(&projectForce, "force", false, "Overwrite existing project config")
	projectInitCmd.Flags().BoolVar(&projectWizard, "wizard", false, "Use the full project scaffolding wizard")
	projectCmd.AddCommand(projectInitCmd)
}
