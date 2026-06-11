package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/config"
)

var projectForce bool

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage ClusterForge projects",
}

var projectInitCmd = &cobra.Command{
	Use:   "init <name>",
	Short: "Initialize a ClusterForge project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		cfgPath := opts.ConfigPath
		if !projectForce {
			if _, err := os.Stat(cfgPath); err == nil {
				return fmt.Errorf("%s already exists; use --force to overwrite it", cfgPath)
			}
		}

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
	projectCmd.AddCommand(projectInitCmd)
}
