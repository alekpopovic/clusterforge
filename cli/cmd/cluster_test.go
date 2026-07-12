package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alekpopovic/clusterforge/cli/internal/inventory"
)

func TestClusterListAndShow(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "clusterforge.yaml")
	data := `project:
  name: demo
environments:
  dev:
    path: live/dev/aws-eks
clusters:
  dev-eks:
    environment: dev
    cloud: aws
    orchestrator: eks
    region: eu-central-1
    path: live/dev/aws-eks
    status: experimental
`
	if err := os.WriteFile(configPath, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	oldPath, oldListJSON, oldShowJSON := opts.ConfigPath, clusterListJSON, clusterShowJSON
	opts.ConfigPath, clusterListJSON, clusterShowJSON = configPath, true, false
	t.Cleanup(func() { opts.ConfigPath, clusterListJSON, clusterShowJSON = oldPath, oldListJSON, oldShowJSON })

	var listOut bytes.Buffer
	clusterListCmd.SetOut(&listOut)
	t.Cleanup(func() { clusterListCmd.SetOut(nil) })
	if err := clusterListCmd.RunE(clusterListCmd, nil); err != nil {
		t.Fatalf("list: %v", err)
	}
	var response struct {
		Clusters []inventory.Cluster `json:"clusters"`
	}
	if err := json.Unmarshal(listOut.Bytes(), &response); err != nil {
		t.Fatalf("json: %v", err)
	}
	if len(response.Clusters) != 1 || response.Clusters[0].Name != "dev-eks" {
		t.Fatalf("response = %#v", response)
	}

	var showOut bytes.Buffer
	clusterShowCmd.SetOut(&showOut)
	t.Cleanup(func() { clusterShowCmd.SetOut(nil) })
	if err := clusterShowCmd.RunE(clusterShowCmd, []string{"dev-eks"}); err != nil {
		t.Fatalf("show: %v", err)
	}
	if !strings.Contains(showOut.String(), "status: experimental") {
		t.Fatalf("show output = %q", showOut.String())
	}
}

func TestClusterShowMissingFailsClearly(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "clusterforge.yaml")
	if err := os.WriteFile(path, []byte("project:\n  name: demo\nenvironments: {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	oldPath := opts.ConfigPath
	opts.ConfigPath = path
	t.Cleanup(func() { opts.ConfigPath = oldPath })
	if err := clusterShowCmd.RunE(clusterShowCmd, []string{"missing"}); err == nil || !strings.Contains(err.Error(), "cluster \"missing\" not found") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestInspectClusterDoesNotRequireCloudAccess(t *testing.T) {
	dir := t.TempDir()
	report := inspectCluster(inventory.Cluster{Name: "dev", Path: dir, Orchestrator: "eks"})
	if report.Status == doctorFail || len(report.Checks) != 2 || report.Checks[1].Status != doctorWarn {
		t.Fatalf("report = %#v", report)
	}
}
