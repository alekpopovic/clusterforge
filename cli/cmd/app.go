package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	cfapp "github.com/textracta/clusterforge/cli/internal/app"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

var appAddOptions cfapp.AddOptions
var appRenderEnv string
var appRenderTemplatePack string
var appValidateAll bool
var appListJSON bool

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Manage ClusterForge app manifests",
}

var appAddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Create an app manifest",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		prompts := newPromptSession(cmd)
		name, err := requireValueWithPrompt(optionalArg(args, 0), "app name", prompts)
		if err != nil {
			return err
		}
		options := appAddOptions
		if strings.TrimSpace(options.Image) == "" {
			if opts.NonInteractive {
				return fmt.Errorf("--image is required in non-interactive mode")
			}
			options.Type, err = prompts.String("app type", defaultString(options.Type, "web"))
			if err != nil {
				return err
			}
			options.Image, err = prompts.String("image", "")
			if err != nil {
				return err
			}
			options.Port, err = prompts.Int("container port", options.Port)
			if err != nil {
				return err
			}
			options.Replicas, err = prompts.Int("replicas", options.Replicas)
			if err != nil {
				return err
			}
			options.Host, err = prompts.String("ingress host", options.Host)
			if err != nil {
				return err
			}
			options.Autoscaling, err = prompts.Bool("enable autoscaling", false)
			if err != nil {
				return err
			}
		}
		manifest := cfapp.NewManifest(name, options)
		printer.Info(fmt.Sprintf("app: %s", manifest.Name))
		printer.Info(fmt.Sprintf("type: %s", manifest.Type))
		printer.Info(fmt.Sprintf("image: %s", manifest.Image))
		printer.Info(fmt.Sprintf("replicas: %d", manifest.Replicas))
		if len(manifest.Ports) > 0 {
			printer.Info(fmt.Sprintf("port: %d", manifest.Ports[0].ContainerPort))
		}
		if manifest.Ingress.Enabled {
			printer.Info(fmt.Sprintf("ingress host: %s", manifest.Ingress.Host))
		}
		printer.Info(fmt.Sprintf("autoscaling: %t", manifest.Autoscaling.Enabled))

		path, err := cfapp.Add(".", name, options)
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
		if appListJSON {
			return ui.WriteJSON(cmd.OutOrStdout(), appListResponse{Apps: apps})
		}
		for _, app := range apps {
			fmt.Fprintln(cmd.OutOrStdout(), app)
		}
		return nil
	},
}

type appListResponse struct {
	Apps []string `json:"apps"`
}

var appValidateCmd = &cobra.Command{
	Use:           "validate [name]",
	Short:         "Validate app manifest files",
	Args:          cobra.MaximumNArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if appValidateAll && len(args) > 0 {
			return fmt.Errorf("use either app validate <name> or app validate --all, not both")
		}
		if !appValidateAll && len(args) == 0 {
			return fmt.Errorf("app name is required unless --all is set")
		}

		names := args
		if appValidateAll {
			apps, err := cfapp.List(".")
			if err != nil {
				return err
			}
			names = apps
		}

		var failed bool
		for _, name := range names {
			path := cfapp.ManifestPath(".", name)
			if err := cfapp.ValidateFile(path); err != nil {
				failed = true
				for _, line := range strings.Split(err.Error(), "\n") {
					if strings.TrimSpace(line) == "" {
						continue
					}
					fmt.Fprintf(cmd.ErrOrStderr(), "%s: %s\n", path, line)
				}
				continue
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s: ok\n", path)
		}
		if failed {
			return fmt.Errorf("app manifest validation failed")
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
		if appRenderTemplatePack != "" {
			found := false
			for _, pack := range cfg.TemplatePacks {
				if pack.Name == appRenderTemplatePack {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("template pack %q not found in %s", appRenderTemplatePack, opts.ConfigPath)
			}
			printer.Warn("app template pack overrides are not executed; built-in app renderer is used")
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
	appAddCmd.Flags().BoolVar(&appAddOptions.Autoscaling, "autoscaling", false, "Enable default autoscaling settings")
	appAddCmd.Flags().BoolVar(&appAddOptions.Force, "force", false, "Overwrite an existing app manifest")

	appRenderCmd.Flags().StringVar(&appRenderEnv, "env", "", "Environment to render into")
	appRenderCmd.Flags().StringVar(&appRenderTemplatePack, "template-pack", "", "Local template pack name to validate before rendering")
	appValidateCmd.Flags().BoolVar(&appValidateAll, "all", false, "Validate all app manifests")
	appListCmd.Flags().BoolVar(&appListJSON, "json", false, "Print app manifests as JSON")

	appCmd.AddCommand(appAddCmd)
	appCmd.AddCommand(appListCmd)
	appCmd.AddCommand(appValidateCmd)
	appCmd.AddCommand(appRenderCmd)
	appCmd.AddCommand(appRemoveCmd)
}
