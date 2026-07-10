package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/textracta/clusterforge/cli/internal/audit"
)

func TestAuditedCommandWritesEntry(t *testing.T) {
	dir := t.TempDir()
	logPath := filepath.Join(dir, ".cf", "audit.log")
	configPath := filepath.Join(dir, "clusterforge.yaml")
	data := "project:\n  name: demo\nenvironments:\n  dev:\n    path: live/dev\naudit:\n  enabled: true\n  path: " + logPath + "\n"
	if err := os.WriteFile(configPath, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	oldPath, oldPlan, oldJSON := opts.ConfigPath, policyPlanFile, policyCheckJSON
	opts.ConfigPath, policyPlanFile, policyCheckJSON = configPath, "", false
	t.Cleanup(func() { opts.ConfigPath, policyPlanFile, policyCheckJSON = oldPath, oldPlan, oldJSON })
	if err := policyCheckCmd.RunE(policyCheckCmd, []string{"dev"}); err != nil {
		t.Fatalf("policy check: %v", err)
	}
	entries, err := audit.Read(logPath, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].Command != "policy check" || entries[0].Environment != "dev" || entries[0].Result != "success" {
		t.Fatalf("entries = %#v", entries)
	}
}

func TestAuditDisabledWritesNothing(t *testing.T) {
	dir := t.TempDir()
	logPath := filepath.Join(dir, "audit.log")
	configPath := filepath.Join(dir, "clusterforge.yaml")
	data := "project:\n  name: demo\nenvironments:\n  dev:\n    path: live/dev\naudit:\n  enabled: false\n  path: " + logPath + "\n"
	if err := os.WriteFile(configPath, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	oldPath, oldPlan := opts.ConfigPath, policyPlanFile
	opts.ConfigPath, policyPlanFile = configPath, ""
	t.Cleanup(func() { opts.ConfigPath, policyPlanFile = oldPath, oldPlan })
	if err := policyCheckCmd.RunE(policyCheckCmd, []string{"dev"}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(logPath); !os.IsNotExist(err) {
		t.Fatalf("audit log should not exist: %v", err)
	}
}

func TestAuditClearRequiresConfirmation(t *testing.T) {
	oldNonInteractive, oldYes := opts.NonInteractive, auditClearYes
	opts.NonInteractive, auditClearYes = true, false
	t.Cleanup(func() { opts.NonInteractive, auditClearYes = oldNonInteractive, oldYes })
	err := auditClearCmd.RunE(auditClearCmd, nil)
	if err == nil || !strings.Contains(err.Error(), "requires --yes") {
		t.Fatalf("unexpected error: %v", err)
	}
}
