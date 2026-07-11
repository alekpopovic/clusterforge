package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/upgradeplanner"
)

var platformJSON bool
var platformCmd = &cobra.Command{Use: "platform", Short: "Inspect and plan platform add-on versions"}

func platformCommand(use string) *cobra.Command {
	return &cobra.Command{Use: use + " <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		env, ok := cfg.Environments[args[0]]
		if !ok {
			return fmt.Errorf("environment %q not found", args[0])
		}
		plan := upgradeplanner.PlatformUpgrade(env.Path, cfg.PlatformVersions)
		if platformJSON {
			data, _ := json.MarshalIndent(plan, "", "  ")
			fmt.Fprintln(cmd.OutOrStdout(), string(data))
			return nil
		}
		for _, c := range plan.Changes {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%t\t%s\n", c.Component, c.Current, c.Desired, c.Change, c.Warning)
		}
		return nil
	}}
}

var platformVersionsCmd = platformCommand("versions")
var platformUpgradeCmd = &cobra.Command{Use: "upgrade"}
var platformPlanCmd = platformCommand("plan")
var platformCheckCmd = platformCommand("check")

func init() {
	for _, c := range []*cobra.Command{platformVersionsCmd, platformPlanCmd, platformCheckCmd} {
		c.Flags().BoolVar(&platformJSON, "json", false, "Print JSON")
	}
	platformUpgradeCmd.AddCommand(platformPlanCmd, platformCheckCmd)
	platformCmd.AddCommand(platformVersionsCmd, platformUpgradeCmd)
}
