package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	cfapp "github.com/textracta/clusterforge/cli/internal/app"
	"github.com/textracta/clusterforge/cli/internal/config"
	"github.com/textracta/clusterforge/cli/internal/scaffold"
)

var wizardDefaults bool
var wizardCmd = &cobra.Command{Use: "wizard", Short: "Interactively scaffold a ClusterForge project", Args: cobra.NoArgs, RunE: runScaffoldWizard}

func runScaffoldWizard(cmd *cobra.Command, args []string) error {
	plan := scaffold.Defaults()
	prompts := newPromptSession(cmd)
	if !wizardDefaults {
		if opts.NonInteractive {
			return fmt.Errorf("wizard requires --defaults in non-interactive mode")
		}
		var err error
		if plan.Project, err = prompts.String("project name", plan.Project); err != nil {
			return err
		}
		if plan.Owner, err = prompts.String("team/owner", plan.Owner); err != nil {
			return err
		}
		if plan.Target, err = prompts.String("target (aws-eks/aws-ecs/existing-kubernetes/local-kind/azure-aks/gcp-gke)", plan.Target); err != nil {
			return err
		}
		if plan.Environment, err = prompts.String("environment (dev/staging/prod)", plan.Environment); err != nil {
			return err
		}
		if plan.Backend, err = prompts.String("backend (local/s3/terraform-cloud)", plan.Backend); err != nil {
			return err
		}
		addons, err := prompts.String("platform add-ons (comma-separated)", strings.Join(plan.Addons, ","))
		if err != nil {
			return err
		}
		plan.Addons = splitCSV(addons)
		if plan.DemoApp, err = prompts.Bool("create demo app", plan.DemoApp); err != nil {
			return err
		}
	}
	if err := plan.Validate(); err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Scaffolding summary\n%s\n", plan.Summary())
	if !wizardDefaults {
		confirmed, err := prompts.Bool("write these files", false)
		if err != nil {
			return err
		}
		if !confirmed {
			return fmt.Errorf("scaffolding cancelled")
		}
	}
	if _, err := os.Stat(opts.ConfigPath); err == nil {
		return fmt.Errorf("%s already exists", opts.ConfigPath)
	}
	cloud, orchestrator, region := scaffold.Target(plan.Target)
	cfg := config.DefaultConfig(plan.Project)
	cfg.Organization = config.Organization{Name: plan.Owner, Owner: plan.Owner}
	cfg.Defaults.Cloud, cfg.Defaults.Orchestrator, cfg.Defaults.Region = cloud, orchestrator, region
	cfg.Policies.RequirePlanFileForApply = plan.ProductionPlanRequired
	cfg.Policies.BlockDestroyInProd = plan.DestroyBlocked
	envPath := filepath.Join("live", plan.Environment, cloud+"-"+orchestrator)
	cfg.Environments[plan.Environment] = config.Environment{Cloud: cloud, Orchestrator: orchestrator, Region: region, Path: envPath, Layout: "simple"}
	if plan.Backend == "s3" {
		cfg.Backends[plan.Environment] = config.Backend{Type: "s3", Bucket: "replace-with-reviewed-state-bucket", Region: region}
	}
	if plan.Backend == "terraform-cloud" {
		cfg.TerraformCloud = config.TerraformCloud{Enabled: true, Organization: "replace-with-organization", Project: plan.Project, Workspaces: map[string]config.TerraformCloudWorkspace{plan.Environment: {Name: plan.Project + "-" + plan.Environment}}}
	}
	if err := os.MkdirAll(envPath, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll("apps", 0o755); err != nil {
		return err
	}
	if err := config.Write(opts.ConfigPath, cfg, false); err != nil {
		return err
	}
	if plan.DemoApp {
		if _, err := cfapp.Add(".", "demo", cfapp.AddOptions{Image: "nginx:1.27", Port: 80, Replicas: 1}); err != nil {
			return err
		}
	}
	fmt.Fprintln(cmd.OutOrStdout(), "Project scaffold created. Review files, then run generate/plan; no apply was run.")
	return nil
}

func splitCSV(value string) []string {
	var result []string
	for _, item := range strings.Split(value, ",") {
		if item = strings.TrimSpace(item); item != "" {
			result = append(result, item)
		}
	}
	return result
}

func init() {
	wizardCmd.Flags().BoolVar(&wizardDefaults, "defaults", false, "Use safe demo defaults without prompts")
}
