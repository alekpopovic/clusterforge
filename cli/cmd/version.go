package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version = "0.4.0-rc.1"
	Commit  = "unknown"
	Date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the ClusterForge CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "version: %s\n", Version)
		fmt.Fprintf(cmd.OutOrStdout(), "commit: %s\n", Commit)
		fmt.Fprintf(cmd.OutOrStdout(), "date: %s\n", Date)
		fmt.Fprintf(cmd.OutOrStdout(), "go: %s\n", runtime.Version())
	},
}
