package policyengine

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuiltInRepositoryPolicy(t *testing.T) {
	result := Evaluate(Input{TrackedFiles: []string{"state/prod.tfstate"}}, Options{Pack: "baseline"})
	if !result.Blocked || len(result.Findings) != 1 || result.Findings[0].PolicyID != "CF-REPO-001" {
		t.Fatalf("unexpected result: %#v", result)
	}
}

func TestLoadPolicyPackOverrides(t *testing.T) {
	directory := t.TempDir()
	data := []byte("policies:\n  - id: CF-IAM-001\n    action: block\n")
	if err := os.WriteFile(filepath.Join(directory, "policies.yaml"), data, 0o644); err != nil {
		t.Fatal(err)
	}
	overrides, err := LoadOverrides(directory)
	if err != nil {
		t.Fatal(err)
	}
	if overrides["CF-IAM-001"] != "block" {
		t.Fatalf("unexpected overrides: %#v", overrides)
	}
}

func TestProductionPackBlocksExpectedIssues(t *testing.T) {
	result := Evaluate(Input{Production: true, BackendType: "local", Image: "example/api:latest"}, Options{Pack: "production"})
	if !result.Blocked || len(result.Findings) < 4 {
		t.Fatalf("expected blocking findings: %#v", result)
	}
}

func TestAdvisoryOverrideDoesNotBlock(t *testing.T) {
	result := Evaluate(Input{Production: true, BackendType: "local", RequirePlanFile: true, BlockProdDestroy: true}, Options{Pack: "production", Overrides: map[string]string{"CF-PROD-003": "advisory"}})
	if result.Blocked {
		t.Fatalf("advisory policy blocked: %#v", result)
	}
}

func TestJSONAndSARIFOutput(t *testing.T) {
	result := Evaluate(Input{TrackedFiles: []string{".env"}}, Options{Pack: "baseline"})
	if _, err := json.Marshal(result); err != nil {
		t.Fatal(err)
	}
	data, err := SARIF(result)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"version": "2.1.0"`) || !json.Valid(data) {
		t.Fatalf("invalid SARIF: %s", data)
	}
}
