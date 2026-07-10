package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	cfgraph "github.com/textracta/clusterforge/cli/internal/graph"
	cfterraform "github.com/textracta/clusterforge/cli/internal/terraform"
)

var graphStack string
var graphFormat string
var graphOutput string
var graphTerraform bool

func writeGraphOutput(out io.Writer, outputPath string, data []byte) error {
	if outputPath == "" {
		_, err := out.Write(data)
		return err
	}
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("create graph output directory: %w", err)
	}
	if err := os.WriteFile(outputPath, data, 0o644); err != nil {
		return fmt.Errorf("write graph output %s: %w", outputPath, err)
	}
	return nil
}

var graphCmd = &cobra.Command{
	Use:   "graph <env>",
	Short: "Generate a logical or Terraform dependency graph without applying",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		envName := args[0]
		env, ok := cfg.Environments[envName]
		if !ok {
			return fmt.Errorf("environment %q not found", envName)
		}
		if graphFormat != "dot" && graphFormat != "text" {
			return fmt.Errorf("format must be dot or text")
		}
		if !graphTerraform {
			logical, err := cfgraph.Build(envName, graphStack)
			if err != nil {
				return err
			}
			data := []byte(logical.DOT())
			if graphFormat == "text" {
				data = []byte(logical.Text())
			}
			return writeGraphOutput(cmd.OutOrStdout(), graphOutput, data)
		}
		if graphFormat != "dot" {
			return fmt.Errorf("Terraform graph supports only dot format")
		}
		if env.EffectiveLayout() == "stacked" && graphStack == "" {
			return fmt.Errorf("--stack is required for Terraform graph in a stacked environment")
		}
		paths, err := resolveStackPaths(env, graphStack)
		if err != nil {
			return err
		}
		if len(paths) != 1 {
			return fmt.Errorf("Terraform graph requires exactly one work directory")
		}
		binary, err := engineBinary(cfg)
		if err != nil {
			return err
		}
		runner := cfterraform.NewRunner(binary, paths[0], opts.Verbose)
		runner.Stderr = cmd.ErrOrStderr()
		data, err := runner.Graph(cmd.Context())
		if err != nil {
			return err
		}
		return writeGraphOutput(cmd.OutOrStdout(), graphOutput, data)
	},
}

func init() {
	graphCmd.Flags().StringVar(&graphStack, "stack", "", "Limit the logical graph or select a Terraform stack")
	graphCmd.Flags().StringVar(&graphFormat, "format", "dot", "Output format: dot or text")
	graphCmd.Flags().StringVar(&graphOutput, "output", "", "Write graph output to a file")
	graphCmd.Flags().BoolVar(&graphTerraform, "terraform", false, "Run terraform/tofu graph in the selected work directory")
}
