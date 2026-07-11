package audit

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Timestamp        time.Time `json:"timestamp"`
	User             string    `json:"user,omitempty"`
	Command          string    `json:"command"`
	Args             []string  `json:"args,omitempty"`
	WorkingDirectory string    `json:"working_directory"`
	Environment      string    `json:"environment,omitempty"`
	Stack            string    `json:"stack,omitempty"`
	Result           string    `json:"result"`
	DurationMS       int64     `json:"duration_ms"`
	CLIVersion       string    `json:"cli_version"`
}

func Redact(args []string) []string {
	redacted := append([]string{}, args...)
	redactNext := false
	for index, arg := range redacted {
		if redactNext {
			redacted[index] = "[REDACTED]"
			redactNext = false
			continue
		}
		if !strings.HasPrefix(arg, "-") {
			continue
		}
		name, _, hasValue := strings.Cut(arg, "=")
		if !sensitiveFlag(name) {
			continue
		}
		if hasValue {
			redacted[index] = name + "=[REDACTED]"
		} else {
			redactNext = true
		}
	}
	return redacted
}

func sensitiveFlag(flag string) bool {
	lower := strings.ToLower(flag)
	for _, word := range []string{"password", "token", "secret", "key"} {
		if strings.Contains(lower, word) {
			return true
		}
	}
	return false
}

func Append(path string, entry Entry) error {
	if path == "" {
		return fmt.Errorf("audit path is required")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create audit directory: %w", err)
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return fmt.Errorf("open audit log: %w", err)
	}
	defer file.Close()
	data, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("encode audit entry: %w", err)
	}
	if _, err := file.Write(append(data, '\n')); err != nil {
		return fmt.Errorf("write audit entry: %w", err)
	}
	return nil
}

func Read(path string, limit int) ([]Entry, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}
		return nil, fmt.Errorf("open audit log: %w", err)
	}
	defer file.Close()
	entries := []Entry{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry Entry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			return nil, fmt.Errorf("parse audit log entry: %w", err)
		}
		entries = append(entries, entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read audit log: %w", err)
	}
	if limit > 0 && len(entries) > limit {
		entries = entries[len(entries)-limit:]
	}
	return entries, nil
}

func Clear(path string) error {
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("clear audit log: %w", err)
	}
	return nil
}

// FilterSince returns entries at or after the cutoff and redacts sensitive
// arguments again so older logs receive the current redaction rules on export.
func FilterSince(entries []Entry, cutoff time.Time) []Entry {
	filtered := make([]Entry, 0, len(entries))
	for _, entry := range entries {
		if !cutoff.IsZero() && entry.Timestamp.Before(cutoff) {
			continue
		}
		entry.Args = Redact(entry.Args)
		filtered = append(filtered, entry)
	}
	return filtered
}

func Export(writer io.Writer, entries []Entry, format string) error {
	entries = FilterSince(entries, time.Time{})
	switch format {
	case "jsonl":
		encoder := json.NewEncoder(writer)
		for _, entry := range entries {
			if err := encoder.Encode(entry); err != nil {
				return fmt.Errorf("encode JSONL audit entry: %w", err)
			}
		}
		return nil
	case "json":
		encoder := json.NewEncoder(writer)
		encoder.SetIndent("", "  ")
		return encoder.Encode(entries)
	case "csv":
		csvWriter := csv.NewWriter(writer)
		if err := csvWriter.Write([]string{"timestamp", "user", "command", "args", "working_directory", "environment", "stack", "result", "duration_ms", "cli_version"}); err != nil {
			return err
		}
		for _, entry := range entries {
			args, _ := json.Marshal(entry.Args)
			record := []string{entry.Timestamp.Format(time.RFC3339Nano), entry.User, entry.Command, string(args), entry.WorkingDirectory, entry.Environment, entry.Stack, entry.Result, strconv.FormatInt(entry.DurationMS, 10), entry.CLIVersion}
			if err := csvWriter.Write(record); err != nil {
				return err
			}
		}
		csvWriter.Flush()
		return csvWriter.Error()
	default:
		return fmt.Errorf("unsupported audit export format %q (use jsonl, json, or csv)", format)
	}
}

// RedactFile rewrites a JSONL audit log without modifying its source.
func RedactFile(inputPath string, writer io.Writer) error {
	entries, err := Read(inputPath, 0)
	if err != nil {
		return err
	}
	return Export(writer, entries, "jsonl")
}
