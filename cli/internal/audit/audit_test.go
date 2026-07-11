package audit

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestAppendAndReadEntry(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".cf", "audit.log")
	want := Entry{Timestamp: time.Now().UTC(), Command: "plan", Result: "success", CLIVersion: "test"}
	if err := Append(path, want); err != nil {
		t.Fatal(err)
	}
	entries, err := Read(path, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].Command != "plan" {
		t.Fatalf("entries = %#v", entries)
	}
}

func TestExportJSONLAndCSV(t *testing.T) {
	entries := []Entry{{Timestamp: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC), Command: "plan", Args: []string{"--token=secret"}, Result: "success"}}
	var jsonl bytes.Buffer
	if err := Export(&jsonl, entries, "jsonl"); err != nil {
		t.Fatal(err)
	}
	var decoded Entry
	if err := json.Unmarshal(jsonl.Bytes(), &decoded); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(jsonl.String(), "secret") || decoded.Args[0] != "--token=[REDACTED]" {
		t.Fatalf("JSONL was not redacted: %s", jsonl.String())
	}
	var csvData bytes.Buffer
	if err := Export(&csvData, entries, "csv"); err != nil {
		t.Fatal(err)
	}
	records, err := csv.NewReader(&csvData).ReadAll()
	if err != nil || len(records) != 2 || records[1][2] != "plan" {
		t.Fatalf("records=%v err=%v", records, err)
	}
}

func TestFilterSince(t *testing.T) {
	now := time.Now().UTC()
	entries := []Entry{{Timestamp: now.Add(-2 * time.Hour)}, {Timestamp: now.Add(-30 * time.Minute)}}
	filtered := FilterSince(entries, now.Add(-time.Hour))
	if len(filtered) != 1 || !filtered[0].Timestamp.Equal(entries[1].Timestamp) {
		t.Fatalf("filtered=%#v", filtered)
	}
}

func TestSensitiveFlagsAreRedacted(t *testing.T) {
	got := Redact([]string{"--token=abc", "--password", "value", "--region", "eu", "--secret-key", "hidden"})
	for _, value := range got {
		if value == "abc" || value == "value" || value == "hidden" || value == "--token=abc" {
			t.Fatalf("unredacted args = %#v", got)
		}
	}
}

func TestDisabledBehaviorIsCallerControlled(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.log")
	entries, err := Read(path, 0)
	if err != nil || len(entries) != 0 {
		t.Fatalf("entries=%#v err=%v", entries, err)
	}
}
