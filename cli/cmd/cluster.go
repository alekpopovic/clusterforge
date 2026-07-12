package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/alekpopovic/clusterforge/cli/internal/inventory"
	cfterraform "github.com/alekpopovic/clusterforge/cli/internal/terraform"
	"github.com/alekpopovic/clusterforge/cli/internal/ui"
	"github.com/spf13/cobra"
)

var clusterListJSON bool
var clusterShowJSON bool
var clusterDoctorJSON bool
var clusterOutputsJSON bool

var clusterCmd = &cobra.Command{Use: "cluster", Short: "Inspect configured cluster inventory"}

var clusterListCmd = &cobra.Command{
	Use: "list", Short: "List configured clusters", Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		clusters := inventory.List(cfg)
		if clusterListJSON {
			return ui.WriteJSON(cmd.OutOrStdout(), map[string]any{"clusters": clusters})
		}
		for _, cluster := range clusters {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%s\t%s\t%s\t%s\n", cluster.Name, cluster.Environment, cluster.Cloud, cluster.Region, cluster.Orchestrator, cluster.Status, cluster.Path)
		}
		return nil
	},
}

var clusterShowCmd = &cobra.Command{
	Use: "show <name>", Short: "Show one cluster inventory record", Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		cluster, err := inventory.Find(cfg, args[0])
		if err != nil {
			return err
		}
		if clusterShowJSON {
			return ui.WriteJSON(cmd.OutOrStdout(), cluster)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "name: %s\nenvironment: %s\ncloud: %s\nregion: %s\norchestrator: %s\nstatus: %s\npath: %s\n", cluster.Name, cluster.Environment, cluster.Cloud, cluster.Region, cluster.Orchestrator, cluster.Status, cluster.Path)
		if cluster.KubeconfigPath != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "kubeconfig_path: %s\n", cluster.KubeconfigPath)
		}
		return nil
	},
}

type clusterDoctorReport struct {
	Cluster inventory.Cluster `json:"cluster"`
	Status  doctorStatus      `json:"status"`
	Checks  []doctorCheck     `json:"checks"`
}

func inspectCluster(cluster inventory.Cluster) clusterDoctorReport {
	report := clusterDoctorReport{Cluster: cluster}
	if info, err := os.Stat(cluster.Path); err != nil || !info.IsDir() {
		report.Checks = append(report.Checks, doctorCheck{Name: "cluster.path", Status: doctorFail, Message: fmt.Sprintf("%s is not an accessible directory", cluster.Path)})
	} else {
		report.Checks = append(report.Checks, doctorCheck{Name: "cluster.path", Status: doctorPass, Message: fmt.Sprintf("%s exists", cluster.Path)})
	}
	if inventory.IsKubernetes(cluster.Orchestrator) {
		if strings.TrimSpace(cluster.KubeconfigPath) == "" {
			report.Checks = append(report.Checks, doctorCheck{Name: "cluster.kubeconfig_reference", Status: doctorWarn, Message: "no kubeconfig_path reference configured; kubeconfig content is never stored"})
		} else {
			report.Checks = append(report.Checks, doctorCheck{Name: "cluster.kubeconfig_reference", Status: doctorPass, Message: cluster.KubeconfigPath})
		}
	}
	report.Status = reportStatus(report.Checks)
	return report
}

var clusterDoctorCmd = &cobra.Command{
	Use: "doctor <name>", Short: "Check one inventory record without live cloud access", Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		cluster, err := inventory.Find(cfg, args[0])
		if err != nil {
			return err
		}
		report := inspectCluster(cluster)
		if clusterDoctorJSON {
			if err := ui.WriteJSON(cmd.OutOrStdout(), report); err != nil {
				return err
			}
		} else {
			for _, check := range report.Checks {
				fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\n", check.Status, check.Name, check.Message)
			}
		}
		if report.Status == doctorFail {
			return fmt.Errorf("cluster doctor found one or more hard failures")
		}
		return nil
	},
}

var clusterKubeconfigCmd = &cobra.Command{
	Use: "kubeconfig <name>", Short: "Print a configured kubeconfig path reference, never its content", Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		cluster, err := inventory.Find(cfg, args[0])
		if err != nil {
			return err
		}
		if !inventory.IsKubernetes(cluster.Orchestrator) {
			return fmt.Errorf("cluster %q orchestrator %q does not use kubeconfig", cluster.Name, cluster.Orchestrator)
		}
		if strings.TrimSpace(cluster.KubeconfigPath) == "" {
			return fmt.Errorf("cluster %q has no kubeconfig_path reference configured", cluster.Name)
		}
		fmt.Fprintln(cmd.OutOrStdout(), cluster.KubeconfigPath)
		return nil
	},
}

var clusterOutputsCmd = &cobra.Command{
	Use: "outputs <name>", Short: "Read Terraform/OpenTofu outputs for one cluster path", Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		cluster, err := inventory.Find(cfg, args[0])
		if err != nil {
			return err
		}
		binary, err := engineBinary(cfg)
		if err != nil {
			return err
		}
		runner := cfterraform.NewRunner(binary, cluster.Path, opts.Verbose)
		runner.Stdout, runner.Stderr = cmd.OutOrStdout(), cmd.ErrOrStderr()
		return runner.Output(cmd.Context(), clusterOutputsJSON)
	},
}

func init() {
	clusterListCmd.Flags().BoolVar(&clusterListJSON, "json", false, "Print cluster inventory as JSON")
	clusterShowCmd.Flags().BoolVar(&clusterShowJSON, "json", false, "Print cluster details as JSON")
	clusterDoctorCmd.Flags().BoolVar(&clusterDoctorJSON, "json", false, "Print cluster doctor report as JSON")
	clusterOutputsCmd.Flags().BoolVar(&clusterOutputsJSON, "json", false, "Request Terraform output JSON")
	clusterCmd.AddCommand(clusterListCmd, clusterShowCmd, clusterDoctorCmd, clusterKubeconfigCmd, clusterOutputsCmd)
}
