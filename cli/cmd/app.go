package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cfapp "github.com/textracta/clusterforge/cli/internal/app"
)

var appAddOptions cfapp.AddOptions
var appRenderEnv string

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Manage ClusterForge app manifests",
}

var appAddCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Create an app manifest",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cfapp.Add(".", args[0], appAddOptions)
		if err != nil {
			return err
		}
		printer.Success(fmt.Sprintf("created app manifest %s", path))
		return nil
	},
}

var appListCmd = &cobra.Command{
	Use:   "list",
	Short: "List app manifests",
	RunE: func(cmd *cobra.Command, args []string) error {
		apps, err := cfapp.List(".")
		if err != nil {
			return err
		}
		for _, app := range apps {
			fmt.Fprintln(cmd.OutOrStdout(), app)
		}
		return nil
	},
}

var appRenderCmd = &cobra.Command{
	Use:   "render <name>",
	Short: "Render an app manifest into an environment Terraform module call",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if appRenderEnv == "" {
			return fmt.Errorf("--env is required")
		}
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		env, ok := cfg.Environments[appRenderEnv]
		if !ok {
			return fmt.Errorf("environment %q not found", appRenderEnv)
		}
		manifest, err := cfapp.Load(cfapp.ManifestPath(".", args[0]))
		if err != nil {
			return err
		}
		outPath, err := cfapp.Render(".", appRenderEnv, env, manifest)
		if err != nil {
			return err
		}
		printer.Success(fmt.Sprintf("rendered app %q to %s", args[0], outPath))
		return nil
	},
}

var appRemoveCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove an app manifest",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cfapp.Remove(".", args[0]); err != nil {
			return err
		}
		printer.Success(fmt.Sprintf("removed app manifest %q", args[0]))
		return nil
	},
}

func init() {
	appAddCmd.Flags().StringVar(&appAddOptions.Image, "image", "", "Container image for the app")
	appAddCmd.Flags().IntVar(&appAddOptions.Port, "port", 8080, "Container port for the default http port")
	appAddCmd.Flags().IntVar(&appAddOptions.Replicas, "replicas", 1, "Replica count")
	appAddCmd.Flags().StringVar(&appAddOptions.Host, "host", "", "Optional ingress host")
	appAddCmd.Flags().StringVar(&appAddOptions.Type, "type", "web", "App type")
	appAddCmd.Flags().BoolVar(&appAddOptions.Force, "force", false, "Overwrite an existing app manifest")

	appRenderCmd.Flags().StringVar(&appRenderEnv, "env", "", "Environment to render into")

	appCmd.AddCommand(appAddCmd)
	appCmd.AddCommand(appListCmd)
	appCmd.AddCommand(appRenderCmd)
	appCmd.AddCommand(appRemoveCmd)
}
