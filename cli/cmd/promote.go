package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	cfapp "github.com/textracta/clusterforge/cli/internal/app"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

var promoteFrom string
var promoteTo string
var promoteJSON bool

var promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Compare environments before Git-based promotion",
}

var promotePlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "Plan a promotion by comparing source and target",
	RunE:  runPromotionCompare,
}

var promoteDiffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show promotion differences between environments",
	RunE:  runPromotionCompare,
}

func runPromotionCompare(cmd *cobra.Command, args []string) error {
	if promoteFrom == "" || promoteTo == "" {
		return fmt.Errorf("--from and --to are required")
	}
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	fromEnv, fromOK := cfg.Environments[promoteFrom]
	toEnv, toOK := cfg.Environments[promoteTo]
	if !fromOK || !toOK {
		return fmt.Errorf("both source and target environments must exist")
	}
	result := promotionResult{
		From:        promoteFrom,
		To:          promoteTo,
		Differences: []promotionDiff{},
	}
	if fromEnv.Cloud != toEnv.Cloud || fromEnv.Orchestrator != toEnv.Orchestrator || fromEnv.Region != toEnv.Region {
		result.Differences = append(result.Differences, promotionDiff{
			Type: "environment",
			Name: "config",
			Detail: fmt.Sprintf("source %s/%s/%s differs from target %s/%s/%s",
				fromEnv.Cloud, fromEnv.Orchestrator, fromEnv.Region,
				toEnv.Cloud, toEnv.Orchestrator, toEnv.Region),
		})
	}
	apps, err := cfapp.List(".")
	if err != nil {
		return err
	}
	for _, name := range apps {
		manifest, err := cfapp.Load(cfapp.ManifestPath(".", name))
		if err != nil {
			return err
		}
		if manifest.Ingress.Enabled && strings.Contains(strings.ToLower(manifest.Ingress.Host), promoteFrom) && !strings.Contains(strings.ToLower(manifest.Ingress.Host), promoteTo) {
			result.Differences = append(result.Differences, promotionDiff{Type: "ingress", Name: name, Detail: "ingress host appears source-specific"})
		}
		if manifest.Image == "" {
			result.Differences = append(result.Differences, promotionDiff{Type: "image", Name: name, Detail: "image is empty"})
		}
	}
	result.Differences = append(result.Differences, compareGeneratedFiles(fromEnv.Path, toEnv.Path)...)
	if promoteJSON {
		return ui.WriteJSON(cmd.OutOrStdout(), result)
	}
	if len(result.Differences) == 0 {
		fmt.Fprintf(cmd.OutOrStdout(), "no promotion differences detected between %s and %s\n", promoteFrom, promoteTo)
		return nil
	}
	fmt.Fprintf(cmd.OutOrStdout(), "promotion differences between %s and %s:\n", promoteFrom, promoteTo)
	for _, diff := range result.Differences {
		fmt.Fprintf(cmd.OutOrStdout(), "- [%s] %s: %s\n", diff.Type, diff.Name, diff.Detail)
	}
	return nil
}

func compareGeneratedFiles(fromPath, toPath string) []promotionDiff {
	diffs := []promotionDiff{}
	for _, file := range []string{"main.tf", "variables.tf", "outputs.tf", "providers.tf", "versions.tf"} {
		fromData, fromErr := os.ReadFile(filepath.Join(fromPath, file))
		toData, toErr := os.ReadFile(filepath.Join(toPath, file))
		if fromErr != nil || toErr != nil {
			continue
		}
		if string(fromData) != string(toData) {
			diffs = append(diffs, promotionDiff{Type: "terraform", Name: file, Detail: "generated file content differs"})
		}
	}
	return diffs
}

type promotionResult struct {
	From        string          `json:"from"`
	To          string          `json:"to"`
	Differences []promotionDiff `json:"differences"`
}

type promotionDiff struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func init() {
	for _, command := range []*cobra.Command{promotePlanCmd, promoteDiffCmd} {
		command.Flags().StringVar(&promoteFrom, "from", "", "Source environment")
		command.Flags().StringVar(&promoteTo, "to", "", "Target environment")
		command.Flags().BoolVar(&promoteJSON, "json", false, "Print comparison as JSON")
	}
	promoteCmd.AddCommand(promotePlanCmd)
	promoteCmd.AddCommand(promoteDiffCmd)
}
