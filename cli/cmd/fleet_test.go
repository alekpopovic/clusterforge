package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alekpopovic/clusterforge/cli/internal/fleet"
)

func TestFleetListWithFiltersJSON(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "clusterforge.yaml")
	data := `project:
  name: demo
environments:
  dev:
    path: live/dev
  prod:
    path: live/prod
clusters:
  dev-eks:
    environment: dev
    cloud: aws
    orchestrator: eks
    region: eu-central-1
    path: live/dev
    status: experimental
  prod-eks:
    environment: prod
    cloud: aws
    orchestrator: eks
    region: eu-central-1
    path: live/prod
    status: production
`
	if err := os.WriteFile(configPath, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	oldPath, oldFilter, oldJSON := opts.ConfigPath, fleetFilter, fleetJSON
	opts.ConfigPath, fleetFilter, fleetJSON = configPath, fleet.Filter{Environment: "prod", Status: "production"}, true
	t.Cleanup(func() { opts.ConfigPath, fleetFilter, fleetJSON = oldPath, oldFilter, oldJSON })
	var out bytes.Buffer
	fleetListCmd.SetOut(&out)
	t.Cleanup(func() { fleetListCmd.SetOut(nil) })
	if err := fleetListCmd.RunE(fleetListCmd, nil); err != nil {
		t.Fatalf("fleet list: %v", err)
	}
	var response struct {
		Clusters []struct {
			Name string `json:"name"`
		} `json:"clusters"`
	}
	if err := json.Unmarshal(out.Bytes(), &response); err != nil {
		t.Fatalf("json: %v", err)
	}
	if len(response.Clusters) != 1 || response.Clusters[0].Name != "prod-eks" {
		t.Fatalf("response = %#v", response)
	}
}

func TestFleetDoctorAggregatesFailuresWithoutFailingCommand(t *testing.T) {
	dir := t.TempDir()
	existing := filepath.Join(dir, "existing")
	if err := os.MkdirAll(existing, 0o755); err != nil {
		t.Fatal(err)
	}
	configPath := filepath.Join(dir, "clusterforge.yaml")
	data := "project:\n  name: demo\nenvironments:\n  dev:\n    path: " + existing + "\n  prod:\n    path: " + filepath.Join(dir, "missing") + "\nclusters:\n  dev-eks:\n    environment: dev\n    path: " + existing + "\n    status: development\n  prod-eks:\n    environment: prod\n    path: " + filepath.Join(dir, "missing") + "\n    status: production\n"
	if err := os.WriteFile(configPath, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	oldPath, oldFilter, oldJSON := opts.ConfigPath, fleetFilter, fleetJSON
	opts.ConfigPath, fleetFilter, fleetJSON = configPath, fleet.Filter{}, false
	t.Cleanup(func() { opts.ConfigPath, fleetFilter, fleetJSON = oldPath, oldFilter, oldJSON })
	var out bytes.Buffer
	fleetDoctorCmd.SetOut(&out)
	t.Cleanup(func() { fleetDoctorCmd.SetOut(nil) })
	if err := fleetDoctorCmd.RunE(fleetDoctorCmd, nil); err != nil {
		t.Fatalf("fleet doctor should aggregate: %v", err)
	}
	if !strings.Contains(out.String(), "dev-eks") || !strings.Contains(out.String(), "prod-eks") || !strings.Contains(out.String(), "fail") {
		t.Fatalf("output = %q", out.String())
	}
}

func TestFleetPolicyCheckWarnsForProductionLikeEnvironment(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "clusterforge.yaml")
	data := `project:
  name: demo
policies:
  require_plan_file_for_apply: false
  block_destroy_in_prod: false
environments:
  prd:
    path: live/prd
clusters:
  prd-eks:
    environment: prd
    path: live/prd
    status: development
`
	if err := os.WriteFile(configPath, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	oldPath, oldFilter, oldJSON := opts.ConfigPath, fleetFilter, fleetJSON
	opts.ConfigPath, fleetFilter, fleetJSON = configPath, fleet.Filter{}, false
	t.Cleanup(func() { opts.ConfigPath, fleetFilter, fleetJSON = oldPath, oldFilter, oldJSON })
	var out bytes.Buffer
	fleetPolicyCheckCmd.SetOut(&out)
	t.Cleanup(func() { fleetPolicyCheckCmd.SetOut(nil) })
	if err := fleetPolicyCheckCmd.RunE(fleetPolicyCheckCmd, nil); err != nil {
		t.Fatalf("fleet policy check: %v", err)
	}
	if !strings.Contains(out.String(), "prd-eks\tprd\twarn\tproduction safety gates are incomplete") {
		t.Fatalf("output = %q", out.String())
	}
}

func TestSafeFleetNamePreventsPathTraversal(t *testing.T) {
	if got := safeFleetName("../../prod cluster"); strings.Contains(got, "/") || strings.Contains(got, " ") {
		t.Fatalf("safe name = %q", got)
	}
}
