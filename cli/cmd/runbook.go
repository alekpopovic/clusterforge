package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/runbooks"
	"strings"
)

var runbookCmd = &cobra.Command{Use: "runbook", Short: "Discover local operational runbooks"}

func allRunbooks() ([]runbooks.Runbook, error) {
	return runbooks.Discover("docs/incident-response", "docs/dr")
}
func printRunbooks(cmd *cobra.Command, books []runbooks.Runbook) {
	for _, book := range books {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\n", book.Name, book.Path, book.Summary)
	}
}

var runbookListCmd = &cobra.Command{Use: "list", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
	books, err := allRunbooks()
	if err != nil {
		return err
	}
	printRunbooks(cmd, books)
	return nil
}}
var runbookShowCmd = &cobra.Command{Use: "show <name>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	books, err := allRunbooks()
	if err != nil {
		return err
	}
	book, err := runbooks.Find(books, args[0])
	if err != nil {
		return err
	}
	fmt.Fprint(cmd.OutOrStdout(), book.Content)
	return nil
}}
var runbookSearchCmd = &cobra.Command{Use: "search <query>", Args: cobra.MinimumNArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	books, err := allRunbooks()
	if err != nil {
		return err
	}
	printRunbooks(cmd, runbooks.Search(books, strings.Join(args, " ")))
	return nil
}}
var runbookOpenCmd = &cobra.Command{Use: "open <name>", Short: "Print a runbook path without starting a process", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	books, err := allRunbooks()
	if err != nil {
		return err
	}
	book, err := runbooks.Find(books, args[0])
	if err != nil {
		return err
	}
	fmt.Fprintln(cmd.OutOrStdout(), book.Path)
	return nil
}}

func init() { runbookCmd.AddCommand(runbookListCmd, runbookShowCmd, runbookSearchCmd, runbookOpenCmd) }
