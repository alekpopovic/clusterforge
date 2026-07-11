package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/config"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

var envCreateCloud string
var envCreateRegion string
var envCreateOrchestrator string
var envListJSON bool
var envCreateWizard bool

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage ClusterForge environments",
}

var envCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create an environment entry and live directory",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		prompts := newPromptSession(cmd)
		name, err := requireValueWithPrompt(optionalArg(args, 0), "environment name", prompts)
		if err != nil {
			return err
		}

		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		if _, exists := cfg.Environments[name]; exists {
			return fmt.Errorf("environment %q already exists", name)
		}

		cloud := defaultString(envCreateCloud, cfg.Defaults.Cloud)
		region := defaultString(envCreateRegion, cfg.Defaults.Region)
		orchestrator := defaultString(envCreateOrchestrator, cfg.Defaults.Orchestrator)
		path := fmt.Sprintf("live/%s/%s-%s", name, cloud, orchestrator)
		if !opts.NonInteractive {
			if !cmd.Flags().Changed("cloud") {
				cloud, err = prompts.String("cloud", cloud)
				if err != nil {
					return err
				}
			}
			if !cmd.Flags().Changed("orchestrator") {
				orchestrator, err = prompts.String("orchestrator", orchestrator)
				if err != nil {
					return err
				}
			}
			if !cmd.Flags().Changed("region") {
				region, err = prompts.String("region", region)
				if err != nil {
					return err
				}
			}
			path = fmt.Sprintf("live/%s/%s-%s", name, cloud, orchestrator)
			path, err = prompts.String("environment path", path)
			if err != nil {
				return err
			}
		}
		env := config.Environment{
			Cloud:        cloud,
			Region:       region,
			Orchestrator: orchestrator,
			Path:         path,
			Layout:       "simple",
		}
		if warning := dockerTargetWarning(env.Orchestrator); warning != "" {
			printer.Warn(warning)
		}

		printer.Info(fmt.Sprintf("environment: %s", name))
		printer.Info(fmt.Sprintf("cloud: %s", env.Cloud))
		printer.Info(fmt.Sprintf("orchestrator: %s", env.Orchestrator))
		printer.Info(fmt.Sprintf("region: %s", env.Region))
		printer.Info(fmt.Sprintf("path: %s", env.Path))

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

func dockerTargetWarning(orchestrator string) string {
	if orchestrator == "docker" || orchestrator == "swarm" {
		return "Docker and Docker Swarm targets are experimental; review security, networking, secrets, and host lifecycle limitations"
	}
	return ""
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

		environments := make([]envListItem, 0, len(names))
		for _, name := range names {
			env := cfg.Environments[name]
			environments = append(environments, envListItem{
				Name:         name,
				Cloud:        env.Cloud,
				Orchestrator: env.Orchestrator,
				Region:       env.Region,
				Path:         env.Path,
			})
		}
		if envListJSON {
			return ui.WriteJSON(cmd.OutOrStdout(), envListResponse{Environments: environments})
		}
		for _, env := range environments {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%s\t%s\n", env.Name, env.Cloud, env.Region, env.Orchestrator, env.Path)
		}
		return nil
	},
}

type envListResponse struct {
	Environments []envListItem `json:"environments"`
}

type envListItem struct {
	Name         string `json:"name"`
	Cloud        string `json:"cloud"`
	Orchestrator string `json:"orchestrator"`
	Region       string `json:"region"`
	Path         string `json:"path"`
}

func init() {
	envCreateCmd.Flags().StringVar(&envCreateCloud, "cloud", "", "Cloud target for the environment")
	envCreateCmd.Flags().StringVar(&envCreateRegion, "region", "", "Cloud region for the environment")
	envCreateCmd.Flags().StringVar(&envCreateOrchestrator, "orchestrator", "", "Orchestrator target for the environment")
	envCreateCmd.Flags().BoolVar(&envCreateWizard, "wizard", false, "Prompt for environment settings (default interactively)")
	envListCmd.Flags().BoolVar(&envListJSON, "json", false, "Print environments as JSON")

	envCmd.AddCommand(envCreateCmd)
	envCmd.AddCommand(envListCmd)
}

func defaultString(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
