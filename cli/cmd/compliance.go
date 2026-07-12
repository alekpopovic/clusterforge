package cmd

import (
	"fmt"

	"github.com/alekpopovic/clusterforge/cli/internal/compliance"
	"github.com/spf13/cobra"
)

var compliancePack, complianceFormat string
var complianceCmd = &cobra.Command{Use: "compliance", Short: "Render non-certifying compliance control mappings"}
var complianceListCmd = &cobra.Command{Use: "list", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	for _, pack := range compliance.List() {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\n", pack.ID, pack.Name)
	}
	return nil
}}
var complianceReportCmd = &cobra.Command{Use: "report", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	if compliancePack == "" {
		return fmt.Errorf("--pack is required")
	}
	pack, err := compliance.Get(compliancePack)
	if err != nil {
		return err
	}
	return compliance.Render(cmd.OutOrStdout(), pack, complianceFormat)
}}

func init() {
	complianceReportCmd.Flags().StringVar(&compliancePack, "pack", "", "Compliance mapping pack")
	complianceReportCmd.Flags().StringVar(&complianceFormat, "format", "markdown", "Output format: markdown or json")
	complianceCmd.AddCommand(complianceListCmd, complianceReportCmd)
}
