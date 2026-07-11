package backup

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestEvidenceParsingAndJSONReport(t *testing.T) {
	root := t.TempDir()
	evidencePath := filepath.Join(root, EvidencePath)
	data := "backup_tests:\n  prod:\n    last_backup_test: \"2026-07-01\"\n    last_restore_test: \"2026-07-02\"\n    result: passed\n    notes: Restored namespace demo-restore\n"
	if err := os.WriteFile(evidencePath, []byte(data), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(`source = "modules/platform/kubernetes/velero"\nsource = "modules/cloud/aws/velero-backup-bucket"`), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := BuildReport("prod", root, "aws", evidencePath)
	if err != nil {
		t.Fatal(err)
	}
	if report.Evidence == nil || report.Evidence.LastRestoreTest != "2026-07-02" {
		t.Fatalf("report=%#v", report)
	}
	encoded, err := report.JSON()
	if err != nil || !json.Valid(encoded) {
		t.Fatalf("invalid report JSON: %s %v", encoded, err)
	}
}

func TestMissingRestoreTestWarns(t *testing.T) {
	root := t.TempDir()
	evidencePath := filepath.Join(root, EvidencePath)
	if err := os.WriteFile(evidencePath, []byte("backup_tests:\n  prod:\n    last_backup_test: \"2026-07-01\"\n    result: passed\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := BuildReport("prod", root, "gcp", evidencePath)
	if err != nil {
		t.Fatal(err)
	}
	if report.Overall != "warning" || report.Checks[len(report.Checks)-1].Status != "warn" {
		t.Fatalf("report=%#v", report)
	}
}
