package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Manage local ClusterForge template packs",
}

var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured template packs",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		if len(cfg.TemplatePacks) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "no template packs configured")
			return nil
		}
		for _, pack := range cfg.TemplatePacks {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\n", pack.Name, pack.Path)
		}
		return nil
	},
}

var templateValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate configured template pack metadata and directories",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		for _, pack := range cfg.TemplatePacks {
			if err := validateTemplatePack(pack.Name, pack.Path); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s: ok\n", pack.Name)
		}
		return nil
	},
}

type templatePackMetadata struct {
	Name                   string   `yaml:"name"`
	Version                string   `yaml:"version"`
	Description            string   `yaml:"description"`
	SupportedClouds        []string `yaml:"supported_clouds"`
	SupportedOrchestrators []string `yaml:"supported_orchestrators"`
}

func validateTemplatePack(name, path string) error {
	data, err := os.ReadFile(filepath.Join(path, "metadata.yaml"))
	if err != nil {
		return fmt.Errorf("read template pack %q metadata: %w", name, err)
	}
	var metadata templatePackMetadata
	if err := yaml.Unmarshal(data, &metadata); err != nil {
		return fmt.Errorf("parse template pack %q metadata: %w", name, err)
	}
	if metadata.Name == "" || metadata.Version == "" {
		return fmt.Errorf("template pack %q metadata requires name and version", name)
	}
	if metadata.Name != name {
		return fmt.Errorf("template pack %q metadata name is %q", name, metadata.Name)
	}
	for _, dir := range []string{"env", "app"} {
		if _, err := os.Stat(filepath.Join(path, dir)); err != nil {
			return fmt.Errorf("template pack %q missing %s directory: %w", name, dir, err)
		}
	}
	return nil
}

func init() {
	templateCmd.AddCommand(templateListCmd)
	templateCmd.AddCommand(templateValidateCmd)
}
