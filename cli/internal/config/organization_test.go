package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOrganizationConfigurationIsOptional(t *testing.T) {
	cfg := DefaultConfig("demo")
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestLoadOrganizationConfiguration(t *testing.T) {
	cfg := DefaultConfig("demo")
	cfg.Environments["dev"] = Environment{Cloud: "aws", Region: "eu-central-1", Orchestrator: "eks", Path: "live/dev"}
	cfg.Organization = Organization{Name: "example-org", Owner: "platform", Contact: "platform@example.com"}
	cfg.Workspaces["platform"] = Workspace{Environments: []string{"dev"}}
	cfg.Teams["platform"] = Team{Owners: []string{"platform@example.com"}, Namespaces: []string{"platform-system"}}
	path := filepath.Join(t.TempDir(), "clusterforge.yaml")
	if err := cfg.Save(path); err != nil {
		t.Fatal(err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Organization.Name != "example-org" || len(loaded.Workspaces) != 1 || len(loaded.Teams) != 1 {
		t.Fatalf("unexpected config: %#v", loaded)
	}
}

func TestDuplicateWorkspaceYAMLKeysFail(t *testing.T) {
	path := filepath.Join(t.TempDir(), "clusterforge.yaml")
	data := []byte("project: {name: demo, default_engine: terraform}\nengines: {terraform: {binary: terraform}}\ndefaults: {cloud: aws, region: x, orchestrator: eks}\nenvironments: {}\npolicies: {}\nworkspaces:\n  platform: {}\n  platform: {}\n")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(path); err == nil {
		t.Fatal("expected duplicate workspace key error")
	}
}
