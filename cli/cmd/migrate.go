package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/migration"
)

var migratePath string
var migrateJSON bool
var migrateCmd = &cobra.Command{Use: "migrate", Short: "Analyze existing Terraform repositories without modifying them"}
var migrateAnalyzeCmd = &cobra.Command{Use: "analyze", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	report, err := migration.Analyze(migratePath)
	if err != nil {
		return err
	}
	if migrateJSON {
		data, _ := report.JSON()
		fmt.Fprintln(cmd.OutOrStdout(), string(data))
		return nil
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Terraform roots: %d\nProviders: %s\nArchitecture: %s\nRisks: %d\n", len(report.RootModules), strings.Join(report.Providers, ", "), strings.Join(report.Architecture, ", "), len(report.Risks))
	return nil
}}
var migrateReportCmd = &cobra.Command{Use: "report", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	report, err := migration.Analyze(migratePath)
	if err != nil {
		return err
	}
	out := cmd.OutOrStdout()
	fmt.Fprintln(out, "# ClusterForge migration assessment")
	fmt.Fprintln(out)
	fmt.Fprintf(out, "Analyzed `%s` read-only. No Terraform or cloud command was run.\n\n", report.Path)
	fmt.Fprintf(out, "## Detected architecture\n\n- %s\n\n", strings.Join(report.Architecture, "\n- "))
	fmt.Fprintf(out, "## ClusterForge equivalents\n\n- %s\n\n", strings.Join(report.EquivalentModules, "\n- "))
	fmt.Fprintf(out, "## Risks\n\n- %s\n\n", strings.Join(report.Risks, "\n- "))
	fmt.Fprintf(out, "## Suggested migration steps\n\n")
	for index, step := range report.SuggestedSteps {
		fmt.Fprintf(out, "%d. %s\n", index+1, step)
	}
	fmt.Fprintf(out, "\n## Import and adoption notes\n\n- %s\n", strings.Join(report.AdoptionNotes, "\n- "))
	return nil
}}

func init() {
	for _, command := range []*cobra.Command{migrateAnalyzeCmd, migrateReportCmd} {
		command.Flags().StringVar(&migratePath, "path", ".", "Terraform repository path")
	}
	migrateAnalyzeCmd.Flags().BoolVar(&migrateJSON, "json", false, "Print JSON report")
	migrateCmd.AddCommand(migrateAnalyzeCmd, migrateReportCmd)
}
