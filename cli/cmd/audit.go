package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/textracta/clusterforge/cli/internal/audit"
	"github.com/textracta/clusterforge/cli/internal/config"
)

var auditJSON bool
var auditTailLines int
var auditClearYes bool
var auditExportFormat, auditExportSince, auditExportOutput string
var auditRedactInput, auditRedactOutput string

var auditedCommands = map[string]bool{
	"cf project init": true, "cf env create": true, "cf generate": true,
	"cf app add": true, "cf app render": true, "cf app remove": true,
	"cf plan": true, "cf apply": true, "cf destroy": true,
	"cf drift check": true, "cf policy check": true, "cf upgrade apply": true,
}

func auditSettings() (config.Audit, error) {
	cfg, err := config.Load(opts.ConfigPath)
	if err != nil {
		return config.Audit{}, err
	}
	return cfg.Audit, nil
}

func auditArgs(cmd *cobra.Command, positional []string) []string {
	args := append([]string{}, positional...)
	seen := map[string]bool{}
	visit := func(flag *pflag.Flag) {
		if seen[flag.Name] {
			return
		}
		seen[flag.Name] = true
		args = append(args, fmt.Sprintf("--%s=%s", flag.Name, flag.Value.String()))
	}
	cmd.Flags().Visit(visit)
	cmd.InheritedFlags().Visit(visit)
	return audit.Redact(args)
}

func auditContext(cmd *cobra.Command, args []string) (environment, stack string) {
	switch cmd.CommandPath() {
	case "cf plan", "cf apply", "cf destroy", "cf drift check", "cf policy check", "cf env create":
		if len(args) > 0 {
			environment = args[0]
		}
	case "cf app render":
		if flag := cmd.Flags().Lookup("env"); flag != nil {
			environment = flag.Value.String()
		}
	}
	if flag := cmd.Flags().Lookup("stack"); flag != nil {
		stack = flag.Value.String()
	}
	return environment, stack
}

func appendAuditEntry(cmd *cobra.Command, args []string, started time.Time, runErr error) {
	settings, err := auditSettings()
	if err != nil || !settings.Enabled {
		return
	}
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "unknown"
	}
	username := os.Getenv("USER")
	if current, err := user.Current(); err == nil && current.Username != "" {
		username = current.Username
	}
	environment, stack := auditContext(cmd, args)
	result := "success"
	if runErr != nil {
		result = "error"
	}
	entry := audit.Entry{
		Timestamp: time.Now().UTC(), User: username,
		Command: strings.TrimPrefix(cmd.CommandPath(), "cf "), Args: auditArgs(cmd, args),
		WorkingDirectory: cwd, Environment: environment, Stack: stack,
		Result: result, DurationMS: time.Since(started).Milliseconds(), CLIVersion: Version,
	}
	if err := audit.Append(settings.Path, entry); err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "warning: audit log write failed: %v\n", err)
	}
}

func wrapAuditedCommand(command *cobra.Command) {
	if command.RunE == nil || !auditedCommands[command.CommandPath()] {
		return
	}
	original := command.RunE
	command.RunE = func(cmd *cobra.Command, args []string) error {
		started := time.Now()
		err := original(cmd, args)
		appendAuditEntry(cmd, args, started, err)
		return err
	}
}

func installAuditWrappers(command *cobra.Command) {
	wrapAuditedCommand(command)
	for _, child := range command.Commands() {
		installAuditWrappers(child)
	}
}

var auditCmd = &cobra.Command{Use: "audit", Short: "Inspect and manage the local audit log"}

func showAuditEntries(cmd *cobra.Command, limit int) error {
	settings, err := auditSettings()
	if err != nil {
		return err
	}
	entries, err := audit.Read(settings.Path, limit)
	if err != nil {
		return err
	}
	if auditJSON {
		encoder := json.NewEncoder(cmd.OutOrStdout())
		encoder.SetIndent("", "  ")
		return encoder.Encode(entries)
	}
	for _, entry := range entries {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\t%s\t%s\t%s\t%dms\n", entry.Timestamp.Format(time.RFC3339), entry.Result, entry.User, entry.Command, entry.DurationMS)
	}
	return nil
}

var auditShowCmd = &cobra.Command{Use: "show", Short: "Show all local audit entries", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error { return showAuditEntries(cmd, 0) }}
var auditTailCmd = &cobra.Command{Use: "tail", Short: "Show recent local audit entries", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error { return showAuditEntries(cmd, auditTailLines) }}
var auditClearCmd = &cobra.Command{
	Use: "clear", Short: "Clear the local audit log after confirmation", Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !auditClearYes {
			if opts.NonInteractive {
				return fmt.Errorf("audit clear requires --yes in non-interactive mode")
			}
			confirmed, err := newPromptSession(cmd).Bool("clear local audit log", false)
			if err != nil {
				return err
			}
			if !confirmed {
				return fmt.Errorf("audit clear cancelled")
			}
		}
		settings, err := auditSettings()
		if err != nil {
			return err
		}
		if err := audit.Clear(settings.Path); err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), "audit log cleared")
		return nil
	},
}

var auditExportCmd = &cobra.Command{
	Use: "export", Short: "Export redacted local audit entries", Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		settings, err := auditSettings()
		if err != nil {
			return err
		}
		entries, err := audit.Read(settings.Path, 0)
		if err != nil {
			return err
		}
		if auditExportSince != "" {
			duration, parseErr := time.ParseDuration(auditExportSince)
			if parseErr != nil || duration < 0 {
				return fmt.Errorf("invalid --since duration %q", auditExportSince)
			}
			entries = audit.FilterSince(entries, time.Now().UTC().Add(-duration))
		}
		file, err := os.OpenFile(auditExportOutput, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
		if err != nil {
			return err
		}
		if err := audit.Export(file, entries, auditExportFormat); err != nil {
			file.Close()
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Exported %d audit entries to %s\n", len(entries), auditExportOutput)
		return nil
	},
}

var auditRedactCmd = &cobra.Command{
	Use: "redact", Short: "Create a redacted copy of a JSONL audit log", Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if auditRedactInput == "" || auditRedactOutput == "" {
			return fmt.Errorf("--input and --output are required")
		}
		if auditRedactInput == auditRedactOutput {
			return fmt.Errorf("--input and --output must differ")
		}
		file, err := os.OpenFile(auditRedactOutput, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
		if err != nil {
			return err
		}
		if err := audit.RedactFile(auditRedactInput, file); err != nil {
			file.Close()
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Redacted audit log written to %s\n", auditRedactOutput)
		return nil
	},
}

func init() {
	auditShowCmd.Flags().BoolVar(&auditJSON, "json", false, "Print audit entries as JSON")
	auditTailCmd.Flags().BoolVar(&auditJSON, "json", false, "Print audit entries as JSON")
	auditTailCmd.Flags().IntVar(&auditTailLines, "lines", 20, "Number of recent entries to show")
	auditClearCmd.Flags().BoolVar(&auditClearYes, "yes", false, "Confirm audit log deletion")
	auditExportCmd.Flags().StringVar(&auditExportFormat, "format", "jsonl", "Export format: jsonl, json, or csv")
	auditExportCmd.Flags().StringVar(&auditExportSince, "since", "", "Only include entries from this duration (for example 24h)")
	auditExportCmd.Flags().StringVar(&auditExportOutput, "output", "audit.jsonl", "Output file")
	auditRedactCmd.Flags().StringVar(&auditRedactInput, "input", "", "Input JSONL audit log")
	auditRedactCmd.Flags().StringVar(&auditRedactOutput, "output", "", "Output redacted JSONL file")
	auditCmd.AddCommand(auditShowCmd, auditTailCmd, auditClearCmd, auditExportCmd, auditRedactCmd)
}
