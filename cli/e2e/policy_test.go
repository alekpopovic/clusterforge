package e2e

import (
	"testing"
)

func TestProdApplyAndDestroyPoliciesFailBeforeTerraform(t *testing.T) {
	dir := t.TempDir()
	runCF(t, dir, "project", "init", "demo")
	runCF(t, dir, "env", "create", "prod", "--cloud", "aws", "--orchestrator", "eks", "--region", "eu-central-1")

	applyOutput, applyErr := runCFAllowError(dir, "apply", "prod")
	if applyErr == nil {
		t.Fatalf("expected prod apply without plan file to fail:\n%s", applyOutput)
	}
	assertContains(t, applyOutput, "apply against prod requires --plan-file")

	destroyOutput, destroyErr := runCFAllowError(dir, "destroy", "prod")
	if destroyErr == nil {
		t.Fatalf("expected prod destroy to be blocked:\n%s", destroyOutput)
	}
	assertContains(t, destroyOutput, "destroy against prod is blocked by default")
}

func TestDoctorReportsProjectHealth(t *testing.T) {
	dir := t.TempDir()
	runCF(t, dir, "project", "init", "demo")
	runCF(t, dir, "env", "create", "dev", "--cloud", "aws", "--orchestrator", "eks", "--region", "eu-central-1")

	output, _ := runCFAllowError(dir, "doctor", "--json")
	assertContains(t, output, `"name": "project.config_file"`)
	assertContains(t, output, `"name": "project.environments"`)
	assertContains(t, output, `"name": "binary.git"`)
}
