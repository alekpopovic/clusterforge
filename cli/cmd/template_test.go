package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateTemplatePackRequiresMetadata(t *testing.T) {
	if err := validateTemplatePack("missing", t.TempDir()); err == nil || !strings.Contains(err.Error(), "metadata") {
		t.Fatalf("expected missing metadata error, got %v", err)
	}
}

func TestValidateLocalTemplatePack(t *testing.T) {
	directory := t.TempDir()
	metadata := "name: company\nversion: v1\nsupported_clouds: [aws]\nsupported_orchestrators: [eks]\n"
	if err := os.WriteFile(filepath.Join(directory, "metadata.yaml"), []byte(metadata), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, subdirectory := range []string{"env", "app"} {
		if err := os.MkdirAll(filepath.Join(directory, subdirectory), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(directory, subdirectory, "main.tf.tmpl"), []byte("# template\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if err := validateTemplatePack("company", directory); err != nil {
		t.Fatal(err)
	}
}
