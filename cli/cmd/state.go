package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/alekpopovic/clusterforge/cli/internal/policy"
	cfterraform "github.com/alekpopovic/clusterforge/cli/internal/terraform"
	"github.com/spf13/cobra"
)

var stateStack string
var stateOutput string
var stateAllowRepoOutput bool

var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Run safe read-only Terraform/OpenTofu state operations",
}

var stateListCmd = &cobra.Command{
	Use:   "list <env>",
	Short: "List state addresses",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		envName := args[0]
		runners, err := stateRunners(envName)
		if err != nil {
			return err
		}
		warnState(envName)
		for _, runner := range runners {
			if err := runner.StateList(cmd.Context()); err != nil {
				return err
			}
		}
		return nil
	},
}

var stateShowCmd = &cobra.Command{
	Use:   "show <env> <address>",
	Short: "Show one state object",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		envName := args[0]
		runners, err := stateRunners(envName)
		if err != nil {
			return err
		}
		warnState(envName)
		for _, runner := range runners {
			if err := runner.StateShow(cmd.Context(), args[1]); err != nil {
				return err
			}
		}
		return nil
	},
}

var statePullCmd = &cobra.Command{
	Use:   "pull <env>",
	Short: "Pull state JSON to a required output file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if stateOutput == "" {
			return fmt.Errorf("--output is required")
		}
		if !stateAllowRepoOutput {
			inside, err := pathInsideRepo(stateOutput)
			if err != nil {
				return err
			}
			if inside {
				return fmt.Errorf("refusing to write state inside the repository without --allow-repo-output")
			}
		}
		envName := args[0]
		runners, err := stateRunners(envName)
		if err != nil {
			return err
		}
		if len(runners) != 1 {
			return fmt.Errorf("--stack is required when pulling state from a stacked environment")
		}
		warnState(envName)
		if err := runners[0].StatePull(cmd.Context(), stateOutput); err != nil {
			return err
		}
		printer.Success(fmt.Sprintf("wrote state to %s", stateOutput))
		return nil
	},
}

func stateRunners(envName string) ([]cfterraform.Runner, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}
	env, ok := cfg.Environments[envName]
	if !ok {
		return nil, fmt.Errorf("environment %q not found", envName)
	}
	paths, err := resolveStackPaths(env, stateStack)
	if err != nil {
		return nil, err
	}
	binary, err := engineBinary(cfg)
	if err != nil {
		return nil, err
	}
	runners := make([]cfterraform.Runner, 0, len(paths))
	for _, path := range paths {
		runners = append(runners, cfterraform.NewRunner(binary, path, opts.Verbose))
	}
	return runners, nil
}

func warnState(envName string) {
	printer.Warn("Terraform state can contain sensitive values; handle output carefully")
	if policy.IsProd(envName) {
		printer.Warn("production state inspection should be audited")
	}
}

func pathInsideRepo(path string) (bool, error) {
	absRepo, err := filepath.Abs(".")
	if err != nil {
		return false, err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}
	rel, err := filepath.Rel(absRepo, absPath)
	if err != nil {
		return false, err
	}
	return rel == "." || (!strings.HasPrefix(rel, ".."+string(filepath.Separator)) && rel != ".."), nil
}

func init() {
	for _, command := range []*cobra.Command{stateListCmd, stateShowCmd, statePullCmd} {
		command.Flags().StringVar(&stateStack, "stack", "", "Stack for stacked environments")
	}
	statePullCmd.Flags().StringVar(&stateOutput, "output", "", "Output file for pulled state JSON")
	statePullCmd.Flags().BoolVar(&stateAllowRepoOutput, "allow-repo-output", false, "Allow writing pulled state inside the repository")
	stateCmd.AddCommand(stateListCmd)
	stateCmd.AddCommand(stateShowCmd)
	stateCmd.AddCommand(statePullCmd)
}
