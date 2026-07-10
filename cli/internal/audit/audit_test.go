package audit

import (
	"path/filepath"
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
