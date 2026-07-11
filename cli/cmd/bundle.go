package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/bundle"
)

var bundleEnv, bundleOutput string
var bundleCmd = &cobra.Command{Use: "bundle", Short: "Create and verify offline manifest bundles"}
var bundleCreateCmd = &cobra.Command{Use: "create", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	if err := bundle.Create(".", bundleOutput, bundleEnv, cfg); err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Offline manifest bundle created at %s\n", bundleOutput)
	return nil
}}
var bundleInspectCmd = &cobra.Command{Use: "inspect <bundle>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	summary, err := bundle.Inspect(args[0])
	if err != nil {
		return err
	}
	data, _ := json.MarshalIndent(summary, "", "  ")
	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}}
var bundleVerifyCmd = &cobra.Command{Use: "verify <bundle>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	if err := bundle.Verify(args[0]); err != nil {
		return err
	}
	fmt.Fprintln(cmd.OutOrStdout(), "Bundle checksums verified.")
	return nil
}}

func init() {
	bundleCreateCmd.Flags().StringVar(&bundleEnv, "env", "", "Include one environment")
	bundleCreateCmd.Flags().StringVar(&bundleOutput, "output", "clusterforge-bundle", "Output directory")
	bundleCmd.AddCommand(bundleCreateCmd, bundleInspectCmd, bundleVerifyCmd)
}
