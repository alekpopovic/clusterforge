package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/config"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

type rootOptions struct {
	ConfigPath     string
	Engine         string
	Verbose        bool
	NonInteractive bool
	NoPlugins      bool
	AllowPlugins   bool
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
		var coded interface{ ExitCode() int }
		if errors.As(err, &coded) {
			os.Exit(coded.ExitCode())
		}
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&opts.ConfigPath, "config", config.DefaultPath, "Path to the ClusterForge config file")
	rootCmd.PersistentFlags().StringVar(&opts.Engine, "engine", "", "Override IaC engine: terraform or tofu")
	rootCmd.PersistentFlags().BoolVar(&opts.Verbose, "verbose", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVar(&opts.NonInteractive, "non-interactive", false, "Disable prompts and fail when required values are missing")
	rootCmd.PersistentFlags().BoolVar(&opts.NoPlugins, "no-plugins", false, "Disable plugin discovery and execution")
	rootCmd.PersistentFlags().BoolVar(&opts.AllowPlugins, "allow-plugins", false, "Explicitly allow trusted local plugins in CI")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(appCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(clusterCmd)
	rootCmd.AddCommand(fleetCmd)
	rootCmd.AddCommand(graphCmd)
	rootCmd.AddCommand(auditCmd)
	rootCmd.AddCommand(backendCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(policyCmd)
	rootCmd.AddCommand(driftCmd)
	rootCmd.AddCommand(stateCmd)
	rootCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(promoteCmd)
	rootCmd.AddCommand(costCmd)
	rootCmd.AddCommand(templateCmd)
	rootCmd.AddCommand(moduleCmd)
	rootCmd.AddCommand(localCmd)
	rootCmd.AddCommand(pluginCmd)
	rootCmd.AddCommand(workspaceCmd)
	rootCmd.AddCommand(teamCmd)
	rootCmd.AddCommand(accountCmd)
	rootCmd.AddCommand(regionCmd)
	rootCmd.AddCommand(k8sCmd)
	rootCmd.AddCommand(platformCmd)
	rootCmd.AddCommand(profileCmd)
	rootCmd.AddCommand(tfeCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(planCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(destroyCmd)
	rootCmd.AddCommand(outputCmd)
	rootCmd.AddCommand(doctorCmd)
	installAuditWrappers(rootCmd)
}

type commandExitError struct {
	code    int
	message string
}

func (e commandExitError) Error() string { return e.message }
func (e commandExitError) ExitCode() int { return e.code }

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
