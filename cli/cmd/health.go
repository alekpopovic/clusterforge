package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	cfhealth "github.com/textracta/clusterforge/cli/internal/health"
)

var healthJSON bool
var healthCmd = &cobra.Command{Use: "health", Short: "Run read-only platform health checks"}

func healthCommand(use string) *cobra.Command {
	return &cobra.Command{Use: use + " <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		env, ok := cfg.Environments[args[0]]
		if !ok {
			return fmt.Errorf("environment %q not found", args[0])
		}
		settings := cfg.Health.Environments[args[0]]
		kubeconfig := ""
		for _, cluster := range cfg.Clusters {
			if cluster.Environment == args[0] {
				kubeconfig = cluster.KubeconfigPath
				break
			}
		}
		report := cfhealth.Evaluate(cmd.Context(), cfhealth.Input{Environment: args[0], Path: env.Path, Kubeconfig: kubeconfig, SLO: settings.SLO, CheckNodes: settings.Checks.KubernetesNodes, CheckAddons: settings.Checks.PlatformAddons, CheckIngress: settings.Checks.Ingress, CheckApps: settings.Checks.AppHealth})
		if healthJSON {
			data, _ := json.MarshalIndent(report, "", "  ")
			fmt.Fprintln(cmd.OutOrStdout(), string(data))
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "environment: %s\nstatus: %s\n", report.Environment, report.Status)
			for _, check := range report.Checks {
				fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\n", check.Status, check.Name, check.Message)
			}
		}
		if report.Status == cfhealth.Fail {
			return commandExitError{code: 2, message: "health check failed"}
		}
		return nil
	}}
}

var healthCheckCmd = healthCommand("check")
var healthReportCmd = healthCommand("report")

func init() {
	healthCheckCmd.Flags().BoolVar(&healthJSON, "json", false, "Print JSON")
	healthReportCmd.Flags().BoolVar(&healthJSON, "json", false, "Print JSON")
	healthCmd.AddCommand(healthCheckCmd, healthReportCmd)
}
