package e2e

import (
	"path/filepath"
	"testing"
)

func TestAppAddValidateAndRender(t *testing.T) {
	dir := t.TempDir()
	runCF(t, dir, "project", "init", "demo")
	runCF(t, dir, "env", "create", "dev", "--cloud", "aws", "--orchestrator", "eks", "--region", "eu-central-1")
	runCF(t, dir, "generate", "dev")

	runCF(t, dir, "app", "add", "api", "--image", "nginx:1.25", "--port", "80")
	validateOutput := runCF(t, dir, "app", "validate", "api")
	assertContains(t, validateOutput, "apps/api.yaml: ok")
	runCF(t, dir, "app", "render", "api", "--env", "dev")

	rendered := readFile(t, filepath.Join(dir, "live", "dev", "aws-eks", "apps", "api.tf"))
	assertContains(t, rendered, `module "api"`)
	assertContains(t, rendered, `source = "../../../../modules/workloads/kubernetes/app"`)
	assertContains(t, rendered, `image     = "nginx:1.25"`)
}
