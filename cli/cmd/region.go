package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
)

var regionCmd = &cobra.Command{Use: "region", Short: "Inspect region metadata"}
var regionListCmd = &cobra.Command{Use: "list", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	names := make([]string, 0, len(cfg.Regions))
	for name := range cfg.Regions {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\n", name, cfg.Regions[name])
	}
	return nil
}}
var regionShowCmd = &cobra.Command{Use: "show <name>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	value, ok := cfg.Regions[args[0]]
	if !ok {
		return fmt.Errorf("region alias %q not found", args[0])
	}
	fmt.Fprintf(cmd.OutOrStdout(), "name: %s\nregion: %s\n", args[0], value)
	return nil
}}

func init() { regionCmd.AddCommand(regionListCmd, regionShowCmd) }
