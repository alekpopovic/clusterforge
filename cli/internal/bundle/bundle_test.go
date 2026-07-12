package bundle

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alekpopovic/clusterforge/cli/internal/config"
)

func TestCreateExcludesSensitiveArtifactsAndVerify(t *testing.T) {
	root := t.TempDir()
	for _, dir := range []string{"modules/demo", "live/prod", "apps", "policies/demo", "docs/dr"} {
		if err := os.MkdirAll(filepath.Join(root, dir), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	files := map[string]string{"modules/demo/main.tf": "terraform {}", "live/prod/main.tf": "module \"x\" { source = \"../../modules/demo\" }", "live/prod/terraform.tfstate": "super-secret-state", "live/prod/kubeconfig": "super-secret-kubeconfig", "apps/api.yaml": "name: api\ntype: web\nimage: registry.example/api@sha256:abc\nenv:\n  PASSWORD: super-secret-value\nsecret_env:\n  TOKEN:\n    secret_name: api-secret\n    secret_key: token\n", "policies/demo/README.md": "policy", "docs/dr/restore.md": "# Restore"}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(root, name), []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	cfg := config.DefaultConfig("demo")
	cfg.Environments["prod"] = config.Environment{Path: "live/prod"}
	output := filepath.Join(root, "out")
	if err := Create(root, output, "prod", cfg); err != nil {
		t.Fatal(err)
	}
	if err := Verify(output); err != nil {
		t.Fatal(err)
	}
	var joined strings.Builder
	filepath.WalkDir(output, func(path string, entry os.DirEntry, err error) error {
		if err == nil && !entry.IsDir() {
			data, _ := os.ReadFile(path)
			joined.Write(data)
		}
		return err
	})
	if strings.Contains(joined.String(), "super-secret") || strings.Contains(joined.String(), "tfstate") || strings.Contains(joined.String(), "kubeconfig") {
		t.Fatalf("sensitive artifact leaked: %s", joined.String())
	}
	if err := os.WriteFile(filepath.Join(output, "images.txt"), []byte("tampered\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := Verify(output); err == nil {
		t.Fatal("expected checksum mismatch")
	}
}
