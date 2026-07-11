package cmd

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type fakeDoctorRunner struct {
	paths   map[string]string
	outputs map[string]string
	errs    map[string]error
}

func (f fakeDoctorRunner) LookPath(file string) (string, error) {
	path, ok := f.paths[file]
	if !ok {
		return "", errors.New("not found")
	}
	return path, nil
}

func (f fakeDoctorRunner) Output(ctx context.Context, name string, args ...string) ([]byte, error) {
	key := name + " " + strings.Join(args, " ")
	if err, ok := f.errs[key]; ok {
		return nil, err
	}
	output, ok := f.outputs[key]
	if !ok {
		return nil, errors.New("no fake output")
	}
	return []byte(output), nil
}

func TestParseTerraformVersion(t *testing.T) {
	version, ok := parseTerraformVersion("Terraform v1.8.5\non linux_amd64\n")
	if !ok {
		t.Fatal("expected version to parse")
	}
	if version != "1.8.5" {
		t.Fatalf("version = %q", version)
	}
}

func TestRunDoctorValidConfigPassesProjectChecks(t *testing.T) {
	dir := t.TempDir()
	envPath := filepath.Join(dir, "live", "dev", "aws-eks")
	if err := os.MkdirAll(envPath, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	configPath := writeDoctorConfig(t, dir, `
project:
  name: demo
environments:
  dev:
    path: `+envPath+`
`)

	report := runDoctor(context.Background(), configPath, healthyDoctorRunner())
	assertCheck(t, report, "project.config_load", doctorPass)
	assertCheck(t, report, "project.name", doctorPass)
	assertCheck(t, report, "project.environments", doctorPass)
	assertCheck(t, report, "environment.dev.path", doctorPass)
	if doctorHasFailure(report) {
		t.Fatalf("expected no failures: %#v", report.Checks)
	}
}

func TestRunDoctorInvalidConfigFails(t *testing.T) {
	dir := t.TempDir()
	configPath := writeDoctorConfig(t, dir, `
project: {}
environments: {}
`)

	report := runDoctor(context.Background(), configPath, healthyDoctorRunner())
	assertCheck(t, report, "project.config_load", doctorFail)
	if !doctorHasFailure(report) {
		t.Fatal("expected hard failure")
	}
}

func TestRunDoctorWarnsForProdLocalBackend(t *testing.T) {
	dir := t.TempDir()
	envPath := filepath.Join(dir, "live", "prod", "aws-eks")
	if err := os.MkdirAll(envPath, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	configPath := writeDoctorConfig(t, dir, `
project:
  name: demo
environments:
  prod:
    path: `+envPath+`
backends:
  prod:
    type: local
`)

	report := runDoctor(context.Background(), configPath, healthyDoctorRunner())
	assertCheck(t, report, "safety.prod.backend", doctorWarn)
}

func TestRunDoctorWarnsForDockerProd(t *testing.T) {
	dir := t.TempDir()
	envPath := filepath.Join(dir, "live", "prod", "docker")
	if err := os.MkdirAll(envPath, 0755); err != nil {
		t.Fatal(err)
	}
	configPath := writeDoctorConfig(t, dir, "\nproject:\n  name: demo\nenvironments:\n  prod:\n    cloud: local\n    orchestrator: docker\n    path: "+envPath+"\n")
	report := runDoctor(context.Background(), configPath, healthyDoctorRunner())
	assertCheck(t, report, "safety.prod.docker_target", doctorWarn)
}

func TestDockerTargetWarning(t *testing.T) {
	if dockerTargetWarning("docker") == "" || dockerTargetWarning("swarm") == "" {
		t.Fatal("expected warning")
	}
	if dockerTargetWarning("eks") != "" {
		t.Fatal("unexpected warning")
	}
}

func TestTrackedSensitiveFilesWarn(t *testing.T) {
	checks := checkTrackedSensitiveFiles([]string{
		"live/prod/terraform.tfstate",
		".env",
		"kubeconfig-prod",
	})

	statuses := map[string]doctorStatus{}
	for _, check := range checks {
		statuses[check.Name] = check.Status
	}
	if statuses["git.tracked_tfstate"] != doctorWarn {
		t.Fatalf("tfstate status = %s", statuses["git.tracked_tfstate"])
	}
	if statuses["git.tracked_env"] != doctorWarn {
		t.Fatalf("env status = %s", statuses["git.tracked_env"])
	}
	if statuses["git.tracked_kubeconfig"] != doctorWarn {
		t.Fatalf("kubeconfig status = %s", statuses["git.tracked_kubeconfig"])
	}
}

func healthyDoctorRunner() fakeDoctorRunner {
	return fakeDoctorRunner{
		paths: map[string]string{
			"terraform": "/usr/bin/terraform",
			"git":       "/usr/bin/git",
			"go":        "/usr/bin/go",
		},
		outputs: map[string]string{
			"terraform version":                   "Terraform v1.8.5\n",
			"go version":                          "go version go1.22.0 linux/amd64\n",
			"git rev-parse --is-inside-work-tree": "true\n",
			"git ls-files":                        "README.md\nmodules/example/main.tf\n",
		},
	}
}

func writeDoctorConfig(t *testing.T, dir, content string) string {
	t.Helper()
	path := filepath.Join(dir, "clusterforge.yaml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}
	return path
}

func assertCheck(t *testing.T, report doctorReport, name string, status doctorStatus) {
	t.Helper()
	for _, check := range report.Checks {
		if check.Name == name {
			if check.Status != status {
				t.Fatalf("%s status = %s, want %s: %s", name, check.Status, status, check.Message)
			}
			return
		}
	}
	t.Fatalf("check %s not found in %#v", name, report.Checks)
}
