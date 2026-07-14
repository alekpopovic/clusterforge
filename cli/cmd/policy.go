package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	cfapp "github.com/alekpopovic/clusterforge/cli/internal/app"
	cfenvironment "github.com/alekpopovic/clusterforge/cli/internal/environment"
	"github.com/alekpopovic/clusterforge/cli/internal/policyengine"
	cfterraform "github.com/alekpopovic/clusterforge/cli/internal/terraform"
	"github.com/spf13/cobra"
)

var policyPlanFile, policyPack, policyStack, policyApp, policyFormat string
var policyCheckJSON bool

var policyCmd = &cobra.Command{Use: "policy", Short: "Evaluate unified ClusterForge policies"}

var policyListCmd = &cobra.Command{Use: "list", Short: "List built-in policies", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	for _, policy := range policyengine.BuiltIns() {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%s\n", policy.ID, policy.Severity, policy.Scope, policy.Title)
	}
	return nil
}}

var policyShowCmd = &cobra.Command{Use: "show <id>", Short: "Show one built-in policy", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	for _, policy := range policyengine.BuiltIns() {
		if policy.ID == args[0] {
			data, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Fprintln(cmd.OutOrStdout(), string(data))
			return nil
		}
	}
	return fmt.Errorf("policy %q not found", args[0])
}}

var policyCheckCmd = &cobra.Command{
	Use: "check [env]", Short: "Evaluate repository, environment, plan, or app policies", Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		format := policyFormat
		if policyCheckJSON {
			format = "json"
		}
		if format != "table" && format != "json" && format != "sarif" {
			return fmt.Errorf("--format must be table, json, or sarif")
		}
		cfg, err := loadConfig()
		if err != nil {
			return err
		}
		input := policyengine.Input{RequirePlanFile: cfg.Policies.RequirePlanFileForApply, BlockProdDestroy: cfg.Policies.BlockDestroyInProd}
		tracked, err := exec.Command("git", "ls-files").Output()
		if err == nil {
			input.TrackedFiles = strings.Fields(string(tracked))
		}
		if len(args) == 1 {
			input.Environment = args[0]
			input.Production = cfenvironment.IsProduction(args[0])
			env, ok := cfg.Environments[args[0]]
			if !ok {
				return fmt.Errorf("environment %q not found", args[0])
			}
			input.BackendType = cfg.BackendFor(args[0]).EffectiveType()
			root := env.Path
			if policyStack != "" {
				paths, err := env.StackPaths(policyStack)
				if err != nil {
					return err
				}
				root = paths[0]
			}
			input.TerraformFiles, err = readPolicyFiles(root)
			if err != nil {
				return err
			}
			if policyPlanFile != "" {
				binary, err := engineBinary(cfg)
				if err != nil {
					return err
				}
				data, err := cfterraform.NewRunner(binary, env.Path, opts.Verbose).ShowPlanJSON(cmd.Context(), policyPlanFile)
				if err != nil {
					return err
				}
				input.TerraformFiles[policyPlanFile+".json"] = data
			}
		}
		if policyApp != "" {
			path := cfapp.ManifestPath(".", policyApp)
			data, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read app manifest: %w", err)
			}
			manifest, err := cfapp.Load(path)
			if err != nil {
				return err
			}
			input.AppPath, input.AppYAML, input.Image, input.IngressEnabled = path, data, manifest.Image, manifest.Ingress.Enabled
		}
		selectedPack := defaultString(policyPack, "baseline")
		overrides, err := policyengine.LoadOverrides(filepath.Join("policies", "packs", selectedPack), filepath.Join(".clusterforge", "policies"))
		if err != nil {
			return err
		}
		result := policyengine.Evaluate(input, policyengine.Options{Pack: selectedPack, Overrides: overrides})
		if err := writePolicyResult(cmd, format, result); err != nil {
			return err
		}
		if result.Blocked {
			return commandExitError{code: 2, message: "policy check blocked"}
		}
		return nil
	},
}

func readPolicyFiles(root string) (map[string][]byte, error) {
	result := map[string][]byte{}
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() && (entry.Name() == ".terraform" || entry.Name() == ".git") {
			return filepath.SkipDir
		}
		if entry.IsDir() || (filepath.Ext(path) != ".tf" && filepath.Ext(path) != ".json") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		result[path] = data
		return nil
	})
	if os.IsNotExist(err) {
		return result, nil
	}
	return result, err
}

func writePolicyResult(cmd *cobra.Command, format string, result policyengine.Result) error {
	switch format {
	case "json":
		return json.NewEncoder(cmd.OutOrStdout()).Encode(result)
	case "sarif":
		data, err := policyengine.SARIF(result)
		if err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), string(data))
		return nil
	default:
		fmt.Fprintln(cmd.OutOrStdout(), "POLICY\tSEVERITY\tACTION\tSCOPE\tLOCATION\tMESSAGE")
		for _, finding := range result.Findings {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%s\t%s\t%s\n", finding.PolicyID, finding.Severity, finding.Action, finding.Scope, finding.Location, finding.Message)
		}
		if len(result.Findings) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "PASS\t-\t-\t-\t-\tno findings")
		}
		return nil
	}
}

func init() {
	policyCheckCmd.Flags().StringVar(&policyPlanFile, "plan-file", "", "Existing plan file to inspect")
	policyCheckCmd.Flags().StringVar(&policyPack, "pack", "baseline", "Policy pack: baseline or production")
	policyCheckCmd.Flags().StringVar(&policyStack, "stack", "", "Environment stack to inspect")
	policyCheckCmd.Flags().StringVar(&policyApp, "app", "", "App manifest name to inspect")
	policyCheckCmd.Flags().StringVar(&policyFormat, "format", "table", "Output format: table, json, or sarif")
	policyCheckCmd.Flags().BoolVar(&policyCheckJSON, "json", false, "Alias for --format json")
	policyCmd.AddCommand(policyListCmd, policyShowCmd, policyCheckCmd)
}
