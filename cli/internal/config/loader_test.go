package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "clusterforge.yaml")
	data := []byte(`project:
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
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

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
