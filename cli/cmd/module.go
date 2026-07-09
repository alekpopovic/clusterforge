package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/modulecheck"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

var moduleCheckPath string
var moduleCheckJSON bool

var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Inspect ClusterForge Terraform modules",
}

var moduleCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check Terraform module conformance",
	RunE: func(cmd *cobra.Command, args []string) error {
		report, err := modulecheck.Check(modulecheck.Options{Path: moduleCheckPath})
		if err != nil {
			return err
		}
		if moduleCheckJSON {
			if err := ui.WriteJSON(cmd.OutOrStdout(), report); err != nil {
				return err
			}
		} else {
			for _, module := range report.Modules {
				fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\n", module.Status, module.Path)
				for _, finding := range module.Findings {
					if finding.Severity == modulecheck.Pass {
						continue
					}
					fmt.Fprintf(cmd.OutOrStdout(), "  %s\t%s\t%s\n", finding.Severity, finding.Check, finding.Message)
				}
			}
			fmt.Fprintf(cmd.OutOrStdout(), "status: %s\n", report.Status)
		}
		if report.Status == modulecheck.Fail {
			return fmt.Errorf("module conformance failed")
		}
		return nil
	},
}

func init() {
	moduleCheckCmd.Flags().StringVar(&moduleCheckPath, "path", "modules", "Module directory or module tree to check")
	moduleCheckCmd.Flags().BoolVar(&moduleCheckJSON, "json", false, "Print conformance report as JSON")
	moduleCmd.AddCommand(moduleCheckCmd)
}
