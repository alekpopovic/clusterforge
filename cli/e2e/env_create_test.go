package e2e

import (
	"path/filepath"
	"testing"
)

func TestEnvCreateAndGenerateAWSEKS(t *testing.T) {
	dir := t.TempDir()
	runCF(t, dir, "project", "init", "demo")
	runCF(t, dir, "env", "create", "dev", "--cloud", "aws", "--orchestrator", "eks", "--region", "eu-central-1")
	runCF(t, dir, "generate", "dev")

	root := filepath.Join(dir, "live", "dev", "aws-eks")
	for _, file := range []string{
		"versions.tf",
		"backend.tf",
		"providers.tf",
		"main.tf",
		"variables.tf",
		"outputs.tf",
		"terraform.tfvars.example",
		"README.md",
	} {
		assertExists(t, filepath.Join(root, file))
	}
	assertContains(t, readFile(t, filepath.Join(root, "main.tf")), `module "eks"`)
}
