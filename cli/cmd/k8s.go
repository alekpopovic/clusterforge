package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/alekpopovic/clusterforge/cli/internal/upgradeplanner"
	"github.com/spf13/cobra"
	"os"
)

var k8sTarget string
var k8sCmd = &cobra.Command{Use: "k8s", Short: "Plan Kubernetes upgrades without applying changes"}
var k8sUpgradeCmd = &cobra.Command{Use: "upgrade", Short: "Plan or check an upgrade"}

func k8sPlanCommand(use string) *cobra.Command {
	return &cobra.Command{Use: use + " <cluster-or-env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
		if k8sTarget == "" {
			return fmt.Errorf("--target-version is required")
		}
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		current, node, root := "", "", ""
		if cluster, ok := cfg.Clusters[args[0]]; ok {
			current, node, root = cluster.KubernetesVersion, cluster.NodeVersion, cluster.Path
		} else if env, ok := cfg.Environments[args[0]]; ok {
			current, root = env.KubernetesVersion, env.Path
		} else {
			return fmt.Errorf("cluster or environment %q not found", args[0])
		}
		matrix, _ := os.ReadFile("VERSION_MATRIX.md")
		plan := upgradeplanner.KubernetesUpgrade(upgradeplanner.KubernetesInput{Current: current, Target: k8sTarget, NodeVersion: node, Root: root, Matrix: matrix})
		data, _ := json.MarshalIndent(plan, "", "  ")
		fmt.Fprintln(cmd.OutOrStdout(), string(data))
		if use == "check" && len(plan.Blocking) > 0 {
			return commandExitError{code: 2, message: "Kubernetes upgrade check blocked"}
		}
		return nil
	}}
}

var k8sPlanCmd = k8sPlanCommand("plan")
var k8sCheckCmd = k8sPlanCommand("check")
var k8sVersionsCmd = &cobra.Command{Use: "versions <cluster-or-env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	if c, ok := cfg.Clusters[args[0]]; ok {
		fmt.Fprintf(cmd.OutOrStdout(), "control_plane: %s\nnodes: %s\n", c.KubernetesVersion, c.NodeVersion)
		return nil
	}
	if e, ok := cfg.Environments[args[0]]; ok {
		fmt.Fprintf(cmd.OutOrStdout(), "configured: %s\n", e.KubernetesVersion)
		return nil
	}
	return fmt.Errorf("cluster or environment %q not found", args[0])
}}

func init() {
	k8sPlanCmd.Flags().StringVar(&k8sTarget, "target-version", "", "Target Kubernetes major.minor")
	k8sCheckCmd.Flags().StringVar(&k8sTarget, "target-version", "", "Target Kubernetes major.minor")
	k8sUpgradeCmd.AddCommand(k8sPlanCmd, k8sCheckCmd)
	k8sCmd.AddCommand(k8sUpgradeCmd, k8sVersionsCmd)
}
