package planjson

import (
	"os"
	"strings"
	"testing"
)

func TestParseMixedPlan(t *testing.T) {
	data := readFixture(t, "mixed.json")
	summary, err := Parse(data)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if summary.Creates != 1 || summary.Updates != 1 || summary.Deletes != 1 || summary.Replacements != 1 || summary.NoOps != 1 {
		t.Fatalf("summary = %#v", summary)
	}
	if len(summary.Addresses) != 4 {
		t.Fatalf("addresses = %#v", summary.Addresses)
	}
}

func TestParseInvalidJSON(t *testing.T) {
	if _, err := Parse([]byte(`{`)); err == nil {
		t.Fatal("expected parse error")
	}
}

func TestPrintSummary(t *testing.T) {
	summary, err := Parse(readFixture(t, "creates_updates.json"))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	var out strings.Builder
	Print(&out, "dev", summary, "LOW", "apply allowed")
	if !strings.Contains(out.String(), "Plan summary for dev:") {
		t.Fatalf("output = %q", out.String())
	}
	if !strings.Contains(out.String(), "Risk: LOW") {
		t.Fatalf("output = %q", out.String())
	}
}

func readFixture(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	return data
}
