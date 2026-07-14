package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"github.com/alekpopovic/clusterforge/cli/internal/environment"
	"github.com/alekpopovic/clusterforge/cli/internal/fleet"
	"github.com/alekpopovic/clusterforge/cli/internal/inventory"
	cfterraform "github.com/alekpopovic/clusterforge/cli/internal/terraform"
	"github.com/alekpopovic/clusterforge/cli/internal/ui"
	"github.com/spf13/cobra"
)

var fleetFilter fleet.Filter
var fleetJSON bool
var fleetFailFast bool

var fleetCmd = &cobra.Command{Use: "fleet", Short: "Run read-only operations across cluster inventory"}

func filteredFleet() (*config.Config, []inventory.Cluster, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, nil, err
	}
	return cfg, fleet.Apply(inventory.List(cfg), fleetFilter), nil
}

func writeFleetList(cmd *cobra.Command, clusters []inventory.Cluster) error {
	if fleetJSON {
		return ui.WriteJSON(cmd.OutOrStdout(), map[string]any{"clusters": clusters})
	}
	for _, cluster := range clusters {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%s\t%s\t%s\n", cluster.Name, cluster.Environment, cluster.Cloud, cluster.Orchestrator, cluster.Status, cluster.Path)
	}
	return nil
}

func writeFleetResults(cmd *cobra.Command, operation string, results []fleet.Result) error {
	if fleetJSON {
		return ui.WriteJSON(cmd.OutOrStdout(), map[string]any{"operation": operation, "results": results})
	}
	for _, result := range results {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%s\n", result.Cluster.Name, result.Cluster.Environment, result.Status, result.Message)
	}
	return nil
}

var fleetListCmd = &cobra.Command{
	Use: "list", Short: "List filtered fleet inventory", Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, clusters, err := filteredFleet()
		if err != nil {
			return err
		}
		return writeFleetList(cmd, clusters)
	},
}

var fleetStatusCmd = &cobra.Command{
	Use: "status", Short: "Summarize local fleet inventory readiness", Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, clusters, err := filteredFleet()
		if err != nil {
			return err
		}
		results, _ := fleet.Run(cmd.Context(), clusters, "status", false, func(_ context.Context, cluster inventory.Cluster) (fleet.Result, error) {
			info, err := os.Stat(cluster.Path)
			if err != nil || !info.IsDir() {
				return fleet.Result{}, fmt.Errorf("path %s is not accessible", cluster.Path)
			}
			return fleet.Result{Message: fmt.Sprintf("%s/%s in %s", cluster.Cloud, cluster.Orchestrator, cluster.Region)}, nil
		})
		return writeFleetResults(cmd, "status", results)
	},
}

var fleetDoctorCmd = &cobra.Command{
	Use: "doctor", Short: "Aggregate local doctor checks across the fleet", Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, clusters, err := filteredFleet()
		if err != nil {
			return err
		}
		results, _ := fleet.Run(cmd.Context(), clusters, "doctor", false, func(_ context.Context, cluster inventory.Cluster) (fleet.Result, error) {
			report := inspectCluster(cluster)
			result := fleet.Result{Status: string(report.Status), Message: fmt.Sprintf("%d local check(s)", len(report.Checks))}
			if report.Status == doctorFail {
				return result, fmt.Errorf("one or more local checks failed")
			}
			return result, nil
		})
		return writeFleetResults(cmd, "doctor", results)
	},
}

var fleetPolicyCmd = &cobra.Command{Use: "policy", Short: "Run read-only fleet policy inspection"}

var fleetPolicyCheckCmd = &cobra.Command{
	Use: "check", Short: "Check configured safety gates across the fleet", Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, clusters, err := filteredFleet()
		if err != nil {
			return err
		}
		results, _ := fleet.Run(cmd.Context(), clusters, "policy-check", false, func(_ context.Context, cluster inventory.Cluster) (fleet.Result, error) {
			production := strings.EqualFold(cluster.Status, "production") || environment.IsProduction(cluster.Environment)
			if production && (!cfg.Policies.RequirePlanFileForApply || !cfg.Policies.BlockDestroyInProd) {
				return fleet.Result{Status: "warn", Message: "production safety gates are incomplete"}, nil
			}
			return fleet.Result{Message: "configured read-only policy inspection passed"}, nil
		})
		return writeFleetResults(cmd, "policy-check", results)
	},
}

func fleetDriftPaths(cfg *config.Config, cluster inventory.Cluster) ([]string, error) {
	env, ok := cfg.Environments[cluster.Environment]
	if ok && env.EffectiveLayout() == "stacked" {
		return env.StackPaths("")
	}
	return []string{cluster.Path}, nil
}

func runFleetDrift(ctx context.Context, cfg *config.Config, binary string, cluster inventory.Cluster, stderr io.Writer) (fleet.Result, error) {
	paths, err := fleetDriftPaths(cfg, cluster)
	if err != nil {
		return fleet.Result{}, err
	}
	drifted := false
	for index, path := range paths {
		planFile := filepath.Join(".cf", "plans", fmt.Sprintf("fleet-%s-%d.tfplan", safeFleetName(cluster.Name), index+1))
		if err := os.MkdirAll(filepath.Join(path, filepath.Dir(planFile)), 0o755); err != nil {
			return fleet.Result{}, fmt.Errorf("create drift plan directory: %w", err)
		}
		runner := cfterraform.NewRunner(binary, path, opts.Verbose)
		runner.Stdout, runner.Stderr = io.Discard, stderr
		code, err := runner.PlanDetailedExitCode(ctx, planFile, nil)
		if err != nil {
			return fleet.Result{}, err
		}
		if code == 2 {
			drifted = true
		}
	}
	if drifted {
		return fleet.Result{Status: "warn", Message: "drift detected", Drift: true}, nil
	}
	return fleet.Result{Message: "no drift detected"}, nil
}

func safeFleetName(name string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '-' || r == '_' {
			return r
		}
		return '_'
	}, name)
}

var fleetDriftCmd = &cobra.Command{
	Use: "drift", Short: "Run read-only detailed-exitcode plans across the fleet", Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, clusters, err := filteredFleet()
		if err != nil {
			return err
		}
		binary, err := engineBinary(cfg)
		if err != nil {
			return err
		}
		results, runErr := fleet.Run(cmd.Context(), clusters, "drift", fleetFailFast, func(ctx context.Context, cluster inventory.Cluster) (fleet.Result, error) {
			return runFleetDrift(ctx, cfg, binary, cluster, cmd.ErrOrStderr())
		})
		if err := writeFleetResults(cmd, "drift", results); err != nil {
			return err
		}
		return runErr
	},
}

func init() {
	fleetCmd.PersistentFlags().StringVar(&fleetFilter.Environment, "environment", "", "Filter by environment")
	fleetCmd.PersistentFlags().StringVar(&fleetFilter.Cloud, "cloud", "", "Filter by cloud")
	fleetCmd.PersistentFlags().StringVar(&fleetFilter.Orchestrator, "orchestrator", "", "Filter by orchestrator")
	fleetCmd.PersistentFlags().StringVar(&fleetFilter.Status, "status", "", "Filter by lifecycle status")
	fleetCmd.PersistentFlags().StringVar(&fleetFilter.Region, "region", "", "Filter by region")
	fleetCmd.PersistentFlags().BoolVar(&fleetJSON, "json", false, "Print structured JSON output")
	fleetDriftCmd.Flags().BoolVar(&fleetFailFast, "fail-fast", false, "Stop and fail on the first cluster error")
	fleetPolicyCmd.AddCommand(fleetPolicyCheckCmd)
	fleetCmd.AddCommand(fleetListCmd, fleetStatusCmd, fleetDoctorCmd, fleetDriftCmd, fleetPolicyCmd)
}
