package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alekpopovic/clusterforge/cli/internal/assetinventory"
	cfterraform "github.com/alekpopovic/clusterforge/cli/internal/terraform"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var inventoryFleet bool
var inventoryFormat, inventoryOutput, inventoryStack string
var inventoryCmd = &cobra.Command{Use: "inventory", Short: "Export redacted managed asset inventory"}
var inventoryExportCmd = &cobra.Command{Use: "export [env]", Args: cobra.MaximumNArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	if inventoryFleet && len(args) > 0 {
		return fmt.Errorf("environment and --fleet are mutually exclusive")
	}
	if !inventoryFleet && len(args) == 0 {
		return fmt.Errorf("environment is required unless --fleet is used")
	}
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	names := args
	if inventoryFleet {
		names = nil
		for name := range cfg.Environments {
			names = append(names, name)
		}
		sort.Strings(names)
	}
	var assets []assetinventory.Asset
	binary, err := engineBinary(cfg)
	if err != nil {
		return err
	}
	for _, name := range names {
		env, ok := cfg.Environments[name]
		if !ok {
			return fmt.Errorf("environment %q not found", name)
		}
		paths, err := resolveStackPaths(env, inventoryStack)
		if err != nil {
			return err
		}
		for _, path := range paths {
			stack := inventoryStack
			if data, err := cfterraform.NewRunner(binary, path, opts.Verbose).ShowStateJSON(cmd.Context()); err == nil {
				parsed, err := assetinventory.ParseState(data, name, stack, env.Cloud, env.Region)
				if err != nil {
					return err
				}
				assets = append(assets, parsed...)
			}
			assets = append(assets, moduleAssets(path, name, stack, env.Cloud, env.Region)...)
		}
		assets = append(assets, assetinventory.Asset{Address: "environment." + name, Type: "clusterforge_environment", Name: name, Environment: name, Cloud: env.Cloud, Region: env.Region, Source: "clusterforge-config"})
	}
	for _, path := range appManifestPaths() {
		name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		assets = append(assets, assetinventory.Asset{Address: "app." + name, Type: "clusterforge_app_manifest", Name: name, Source: "app-manifest"})
	}
	var out bytes.Buffer
	switch inventoryFormat {
	case "json":
		err = json.NewEncoder(&out).Encode(assets)
	case "csv":
		err = assetinventory.WriteCSV(&out, assets)
	case "markdown":
		err = assetinventory.WriteMarkdown(&out, assets)
	default:
		return fmt.Errorf("--format must be json, csv, or markdown")
	}
	if err != nil {
		return err
	}
	if inventoryOutput != "" {
		if err := os.WriteFile(inventoryOutput, out.Bytes(), 0600); err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), inventoryOutput)
		return nil
	}
	_, err = cmd.OutOrStdout().Write(out.Bytes())
	return err
}}
var moduleSourceRE = regexp.MustCompile(`(?m)source\s*=\s*"([^"]+)"`)

func moduleAssets(root, env, stack, cloud, region string) []assetinventory.Asset {
	var result []assetinventory.Asset
	_ = filepath.WalkDir(root, func(path string, e os.DirEntry, err error) error {
		if err != nil || e.IsDir() || filepath.Ext(path) != ".tf" {
			return nil
		}
		data, _ := os.ReadFile(path)
		for _, m := range moduleSourceRE.FindAllSubmatch(data, -1) {
			source := string(m[1])
			result = append(result, assetinventory.Asset{Address: "module-source." + filepath.Base(source), Type: "terraform_module_source", Name: filepath.Base(source), Module: source, Environment: env, Stack: stack, Cloud: cloud, Region: region, Source: "terraform-config"})
		}
		return nil
	})
	return result
}
func appManifestPaths() []string {
	paths, _ := filepath.Glob(filepath.Join("apps", "*.yaml"))
	return paths
}
func init() {
	inventoryExportCmd.Flags().BoolVar(&inventoryFleet, "fleet", false, "Export every configured environment")
	inventoryExportCmd.Flags().StringVar(&inventoryStack, "stack", "", "Environment stack")
	inventoryExportCmd.Flags().StringVar(&inventoryFormat, "format", "json", "Output format: json, csv, markdown")
	inventoryExportCmd.Flags().StringVar(&inventoryOutput, "output", "", "Write output with restrictive permissions")
	inventoryCmd.AddCommand(inventoryExportCmd)
}
