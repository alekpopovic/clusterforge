package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/config"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

const version = "dev"

type rootOptions struct {
	ConfigPath string
	Engine     string
	Verbose    bool
}

var opts = rootOptions{
	ConfigPath: config.DefaultPath,
}

var printer = ui.NewPrinter(os.Stdout, os.Stderr)

var rootCmd = &cobra.Command{
	Use:   "cf",
	Short: "ClusterForge CLI",
	Long:  "ClusterForge CLI is a wrapper and generator. Terraform and OpenTofu files stay visible and editable.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&opts.ConfigPath, "config", config.DefaultPath, "Path to the ClusterForge config file")
	rootCmd.PersistentFlags().StringVar(&opts.Engine, "engine", "", "Override IaC engine: terraform or tofu")
	rootCmd.PersistentFlags().BoolVar(&opts.Verbose, "verbose", false, "Enable verbose output")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(appCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(planCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(destroyCmd)
	rootCmd.AddCommand(doctorCmd)
}

func loadConfig() (*config.Config, error) {
	return config.Load(opts.ConfigPath)
}

func engineBinary(cfg *config.Config) (string, error) {
	engine := opts.Engine
	if engine == "" {
		engine = cfg.Project.DefaultEngine
	}
	if engine == "" {
		engine = "terraform"
	}
	return cfg.EngineBinary(engine)
}
