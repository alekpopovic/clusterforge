package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{Use: "workspace", Short: "Inspect workspace metadata"}
var teamCmd = &cobra.Command{Use: "team", Short: "Inspect team ownership metadata"}

var workspaceListCmd = &cobra.Command{Use: "list", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	names := make([]string, 0, len(cfg.Workspaces))
	for name := range cfg.Workspaces {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\n", name, cfg.Workspaces[name].Description)
	}
	return nil
}}
var workspaceShowCmd = &cobra.Command{Use: "show <name>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	workspace, ok := cfg.Workspaces[args[0]]
	if !ok {
		return fmt.Errorf("workspace %q not found", args[0])
	}
	fmt.Fprintf(cmd.OutOrStdout(), "name: %s\ndescription: %s\nenvironments:\n", args[0], workspace.Description)
	for _, env := range workspace.Environments {
		fmt.Fprintf(cmd.OutOrStdout(), "- %s\n", env)
	}
	return nil
}}
var workspaceDoctorCmd = &cobra.Command{Use: "doctor <name>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	workspace, ok := cfg.Workspaces[args[0]]
	if !ok {
		return fmt.Errorf("workspace %q not found", args[0])
	}
	if len(workspace.Environments) == 0 {
		fmt.Fprintln(cmd.OutOrStdout(), "WARN\tworkspace has no environments")
		return nil
	}
	for _, env := range workspace.Environments {
		fmt.Fprintf(cmd.OutOrStdout(), "PASS\t%s\t%s\n", env, cfg.Environments[env].Path)
	}
	return nil
}}
var teamListCmd = &cobra.Command{Use: "list", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	names := make([]string, 0, len(cfg.Teams))
	for name := range cfg.Teams {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintln(cmd.OutOrStdout(), name)
	}
	return nil
}}
var teamShowCmd = &cobra.Command{Use: "show <name>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	team, ok := cfg.Teams[args[0]]
	if !ok {
		return fmt.Errorf("team %q not found", args[0])
	}
	fmt.Fprintf(cmd.OutOrStdout(), "name: %s\nowners:\n", args[0])
	for _, owner := range team.Owners {
		fmt.Fprintf(cmd.OutOrStdout(), "- %s\n", owner)
	}
	fmt.Fprintln(cmd.OutOrStdout(), "namespaces:")
	for _, namespace := range team.Namespaces {
		fmt.Fprintf(cmd.OutOrStdout(), "- %s\n", namespace)
	}
	return nil
}}

func init() {
	workspaceCmd.AddCommand(workspaceListCmd, workspaceShowCmd, workspaceDoctorCmd)
	teamCmd.AddCommand(teamListCmd, teamShowCmd)
}
