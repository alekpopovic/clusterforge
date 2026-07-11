package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	cfapp "github.com/textracta/clusterforge/cli/internal/app"
	"github.com/textracta/clusterforge/cli/internal/dashboard"
	"github.com/textracta/clusterforge/cli/internal/runbooks"
	"github.com/textracta/clusterforge/cli/internal/servicecatalog"
)

var dashboardEnv, dashboardOutput string
var dashboardFleet bool

var dashboardCmd = &cobra.Command{Use: "dashboard", Short: "Prepare local data for a dashboard"}

var dashboardExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export dashboard data from local project files",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if dashboardFleet && dashboardEnv != "" {
			return fmt.Errorf("--env and --fleet cannot be used together")
		}
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		if dashboardEnv != "" {
			if _, ok := cfg.Environments[dashboardEnv]; !ok {
				return fmt.Errorf("environment %q not found", dashboardEnv)
			}
		}
		apps := map[string]cfapp.Manifest{}
		names, err := cfapp.List(".")
		if err != nil {
			return err
		}
		for _, name := range names {
			manifest, loadErr := cfapp.Load(cfapp.ManifestPath(".", name))
			if loadErr != nil {
				return loadErr
			}
			apps[name] = manifest
		}
		catalog := servicecatalog.Catalog{}
		if _, statErr := os.Stat(servicecatalog.DefaultPath); statErr == nil {
			catalog, err = servicecatalog.Load(servicecatalog.DefaultPath)
			if err != nil {
				return err
			}
		} else if !os.IsNotExist(statErr) {
			return statErr
		}
		books, err := runbooks.Discover("docs/incident-response", "docs/dr")
		if err != nil {
			return err
		}
		moduleCatalog, err := readOptionalFile("MODULE_CATALOG.md")
		if err != nil {
			return err
		}
		evidence := map[string]bool{}
		for _, kind := range []string{"policy", "drift", "cost"} {
			_, statErr := os.Stat(".cf/dashboard/" + kind + ".json")
			evidence[kind] = statErr == nil
		}
		data, err := json.MarshalIndent(dashboard.Build(cfg, apps, catalog, books, dashboardEnv, moduleCatalog, evidence), "", "  ")
		if err != nil {
			return err
		}
		data = append(data, '\n')
		if err := os.WriteFile(dashboardOutput, data, 0600); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Dashboard data exported to %s\n", dashboardOutput)
		return nil
	},
}

func readOptionalFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return "", nil
	}
	return string(data), err
}

func init() {
	dashboardExportCmd.Flags().StringVar(&dashboardEnv, "env", "", "Export one environment")
	dashboardExportCmd.Flags().BoolVar(&dashboardFleet, "fleet", false, "Export the entire fleet")
	dashboardExportCmd.Flags().StringVar(&dashboardOutput, "output", "dashboard-data.json", "Output JSON path")
	dashboardCmd.AddCommand(dashboardExportCmd)
}
