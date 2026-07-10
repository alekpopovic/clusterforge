package inventory

import (
	"testing"

	"github.com/textracta/clusterforge/cli/internal/config"
)

func TestListUsesConfiguredClusters(t *testing.T) {
	cfg := &config.Config{Clusters: map[string]config.Cluster{
		"prod-eks": {Environment: "prod", Cloud: "aws", Orchestrator: "eks", Region: "eu-central-1", Path: "live/prod/aws-eks", Status: "production"},
		"dev-eks":  {Environment: "dev", Cloud: "aws", Orchestrator: "eks", Region: "eu-central-1", Path: "live/dev/aws-eks", Status: "experimental"},
	}}
	clusters := List(cfg)
	if len(clusters) != 2 || clusters[0].Name != "dev-eks" || clusters[1].Status != "production" {
		t.Fatalf("clusters = %#v", clusters)
	}
}

func TestListFallsBackToEnvironments(t *testing.T) {
	cfg := &config.Config{Environments: map[string]config.Environment{"dev": {Cloud: "aws", Orchestrator: "eks", Path: "live/dev"}}}
	clusters := List(cfg)
	if len(clusters) != 1 || !clusters[0].Legacy || clusters[0].Environment != "dev" {
		t.Fatalf("clusters = %#v", clusters)
	}
}

func TestFindMissingCluster(t *testing.T) {
	if _, err := Find(&config.Config{}, "missing"); err == nil {
		t.Fatal("expected missing cluster error")
	}
}
