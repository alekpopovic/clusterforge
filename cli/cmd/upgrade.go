package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"github.com/spf13/cobra"
)

const supportedConfigVersion = "0.1.0"

var upgradeYes bool

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Check and apply safe ClusterForge project migrations",
}

var upgradeCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check whether clusterforge.yaml needs migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		actions, err := plannedUpgradeActions(opts.ConfigPath)
		if err != nil {
			return err
		}
		if len(actions) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "clusterforge.yaml is up to date")
			return nil
		}
		fmt.Fprintf(cmd.OutOrStdout(), "migration needed: %d action(s)\n", len(actions))
		for _, action := range actions {
			fmt.Fprintf(cmd.OutOrStdout(), "- %s\n", action)
		}
		return nil
	},
}

var upgradePlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "Show proposed config migrations without writing files",
	RunE: func(cmd *cobra.Command, args []string) error {
		actions, err := plannedUpgradeActions(opts.ConfigPath)
		if err != nil {
			return err
		}
		if len(actions) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "no changes proposed")
			return nil
		}
		fmt.Fprintln(cmd.OutOrStdout(), "proposed changes:")
		for _, action := range actions {
			fmt.Fprintf(cmd.OutOrStdout(), "- %s\n", action)
		}
		return nil
	},
}

var upgradeApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply safe config migrations after creating a backup",
	RunE: func(cmd *cobra.Command, args []string) error {
		actions, err := plannedUpgradeActions(opts.ConfigPath)
		if err != nil {
			return err
		}
		if len(actions) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "clusterforge.yaml is already up to date")
			return nil
		}
		if !upgradeYes {
			return fmt.Errorf("upgrade apply requires --yes after reviewing cf upgrade plan")
		}
		cfg, err := config.Load(opts.ConfigPath)
		if err != nil {
			return err
		}
		backupDir := filepath.Join(".cf", "backups", time.Now().UTC().Format("20060102T150405Z"))
		if err := os.MkdirAll(backupDir, 0o755); err != nil {
			return fmt.Errorf("create backup directory: %w", err)
		}
		data, err := os.ReadFile(opts.ConfigPath)
		if err != nil {
			return fmt.Errorf("read config for backup: %w", err)
		}
		backupPath := filepath.Join(backupDir, filepath.Base(opts.ConfigPath))
		if err := os.WriteFile(backupPath, data, 0o600); err != nil {
			return fmt.Errorf("write config backup: %w", err)
		}
		applyUpgradeDefaults(cfg)
		if err := cfg.Save(opts.ConfigPath); err != nil {
			return err
		}
		printer.Success(fmt.Sprintf("upgraded %s; backup written to %s", opts.ConfigPath, backupPath))
		return nil
	},
}

func plannedUpgradeActions(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}
	cfg, err := config.Load(path)
	if err != nil {
		return nil, err
	}
	text := string(data)
	actions := []string{}
	if !strings.Contains(text, "clusterforge_version:") {
		actions = append(actions, "add clusterforge_version")
	} else if cfg.ClusterForgeVersion != supportedConfigVersion {
		actions = append(actions, fmt.Sprintf("record supported config version %s", supportedConfigVersion))
	}
	if !strings.Contains(text, "policies:") {
		actions = append(actions, "add default policies")
	}
	if !strings.Contains(text, "backends:") {
		actions = append(actions, "add backend config skeleton")
	}
	for name, env := range cfg.Environments {
		expected := fmt.Sprintf("live/%s/%s-%s", name, env.Cloud, env.Orchestrator)
		if strings.TrimSpace(env.Path) == "" || strings.Contains(env.Path, "//") {
			actions = append(actions, fmt.Sprintf("normalize environment path for %s to %s", name, expected))
		}
	}
	return actions, nil
}

func applyUpgradeDefaults(cfg *config.Config) {
	cfg.ClusterForgeVersion = supportedConfigVersion
	cfg.ApplyDefaults()
	if cfg.Backends == nil {
		cfg.Backends = map[string]config.Backend{}
	}
	for name, env := range cfg.Environments {
		if strings.TrimSpace(env.Path) == "" {
			env.Path = fmt.Sprintf("live/%s/%s-%s", name, env.Cloud, env.Orchestrator)
			cfg.Environments[name] = env
		}
		if _, ok := cfg.Backends[name]; !ok {
			cfg.Backends[name] = config.Backend{Type: "local"}
		}
	}
}

func init() {
	upgradeApplyCmd.Flags().BoolVar(&upgradeYes, "yes", false, "Apply the proposed migration without an interactive prompt")
	upgradeCmd.AddCommand(upgradeCheckCmd)
	upgradeCmd.AddCommand(upgradePlanCmd)
	upgradeCmd.AddCommand(upgradeApplyCmd)
	wrapAuditedCommand(upgradeApplyCmd)
}
