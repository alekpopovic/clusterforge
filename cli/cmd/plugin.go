package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/plugins"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Manage trusted local executable plugins",
	Long:  "Plugins are trusted local executables. ClusterForge never downloads or automatically runs them.",
}

var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List discovered plugins and their status",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		found, err := discoverPlugins()
		if err != nil {
			return err
		}
		if len(found) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "no plugins discovered")
			return nil
		}
		for _, plugin := range found {
			status := "enabled"
			if plugin.Disabled {
				status = "disabled"
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\n", plugin.Name, status, plugin.Path)
		}
		return nil
	},
}

var pluginDiscoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Show plugin discovery results without executing plugins",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		found, err := discoverPlugins()
		if err != nil {
			return err
		}
		for _, plugin := range found {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\n", plugin.Name, plugin.Path)
		}
		if len(found) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "no plugins discovered")
		}
		return nil
	},
}

var pluginInfoCmd = &cobra.Command{
	Use:   "info <name>",
	Short: "Read a plugin's JSON metadata",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		plugin, err := findPlugin(args[0])
		if err != nil {
			return err
		}
		info, err := plugins.ReadInfo(plugin)
		if err != nil {
			return err
		}
		encoded, err := json.MarshalIndent(info, "", "  ")
		if err != nil {
			return fmt.Errorf("encode plugin info: %w", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), string(encoded))
		return nil
	},
}

var pluginRunCmd = &cobra.Command{
	Use:                "run <name> -- <args>",
	Short:              "Explicitly run a trusted local plugin",
	Args:               cobra.MinimumNArgs(1),
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		plugin, err := findPlugin(args[0])
		if err != nil {
			return err
		}
		pluginArgs := args[1:]
		if len(pluginArgs) > 0 && pluginArgs[0] == "--" {
			pluginArgs = pluginArgs[1:]
		}
		fmt.Fprintf(cmd.ErrOrStderr(), "running trusted plugin: %s\n", plugin.Path)
		return plugins.Run(plugin, pluginArgs, cmd.InOrStdin(), cmd.OutOrStdout(), cmd.ErrOrStderr())
	},
}

func pluginToggleCommand(enable bool) *cobra.Command {
	action := "disable"
	if enable {
		action = "enable"
	}
	return &cobra.Command{
		Use:   action + " <name>",
		Short: strings.ToUpper(action[:1]) + action[1:] + " a plugin in clusterforge.yaml",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := loadConfig()
			if err != nil {
				return err
			}
			name := strings.TrimSpace(args[0])
			if name == "" || strings.ContainsAny(name, `/\\`) {
				return fmt.Errorf("invalid plugin name %q", args[0])
			}
			if enable {
				cfg.Plugins.Enabled = true
				cfg.Plugins.Disabled = removeString(cfg.Plugins.Disabled, name)
			} else if !containsString(cfg.Plugins.Disabled, name) {
				cfg.Plugins.Disabled = append(cfg.Plugins.Disabled, name)
				sort.Strings(cfg.Plugins.Disabled)
			}
			if err := cfg.Save(opts.ConfigPath); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %sd\n", name, action)
			return nil
		},
	}
}

func discoverPlugins() ([]plugins.Plugin, error) {
	if opts.NoPlugins {
		return nil, nil
	}
	if isCI() && !opts.AllowPlugins {
		return nil, fmt.Errorf("plugins are disabled in CI; pass --allow-plugins only for trusted local executables")
	}
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}
	if !cfg.Plugins.Enabled {
		return nil, nil
	}
	directories := make([]string, 0, len(cfg.Plugins.Directories))
	configDir := filepath.Dir(opts.ConfigPath)
	for _, directory := range cfg.Plugins.Directories {
		if !filepath.IsAbs(directory) {
			directory = filepath.Join(configDir, directory)
		}
		directories = append(directories, directory)
	}
	return plugins.Discover(plugins.DiscoverOptions{
		Directories:      directories,
		AllowPathPlugins: cfg.Plugins.AllowPathPlugins,
		DisabledNames:    cfg.Plugins.Disabled,
	})
}

func findPlugin(name string) (plugins.Plugin, error) {
	found, err := discoverPlugins()
	if err != nil {
		return plugins.Plugin{}, err
	}
	return plugins.Find(found, name)
}

func isCI() bool {
	value := strings.ToLower(strings.TrimSpace(os.Getenv("CI")))
	return value != "" && value != "0" && value != "false" && value != "no"
}

func containsString(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func removeString(values []string, unwanted string) []string {
	result := values[:0]
	for _, value := range values {
		if value != unwanted {
			result = append(result, value)
		}
	}
	return result
}

func init() {
	pluginCmd.AddCommand(pluginListCmd)
	pluginCmd.AddCommand(pluginDiscoverCmd)
	pluginCmd.AddCommand(pluginInfoCmd)
	pluginCmd.AddCommand(pluginRunCmd)
	pluginCmd.AddCommand(pluginToggleCommand(true))
	pluginCmd.AddCommand(pluginToggleCommand(false))
}
