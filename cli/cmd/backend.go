package cmd

import (
	"fmt"

	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"github.com/alekpopovic/clusterforge/cli/internal/policy"
	"github.com/spf13/cobra"
)

var backendType string
var backendBucket string
var backendRegion string
var backendDynamoDBTable string
var backendKeyPrefix string

var backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Manage Terraform/OpenTofu backend configuration",
}

var backendConfigureCmd = &cobra.Command{
	Use:   "configure <env>",
	Short: "Configure backend settings for an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		envName := args[0]
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		if _, ok := cfg.Environments[envName]; !ok {
			return fmt.Errorf("environment %q not found", envName)
		}

		backend := cfg.BackendFor(envName)
		if backendType != "" {
			backend.Type = backendType
		}
		if backend.Type == "" {
			backend.Type = "local"
		}
		if backendBucket != "" {
			backend.Bucket = backendBucket
		}
		if backendRegion != "" {
			backend.Region = backendRegion
		}
		if backendDynamoDBTable != "" {
			backend.DynamoDBTable = backendDynamoDBTable
		}
		if backendKeyPrefix != "" {
			backend.KeyPrefix = backendKeyPrefix
		}
		if err := backend.Validate(envName); err != nil {
			return err
		}
		if cfg.Backends == nil {
			cfg.Backends = map[string]config.Backend{}
		}
		cfg.Backends[envName] = backend
		if policy.IsProd(envName) && backend.EffectiveType() == "local" {
			printer.Warn(fmt.Sprintf("prod environment %q is configured with local backend", envName))
		}
		if err := cfg.Save(opts.ConfigPath); err != nil {
			return err
		}
		printer.Success(fmt.Sprintf("configured %s backend for %s", backend.EffectiveType(), envName))
		return nil
	},
}

var backendShowCmd = &cobra.Command{
	Use:   "show <env>",
	Short: "Show backend settings for an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		envName := args[0]
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		if _, ok := cfg.Environments[envName]; !ok {
			return fmt.Errorf("environment %q not found", envName)
		}
		backend := cfg.BackendFor(envName)
		fmt.Fprintf(cmd.OutOrStdout(), "type: %s\n", backend.EffectiveType())
		if backend.Bucket != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "bucket: %s\n", backend.Bucket)
		}
		if backend.Region != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "region: %s\n", backend.Region)
		}
		if backend.DynamoDBTable != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "dynamodb_table: %s\n", backend.DynamoDBTable)
		}
		if backend.KeyPrefix != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "key_prefix: %s\n", backend.KeyPrefix)
		}
		if policy.IsProd(envName) && backend.EffectiveType() == "local" {
			printer.Warn(fmt.Sprintf("prod environment %q is configured with local backend", envName))
		}
		return nil
	},
}

func init() {
	backendConfigureCmd.Flags().StringVar(&backendType, "backend", "", "Backend type: local, s3, azurerm, or gcs")
	backendConfigureCmd.Flags().StringVar(&backendBucket, "bucket", "", "S3 backend bucket name")
	backendConfigureCmd.Flags().StringVar(&backendRegion, "region", "", "Backend region")
	backendConfigureCmd.Flags().StringVar(&backendDynamoDBTable, "dynamodb-table", "", "S3 backend DynamoDB lock table")
	backendConfigureCmd.Flags().StringVar(&backendKeyPrefix, "key-prefix", "", "State key prefix")

	backendCmd.AddCommand(backendConfigureCmd)
	backendCmd.AddCommand(backendShowCmd)
}
