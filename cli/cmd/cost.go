package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	cfcost "github.com/textracta/clusterforge/cli/internal/cost"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

var costPlanFile, costStack string
var costJSON, costInfracost bool
var costCmd = &cobra.Command{Use: "cost", Short: "Run heuristic or optional Infracost cost analysis"}
var costEstimateCmd = &cobra.Command{Use: "estimate <env>", Args: cobra.ExactArgs(1), RunE: runCostScan}
var costScanCmd = &cobra.Command{Use: "scan <env>", Args: cobra.ExactArgs(1), RunE: runCostScan}
var costReportCmd = &cobra.Command{Use: "report <env>", Args: cobra.ExactArgs(1), RunE: runCostScan}
var costDiffCmd = &cobra.Command{Use: "diff <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	env, ok := cfg.Environments[args[0]]
	if !ok {
		return fmt.Errorf("environment %q not found", args[0])
	}
	paths, err := resolveStackPaths(env, costStack)
	if err != nil {
		return err
	}
	for _, path := range paths {
		if err := cfcost.RunInfracost(cmd.Context(), path, true, cmd.OutOrStdout(), cmd.ErrOrStderr()); err != nil {
			return err
		}
	}
	return nil
}}

func runCostScan(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	env, ok := cfg.Environments[args[0]]
	if !ok {
		return fmt.Errorf("environment %q not found", args[0])
	}
	paths, err := resolveStackPaths(env, costStack)
	if err != nil {
		return err
	}
	if costInfracost {
		for _, path := range paths {
			if err := cfcost.RunInfracost(cmd.Context(), path, false, cmd.OutOrStdout(), cmd.ErrOrStderr()); err != nil {
				return err
			}
		}
		return nil
	}
	if costPlanFile == "" {
		return fmt.Errorf("--plan-file is required for heuristic mode")
	}
	if len(paths) > 1 {
		return fmt.Errorf("--plan-file requires --stack for stacked environments")
	}
	binary, err := engineBinary(cfg)
	if err != nil {
		return err
	}
	data, err := cfterraform.NewRunner(binary, paths[0], opts.Verbose).ShowPlanJSON(cmd.Context(), costPlanFile)
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
	fmt.Fprintln(cmd.OutOrStdout(), "heuristic cost warnings (not price estimates):")
	for _, warning := range summary.Warnings {
		fmt.Fprintf(cmd.OutOrStdout(), "- %s (%s): %s\n", warning.Address, warning.Category, warning.Message)
	}
	return nil
}
func init() {
	for _, command := range []*cobra.Command{costEstimateCmd, costScanCmd, costReportCmd} {
		command.Flags().StringVar(&costPlanFile, "plan-file", "", "Saved plan file")
		command.Flags().StringVar(&costStack, "stack", "", "Environment stack")
		command.Flags().BoolVar(&costJSON, "json", false, "Print JSON")
		command.Flags().BoolVar(&costInfracost, "infracost", false, "Use installed Infracost")
	}
	costDiffCmd.Flags().StringVar(&costStack, "stack", "", "Environment stack")
	costCmd.AddCommand(costEstimateCmd, costDiffCmd, costScanCmd, costReportCmd)
}
