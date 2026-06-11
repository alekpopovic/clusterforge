package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadValidConfig(t *testing.T) {
	path := writeConfig(t, `project:
  name: demo
  default_engine: terraform
engines:
  terraform:
    binary: terraform
  opentofu:
    binary: tofu
defaults:
  cloud: aws
  region: eu-central-1
  orchestrator: eks
environments:
  dev:
    cloud: aws
    region: eu-central-1
    orchestrator: eks
    path: live/dev/aws-eks
policies:
  require_plan_file_for_apply: true
  block_destroy_in_prod: true
  require_manual_approval_for_prod: true
`)

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}
	if cfg.Project.Name != "demo" {
		t.Fatalf("project name = %q", cfg.Project.Name)
	}
	if cfg.Environments["dev"].Path != "live/dev/aws-eks" {
		t.Fatalf("dev path = %q", cfg.Environments["dev"].Path)
	}
}

func TestLoadFailsOnMissingProjectName(t *testing.T) {
	path := writeConfig(t, `project: {}
environments: {}
`)

	if _, err := Load(path); err == nil {
		t.Fatal("expected missing project.name to fail")
	}
}

func TestLoadAppliesDefaults(t *testing.T) {
	path := writeConfig(t, `project:
  name: demo
environments:
  dev:
    path: live/dev/aws-eks
`)

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}
	if cfg.Project.DefaultEngine != "terraform" {
		t.Fatalf("default engine = %q", cfg.Project.DefaultEngine)
	}
	if cfg.Engines["terraform"].Binary != "terraform" {
		t.Fatalf("terraform binary = %q", cfg.Engines["terraform"].Binary)
	}
	if cfg.Engines["opentofu"].Binary != "tofu" {
		t.Fatalf("opentofu binary = %q", cfg.Engines["opentofu"].Binary)
	}
	if cfg.Defaults.Cloud != "aws" || cfg.Defaults.Region != "eu-central-1" || cfg.Defaults.Orchestrator != "eks" {
		t.Fatalf("defaults = %#v", cfg.Defaults)
	}
	if cfg.Environments["dev"].Cloud != "aws" || cfg.Environments["dev"].Orchestrator != "eks" {
		t.Fatalf("dev environment defaults = %#v", cfg.Environments["dev"])
	}
}

func TestAddEnvironmentAndSave(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "clusterforge.yaml")
	cfg := DefaultConfig("demo")
	cfg.Environments["staging"] = Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
		Path:         "live/staging/aws-eks",
	}

	if err := cfg.Save(path); err != nil {
		t.Fatalf("save config: %v", err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("reload config: %v", err)
	}
	if loaded.Environments["staging"].Path != "live/staging/aws-eks" {
		t.Fatalf("staging environment = %#v", loaded.Environments["staging"])
	}
}

func TestValidateUnknownOrchestratorFails(t *testing.T) {
	path := writeConfig(t, `project:
  name: demo
environments:
  dev:
    orchestrator: mesos
    path: live/dev/mesos
`)

	if _, err := Load(path); err == nil {
		t.Fatal("expected unknown orchestrator to fail")
	}
}

func TestDefaultEngineMustExistInEngines(t *testing.T) {
	path := writeConfig(t, `project:
  name: demo
  default_engine: custom
environments:
  dev:
    path: live/dev/aws-eks
`)

	if _, err := Load(path); err == nil {
		t.Fatal("expected unknown default engine to fail")
	}
}

func TestEngineBinaryAliasesToOpenTofu(t *testing.T) {
	cfg := DefaultConfig("demo")
	binary, err := cfg.EngineBinary("tofu")
	if err != nil {
		t.Fatalf("engine binary: %v", err)
	}
	if binary != "tofu" {
		t.Fatalf("binary = %q", binary)
	}
}

func writeConfig(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "clusterforge.yaml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}
	return path
}
