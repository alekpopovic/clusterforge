package cmd

import (
	"fmt"
	cfapp "github.com/alekpopovic/clusterforge/cli/internal/app"
	"github.com/alekpopovic/clusterforge/cli/internal/backstage"
	"github.com/alekpopovic/clusterforge/cli/internal/servicecatalog"
	"github.com/spf13/cobra"
	"os"
)

var backstageApp, backstageEnv, backstageOutput string
var backstageCmd = &cobra.Command{Use: "backstage", Short: "Generate Backstage catalog YAML"}
var backstageGenerateCmd = &cobra.Command{Use: "generate", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	manifests := map[string]cfapp.Manifest{}
	names, err := cfapp.List(".")
	if err != nil {
		return err
	}
	for _, name := range names {
		if backstageApp != "" && name != backstageApp {
			continue
		}
		manifest, err := cfapp.Load(cfapp.ManifestPath(".", name))
		if err != nil {
			return err
		}
		manifests[name] = manifest
	}
	if backstageApp != "" {
		if _, ok := manifests[backstageApp]; !ok {
			return fmt.Errorf("app %q not found", backstageApp)
		}
	}
	catalog := servicecatalog.Catalog{}
	if loaded, loadErr := servicecatalog.Load(servicecatalog.DefaultPath); loadErr == nil {
		catalog = loaded
	}
	data, err := backstage.GenerateWithServices(cfg, manifests, catalog, backstageApp, backstageEnv)
	if err != nil {
		return err
	}
	if backstageOutput != "" {
		return os.WriteFile(backstageOutput, data, 0644)
	}
	_, err = cmd.OutOrStdout().Write(data)
	return err
}}

func init() {
	backstageGenerateCmd.Flags().StringVar(&backstageApp, "app", "", "Only one app")
	backstageGenerateCmd.Flags().StringVar(&backstageEnv, "env", "", "Only one environment")
	backstageGenerateCmd.Flags().StringVar(&backstageOutput, "output", "", "Output YAML path")
	backstageCmd.AddCommand(backstageGenerateCmd)
}
