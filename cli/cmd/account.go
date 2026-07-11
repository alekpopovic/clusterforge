package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{Use: "account", Short: "Inspect AWS account metadata"}
var accountListCmd = &cobra.Command{Use: "list", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	names := make([]string, 0, len(cfg.AWSAccounts))
	for name := range cfg.AWSAccounts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		account := cfg.AWSAccounts[name]
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\n", name, account.AccountID, account.Region)
	}
	return nil
}}
var accountShowCmd = &cobra.Command{Use: "show <name>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	account, ok := cfg.AWSAccounts[args[0]]
	if !ok {
		return fmt.Errorf("AWS account %q not found", args[0])
	}
	fmt.Fprintf(cmd.OutOrStdout(), "name: %s\naccount_id: %s\nregion: %s\nprofile: %s\nrole_arn: %s\n", args[0], account.AccountID, account.Region, account.Profile, account.RoleARN)
	return nil
}}
var accountDoctorCmd = &cobra.Command{Use: "doctor <name>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	account, ok := cfg.AWSAccounts[args[0]]
	if !ok {
		return fmt.Errorf("AWS account %q not found", args[0])
	}
	fmt.Fprintln(cmd.OutOrStdout(), "PASS\taccount metadata is valid")
	if account.Profile == "" && account.RoleARN == "" {
		fmt.Fprintln(cmd.OutOrStdout(), "WARN\tno profile or deployment role configured; ambient credentials would be used")
	}
	if strings.Contains(strings.ToLower(account.RoleARN), ":root") {
		fmt.Fprintln(cmd.OutOrStdout(), "WARN\troot principal must not be used for deployment")
	}
	for envName, env := range cfg.Environments {
		if (envName == "prod" || envName == "production") && env.Account == args[0] && !cfg.AllowSharedProdAWSAccount {
			for otherName, other := range cfg.Environments {
				if otherName != envName && other.Account == env.Account {
					fmt.Fprintf(cmd.OutOrStdout(), "WARN\tproduction and %s share AWS account %s\n", otherName, args[0])
				}
			}
		}
	}
	return nil
}}

var envDoctorCmd = &cobra.Command{Use: "doctor <env>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	env, ok := cfg.Environments[args[0]]
	if !ok {
		return fmt.Errorf("environment %q not found", args[0])
	}
	fmt.Fprintf(cmd.OutOrStdout(), "PASS\tenvironment path\t%s\n", env.Path)
	if env.Cloud == "aws" {
		if env.Account == "" {
			fmt.Fprintln(cmd.OutOrStdout(), "WARN\tAWS environment has no named account")
		} else {
			account := cfg.AWSAccounts[env.Account]
			fmt.Fprintf(cmd.OutOrStdout(), "PASS\tAWS account\t%s (%s)\n", env.Account, account.AccountID)
		}
	}
	return nil
}}

func init() {
	accountCmd.AddCommand(accountListCmd, accountShowCmd, accountDoctorCmd)
	envCmd.AddCommand(envDoctorCmd)
}
