package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/alekpopovic/clusterforge/cli/internal/servicecatalog"
	"github.com/spf13/cobra"
)

var serviceFormat string
var serviceCmd = &cobra.Command{Use: "service", Short: "Inspect the local service catalog"}

func loadServiceCatalog() (servicecatalog.Catalog, error) {
	return servicecatalog.Load(servicecatalog.DefaultPath)
}

var serviceListCmd = &cobra.Command{Use: "list", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	c, err := loadServiceCatalog()
	if err != nil {
		return err
	}
	for _, name := range c.Names() {
		fmt.Fprintln(cmd.OutOrStdout(), name)
	}
	return nil
}}
var serviceShowCmd = &cobra.Command{Use: "show <name>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	c, err := loadServiceCatalog()
	if err != nil {
		return err
	}
	s, ok := c.Services[args[0]]
	if !ok {
		return fmt.Errorf("service %q not found", args[0])
	}
	data, _ := json.MarshalIndent(s, "", "  ")
	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}}
var serviceValidateCmd = &cobra.Command{Use: "validate", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	c, err := loadServiceCatalog()
	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "%d service(s): valid\n", len(c.Services))
	return nil
}}
var serviceExportCmd = &cobra.Command{Use: "export", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	c, err := loadServiceCatalog()
	if err != nil {
		return err
	}
	switch serviceFormat {
	case "json":
		data, _ := c.JSON()
		fmt.Fprintln(cmd.OutOrStdout(), string(data))
	case "markdown":
		return c.WriteMarkdown(cmd.OutOrStdout())
	default:
		return fmt.Errorf("--format must be json or markdown")
	}
	return nil
}}
var serviceGraphCmd = &cobra.Command{Use: "graph", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	if serviceFormat != "dot" {
		return fmt.Errorf("service graph --format must be dot")
	}
	c, err := loadServiceCatalog()
	if err != nil {
		return err
	}
	fmt.Fprint(cmd.OutOrStdout(), c.DOT())
	return nil
}}

func init() {
	serviceExportCmd.Flags().StringVar(&serviceFormat, "format", "json", "json or markdown")
	serviceGraphCmd.Flags().StringVar(&serviceFormat, "format", "dot", "dot")
	serviceCmd.AddCommand(serviceListCmd, serviceShowCmd, serviceValidateCmd, serviceExportCmd, serviceGraphCmd)
}
