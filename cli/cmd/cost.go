package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cfcost "github.com/textracta/clusterforge/cli/internal/cost"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

var costPlanFile string
var costJSON bool

var costCmd = &cobra.Command{
	Use:   "cost",
	Short: "Run heuristic cost-awareness checks",
}

var costEstimateCmd = &cobra.Command{
	Use:   "estimate <env>",
	Short: "Estimate cost risk from a saved plan file",
	Args:  cobra.ExactArgs(1),
	RunE:  runCostScan,
}

var costScanCmd = &cobra.Command{
	Use:   "scan <env>",
	Short: "Scan a saved plan for expensive resource categories",
	Args:  cobra.ExactArgs(1),
	RunE:  runCostScan,
}

func runCostScan(cmd *cobra.Command, args []string) error {
	if costPlanFile == "" {
		return fmt.Errorf("--plan-file is required")
	}
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	env, ok := cfg.Environments[args[0]]
	if !ok {
		return fmt.Errorf("environment %q not found", args[0])
	}
	binary, err := engineBinary(cfg)
	if err != nil {
		return err
	}
	runner := cfterraform.NewRunner(binary, env.Path, opts.Verbose)
	data, err := runner.ShowPlanJSON(cmd.Context(), costPlanFile)
	if err != nil {
		return err
	}
	summary, err := cfcost.ScanPlanJSON(data)
	if err != nil {
		return err
	}
	if costJSON {
		return ui.WriteJSON(cmd.OutOrStdout(), summary)
	}
	if len(summary.Warnings) == 0 {
		fmt.Fprintln(cmd.OutOrStdout(), "no built-in cost warnings detected")
		return nil
	}
	fmt.Fprintln(cmd.OutOrStdout(), "heuristic cost warnings:")
	for _, warning := range summary.Warnings {
		fmt.Fprintf(cmd.OutOrStdout(), "- %s (%s): %s\n", warning.Address, warning.Category, warning.Message)
	}
	fmt.Fprintln(cmd.OutOrStdout(), "Install Infracost for pricing-backed estimates; built-in checks are advisory only.")
	return nil
}

func init() {
	for _, command := range []*cobra.Command{costEstimateCmd, costScanCmd} {
		command.Flags().StringVar(&costPlanFile, "plan-file", "", "Saved Terraform/OpenTofu plan file to scan")
		command.Flags().BoolVar(&costJSON, "json", false, "Print cost scan result as JSON")
	}
	costCmd.AddCommand(costEstimateCmd)
	costCmd.AddCommand(costScanCmd)
}
