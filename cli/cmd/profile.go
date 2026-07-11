package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
	"sort"
)

var profileCmd = &cobra.Command{Use: "profile", Short: "Inspect Terraform/OpenTofu execution profiles"}
var profileListCmd = &cobra.Command{Use: "list", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	names := make([]string, 0, len(cfg.ExecutionProfiles))
	for n := range cfg.ExecutionProfiles {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\n", n, cfg.ExecutionProfiles[n].Engine)
	}
	return nil
}}
var profileShowCmd = &cobra.Command{Use: "show <name>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	p, err := cfterraform.ResolveProfile(cfg.ExecutionProfiles, args[0])
	if err != nil {
		return err
	}
	data, _ := json.MarshalIndent(p, "", "  ")
	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}}

func init() { profileCmd.AddCommand(profileListCmd, profileShowCmd) }
