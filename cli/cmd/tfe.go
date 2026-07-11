package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/generator"
	"os"
	"path/filepath"
	"sort"
)

var tfeCmd = &cobra.Command{Use: "tfe", Short: "Generate optional HCP Terraform configuration"}
var tfeWorkspaceCmd = &cobra.Command{Use: "workspace"}
var tfeBackendCmd = &cobra.Command{Use: "backend"}

func tfeConfig(env string) (string, string, string, bool, error) {
	cfg, err := loadConfig()
	if err != nil {
		return "", "", "", false, err
	}
	if !cfg.TerraformCloud.Enabled {
		return "", "", "", false, nil
	}
	workspace, ok := cfg.TerraformCloud.Workspaces[env]
	if !ok {
		return "", "", "", true, fmt.Errorf("Terraform Cloud workspace for environment %q is not configured", env)
	}
	environment, ok := cfg.Environments[env]
	if !ok {
		return "", "", "", true, fmt.Errorf("environment %q not found", env)
	}
	return cfg.TerraformCloud.Organization, workspace.Name, environment.Path, true, nil
}

var tfeWorkspaceListCmd = &cobra.Command{Use: "list", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	if !cfg.TerraformCloud.Enabled {
		fmt.Fprintln(cmd.OutOrStdout(), "Terraform Cloud integration disabled")
		return nil
	}
	names := make([]string, 0, len(cfg.TerraformCloud.Workspaces))
	for n := range cfg.TerraformCloud.Workspaces {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\n", n, cfg.TerraformCloud.Workspaces[n].Name)
	}
	return nil
}}
var tfeBackendRenderCmd = &cobra.Command{Use: "render <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	org, workspace, _, enabled, err := tfeConfig(args[0])
	if err != nil {
		return err
	}
	if !enabled {
		fmt.Fprintln(cmd.OutOrStdout(), "Terraform Cloud integration disabled")
		return nil
	}
	rendered, err := generator.RenderTerraformCloudBackend(org, workspace)
	if err != nil {
		return err
	}
	fmt.Fprint(cmd.OutOrStdout(), rendered)
	return nil
}}
var tfeWorkspaceGenerateCmd = &cobra.Command{Use: "generate <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	org, workspace, path, enabled, err := tfeConfig(args[0])
	if err != nil {
		return err
	}
	if !enabled {
		fmt.Fprintln(cmd.OutOrStdout(), "Terraform Cloud integration disabled")
		return nil
	}
	rendered, err := generator.RenderTerraformCloudBackend(org, workspace)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	target := filepath.Join(path, "backend.tf")
	if err := os.WriteFile(target, []byte(rendered), 0644); err != nil {
		return err
	}
	fmt.Fprintln(cmd.OutOrStdout(), target)
	return nil
}}

func init() {
	tfeWorkspaceCmd.AddCommand(tfeWorkspaceListCmd, tfeWorkspaceGenerateCmd)
	tfeBackendCmd.AddCommand(tfeBackendRenderCmd)
	tfeCmd.AddCommand(tfeWorkspaceCmd, tfeBackendCmd)
}
