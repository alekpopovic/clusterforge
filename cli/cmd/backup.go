package cmd

import (
	"fmt"

	"github.com/alekpopovic/clusterforge/cli/internal/backup"
	"github.com/spf13/cobra"
)

var backupEvidence string
var backupCmd = &cobra.Command{Use: "backup", Short: "Validate local Kubernetes backup and restore-test evidence"}

func backupReport(environment string) (backup.Report, error) {
	cfg, err := loadConfig()
	if err != nil {
		return backup.Report{}, err
	}
	env, ok := cfg.Environments[environment]
	if !ok {
		return backup.Report{}, fmt.Errorf("environment %q not found", environment)
	}
	if env.Orchestrator != "" && env.Orchestrator != "eks" && env.Orchestrator != "kubernetes" && env.Orchestrator != "gke" && env.Orchestrator != "aks" {
		return backup.Report{}, fmt.Errorf("environment %q is not a Kubernetes environment", environment)
	}
	return backup.BuildReport(environment, env.Path, env.Cloud, backupEvidence)
}

var backupCheckCmd = &cobra.Command{Use: "check <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	report, err := backupReport(args[0])
	if err != nil {
		return err
	}
	for _, check := range report.Checks {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\n", check.Status, check.ID, check.Message)
	}
	fmt.Fprintf(cmd.OutOrStdout(), "overall\t%s\n", report.Overall)
	return nil
}}
var backupPlanCmd = &cobra.Command{Use: "plan <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	if _, err := backupReport(args[0]); err != nil {
		return err
	}
	for index, step := range backup.Plan(args[0]) {
		fmt.Fprintf(cmd.OutOrStdout(), "%d. %s\n", index+1, step)
	}
	return nil
}}
var backupReportCmd = &cobra.Command{Use: "report <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	report, err := backupReport(args[0])
	if err != nil {
		return err
	}
	data, err := report.JSON()
	if err != nil {
		return err
	}
	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}}

func init() {
	for _, command := range []*cobra.Command{backupCheckCmd, backupPlanCmd, backupReportCmd} {
		command.Flags().StringVar(&backupEvidence, "evidence", backup.EvidencePath, "Backup test evidence YAML")
	}
	backupCmd.AddCommand(backupCheckCmd, backupPlanCmd, backupReportCmd)
}
