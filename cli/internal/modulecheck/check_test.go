package modulecheck

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckModulePassesRequiredFiles(t *testing.T) {
	root := t.TempDir()
	module := filepath.Join(root, "modules", "example")
	writeModule(t, module)
	if err := os.WriteFile(filepath.Join(root, "MODULE_CATALOG.md"), []byte("modules/example\n"), 0o644); err != nil {
		t.Fatalf("write catalog: %v", err)
	}

	report, err := Check(Options{Root: root, Path: "modules"})
	if err != nil {
		t.Fatalf("check: %v", err)
	}
	if report.Status != Pass {
		t.Fatalf("status = %s, report = %#v", report.Status, report)
	}
}

func TestCheckModuleFailsMissingRequiredFile(t *testing.T) {
	root := t.TempDir()
	module := filepath.Join(root, "modules", "example")
	writeModule(t, module)
	if err := os.Remove(filepath.Join(module, "outputs.tf")); err != nil {
		t.Fatalf("remove outputs: %v", err)
	}

	report, err := Check(Options{Root: root, Path: "modules/example"})
	if err != nil {
		t.Fatalf("check: %v", err)
	}
	if report.Status != Fail {
		t.Fatalf("status = %s, want fail", report.Status)
	}
}

func TestCheckModuleWarnsForMissingDescription(t *testing.T) {
	root := t.TempDir()
	module := filepath.Join(root, "modules", "example")
	writeModule(t, module)
	if err := os.WriteFile(filepath.Join(module, "variables.tf"), []byte(`variable "name" { type = string }`), 0o644); err != nil {
		t.Fatalf("write variables: %v", err)
	}

	report, err := Check(Options{Root: root, Path: "modules/example"})
	if err != nil {
		t.Fatalf("check: %v", err)
	}
	if report.Status != Warn {
		t.Fatalf("status = %s, want warn", report.Status)
	}
}

func writeModule(t *testing.T, dir string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir module: %v", err)
	}
	files := map[string]string{
		"main.tf":      "",
		"variables.tf": `variable "name" { description = "Name." type = string }`,
		"outputs.tf":   `output "name" { description = "Name." value = var.name }`,
		"versions.tf":  `terraform { required_version = ">= 1.6.0" }`,
		"README.md":    "# Example\n\n## Usage\n\nStatus: stable\n",
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
}
