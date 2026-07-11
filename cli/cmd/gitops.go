package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cfapp "github.com/textracta/clusterforge/cli/internal/app"
	"github.com/textracta/clusterforge/cli/internal/gitops"
)

var gitopsCluster string
var gitopsCmd = &cobra.Command{Use: "gitops", Short: "Render cross-cluster GitOps manifests"}
var gitopsRenderCmd = &cobra.Command{Use: "render", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
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
	data, err := gitops.Render(cfg.GitOps, apps, gitopsCluster)
	if err != nil {
		return err
	}
	_, err = cmd.OutOrStdout().Write(data)
	return err
}}
var gitopsClustersCmd = &cobra.Command{Use: "clusters", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	if err := gitops.Validate(cfg.GitOps); err != nil {
		return err
	}
	for _, cluster := range cfg.GitOps.Clusters {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\n", cluster.Name, cluster.Environment)
	}
	return nil
}}
var gitopsAppsCmd = &cobra.Command{Use: "apps", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	names, err := cfapp.List(".")
	if err != nil {
		return err
	}
	for _, name := range names {
		fmt.Fprintln(cmd.OutOrStdout(), name)
	}
	return nil
}}

func init() {
	gitopsRenderCmd.Flags().StringVar(&gitopsCluster, "cluster", "", "Render one configured cluster")
	gitopsCmd.AddCommand(gitopsRenderCmd, gitopsClustersCmd, gitopsAppsCmd)
}
