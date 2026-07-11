package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/secrets"
)

var secretsApp string
var secretsCmd = &cobra.Command{Use: "secrets", Short: "Inspect secret references without reading secret values"}

func discoverSecretReferences(environment string) ([]secrets.Reference, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}
	if environment == "" {
		return secrets.Discover(".", "", secretsApp)
	}
	env, ok := cfg.Environments[environment]
	if !ok {
		return nil, fmt.Errorf("environment %q not found", environment)
	}
	return secrets.Discover(".", env.Path, secretsApp)
}

var secretsCheckCmd = &cobra.Command{Use: "check <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	refs, err := discoverSecretReferences(args[0])
	if err != nil {
		return err
	}
	if len(refs) == 0 {
		fmt.Fprintln(cmd.OutOrStdout(), "No local secret references found; manual verification is required.")
		return nil
	}
	fmt.Fprintf(cmd.OutOrStdout(), "%d secret reference(s) found; values were not read.\n", len(refs))
	return nil
}}

var secretsReferencesCmd = &cobra.Command{Use: "references [env]", Args: cobra.MaximumNArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	environment := ""
	if len(args) == 1 {
		environment = args[0]
	}
	if environment == "" && secretsApp == "" {
		return fmt.Errorf("provide <env> or --app")
	}
	refs, err := discoverSecretReferences(environment)
	if err != nil {
		return err
	}
	for _, ref := range refs {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s", ref.Kind, ref.Name, ref.Source)
		if ref.Key != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "\tkey=%s", ref.Key)
		}
		if ref.App != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "\tapp=%s", ref.App)
		}
		fmt.Fprintln(cmd.OutOrStdout())
	}
	return nil
}}

var secretsRotationPlanCmd = &cobra.Command{Use: "rotation-plan <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	refs, err := discoverSecretReferences(args[0])
	if err != nil {
		return err
	}
	for index, step := range secrets.RotationPlan(args[0], refs) {
		fmt.Fprintf(cmd.OutOrStdout(), "%d. %s\n", index+1, step)
	}
	return nil
}}

func init() {
	secretsReferencesCmd.Flags().StringVar(&secretsApp, "app", "", "Only references for one app (environment files remain included)")
	secretsCmd.AddCommand(secretsCheckCmd, secretsReferencesCmd, secretsRotationPlanCmd)
}
