package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"github.com/alekpopovic/clusterforge/cli/internal/templatepacks"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var templateCmd = &cobra.Command{Use: "template", Short: "Manage versioned ClusterForge template packs"}

var templateListCmd = &cobra.Command{Use: "list", Short: "List configured template packs", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	if len(cfg.TemplatePacks) == 0 {
		fmt.Fprintln(cmd.OutOrStdout(), "no template packs configured")
		return nil
	}
	for _, pack := range cfg.TemplatePacks {
		status, source := "enabled", pack.Source
		if !pack.IsEnabled() {
			status = "disabled"
		}
		if source == "" {
			source = pack.Path
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%s\n", pack.Name, pack.Version, status, source)
	}
	return nil
}}

var templateFetchCmd = &cobra.Command{Use: "fetch <name>", Short: "Fetch a configured pack into the cache", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error { return fetchTemplatePack(cmd, args[0], false) }}
var templateUpdateCmd = &cobra.Command{Use: "update <name>", Short: "Replace a cached pack from its source", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error { return fetchTemplatePack(cmd, args[0], true) }}
var templateValidateCmd = &cobra.Command{Use: "validate <name>", Short: "Validate one pack without executing it", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	pack, err := configuredTemplatePack(cfg, args[0])
	if err != nil {
		return err
	}
	path := templatePackPath(pack)
	if err := validateTemplatePack(pack.Name, path); err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "%s: ok (%s)\n", pack.Name, path)
	return nil
}}

var templateCacheCmd = &cobra.Command{Use: "cache", Short: "Manage the template pack cache"}
var templateCacheClearCmd = &cobra.Command{Use: "clear", Short: "Remove all cached template packs", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	path := filepath.Join(filepath.Dir(opts.ConfigPath), ".cf", "cache", "template-packs")
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("clear template pack cache: %w", err)
	}
	fmt.Fprintf(cmd.OutOrStdout(), "cleared %s\n", path)
	return nil
}}

type templatePackMetadata struct {
	Name                   string   `yaml:"name"`
	Version                string   `yaml:"version"`
	Description            string   `yaml:"description"`
	SupportedClouds        []string `yaml:"supported_clouds"`
	SupportedOrchestrators []string `yaml:"supported_orchestrators"`
}

var likelySecret = regexp.MustCompile(`(?i)(password|secret|api[_-]?key|access[_-]?key)\s*[:=]\s*["']?[A-Za-z0-9+/=_-]{8,}`)

func fetchTemplatePack(cmd *cobra.Command, name string, update bool) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	pack, err := configuredTemplatePack(cfg, name)
	if err != nil {
		return err
	}
	if !pack.IsEnabled() {
		return fmt.Errorf("template pack %q is disabled", name)
	}
	if pack.Source == "" {
		return fmt.Errorf("template pack %q uses a local path and does not need fetching", name)
	}
	destination := templatePackPath(pack)
	if !update {
		if _, err := os.Stat(destination); err == nil {
			return fmt.Errorf("template pack %q is cached; use template update %s", name, name)
		}
	}
	source := resolveTemplateSource(pack.Source)
	fmt.Fprintf(cmd.OutOrStdout(), "fetching %s from %s\n", name, pack.Source)
	if strings.HasPrefix(source, "git::") {
		parsed, err := templatepacks.ParseGitSource(source)
		if err != nil {
			return err
		}
		if templatepacks.WeakRef(parsed.Ref) {
			fmt.Fprintf(cmd.ErrOrStderr(), "warning: ref %q is mutable; pin a tag or commit SHA\n", parsed.Ref)
		}
	}
	if err := templatepacks.Fetch(source, destination); err != nil {
		return err
	}
	if err := validateTemplatePack(pack.Name, destination); err != nil {
		_ = os.RemoveAll(destination)
		return fmt.Errorf("fetched pack failed validation: %w", err)
	}
	fmt.Fprintf(cmd.OutOrStdout(), "cached at %s\n", destination)
	return nil
}

func configuredTemplatePack(cfg *config.Config, name string) (config.TemplatePack, error) {
	for _, pack := range cfg.TemplatePacks {
		if pack.Name == name {
			return pack, nil
		}
	}
	return config.TemplatePack{}, fmt.Errorf("template pack %q is not configured", name)
}

func resolveTemplateSource(source string) string {
	if strings.HasPrefix(source, "git::") {
		return source
	}
	prefix, path := "", source
	for _, candidate := range []string{"path::", "archive::"} {
		if strings.HasPrefix(source, candidate) {
			prefix, path = candidate, strings.TrimPrefix(source, candidate)
		}
	}
	if !filepath.IsAbs(path) {
		path = filepath.Join(filepath.Dir(opts.ConfigPath), path)
	}
	return prefix + path
}

func templatePackPath(pack config.TemplatePack) string {
	if pack.Source == "" {
		if filepath.IsAbs(pack.Path) {
			return pack.Path
		}
		return filepath.Join(filepath.Dir(opts.ConfigPath), pack.Path)
	}
	return templatepacks.CachePath(filepath.Dir(opts.ConfigPath), pack.Name, pack.Version)
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
	if len(metadata.SupportedClouds) == 0 || len(metadata.SupportedOrchestrators) == 0 {
		return fmt.Errorf("template pack %q must declare supported clouds and orchestrators", name)
	}
	for _, directory := range []string{"env", "app"} {
		if err := requireTemplateFile(filepath.Join(path, directory)); err != nil {
			return fmt.Errorf("template pack %q %s: %w", name, directory, err)
		}
	}
	return filepath.WalkDir(path, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if info.Mode().Perm()&0o111 != 0 {
			return fmt.Errorf("template pack contains executable file %s", filePath)
		}
		if info.Size() > 2*1024*1024 {
			return fmt.Errorf("template file exceeds 2 MiB: %s", filePath)
		}
		content, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		if likelySecret.Match(content) {
			return fmt.Errorf("possible secret in template file %s", filePath)
		}
		return nil
	})
}

func requireTemplateFile(directory string) error {
	found := false
	err := filepath.WalkDir(directory, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			found = true
		}
		return nil
	})
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("must contain at least one template file")
	}
	return nil
}

func init() {
	templateCacheCmd.AddCommand(templateCacheClearCmd)
	templateCmd.AddCommand(templateListCmd, templateFetchCmd, templateUpdateCmd, templateValidateCmd, templateCacheCmd)
}
