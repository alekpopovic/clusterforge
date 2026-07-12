package e2e

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateExistingKubernetesWithoutCloudCredentials(t *testing.T) {
	dir := t.TempDir()
	runCF(t, dir, "project", "init", "demo")
	runCF(t, dir, "env", "create", "dev", "--cloud", "existing", "--orchestrator", "kubernetes", "--region", "local")
	runCF(t, dir, "generate", "dev")

	root := filepath.Join(dir, "live", "dev", "existing-kubernetes")
	assertContains(t, readFile(t, filepath.Join(root, "providers.tf")), "config_path")
	mainTF := readFile(t, filepath.Join(root, "main.tf"))
	if strings.Contains(mainTF, `module "network"`) || strings.Contains(mainTF, `module "eks"`) {
		t.Fatalf("existing Kubernetes generation should not include cloud infrastructure:\n%s", mainTF)
	}
}

func TestDefaultWizardProjectCanGenerate(t *testing.T) {
	dir := t.TempDir()
	runCF(t, dir, "--non-interactive", "wizard", "--defaults")
	runCF(t, dir, "generate", "dev")

	root := filepath.Join(dir, "live", "dev", "existing-kubernetes")
	assertContains(t, readFile(t, filepath.Join(root, "providers.tf")), "config_path")
}
