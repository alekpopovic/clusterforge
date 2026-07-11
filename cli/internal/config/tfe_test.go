package config

import "testing"

func TestTerraformCloudConfig(t *testing.T) {
	cfg := DefaultConfig("demo")
	cfg.Environments["dev"] = Environment{Cloud: "aws", Region: "x", Orchestrator: "eks", Path: "live/dev"}
	cfg.TerraformCloud = TerraformCloud{Enabled: true, Organization: "example-org", Project: "clusterforge", Workspaces: map[string]TerraformCloudWorkspace{"dev": {Name: "clusterforge-dev"}}}
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
}
func TestTerraformCloudDisabledNoOrganization(t *testing.T) {
	cfg := DefaultConfig("demo")
	cfg.TerraformCloud = TerraformCloud{Enabled: false}
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
}
