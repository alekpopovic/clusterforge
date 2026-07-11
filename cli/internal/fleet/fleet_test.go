package fleet

import (
	"context"
	"errors"
	"testing"

	"github.com/textracta/clusterforge/cli/internal/inventory"
)

func TestApplyFilters(t *testing.T) {
	clusters := []inventory.Cluster{
		{Name: "dev-eks", Environment: "dev", Cloud: "aws", Orchestrator: "eks", Status: "experimental"},
		{Name: "prod-eks", Environment: "prod", Cloud: "aws", Orchestrator: "eks", Status: "production"},
		{Name: "prod-gke", Environment: "prod", Cloud: "gcp", Orchestrator: "gke", Status: "production"},
	}
	got := Apply(clusters, Filter{Environment: "prod", Cloud: "aws", Status: "production"})
	if len(got) != 1 || got[0].Name != "prod-eks" {
		t.Fatalf("filtered = %#v", got)
	}
}

func TestApplyFiltersByRegion(t *testing.T) {
	clusters := []inventory.Cluster{{Name: "eu", Region: "eu-central-1"}, {Name: "west", Region: "eu-west-1"}}
	filtered := Apply(clusters, Filter{Region: "eu-central-1"})
	if len(filtered) != 1 || filtered[0].Name != "eu" {
		t.Fatalf("unexpected filter: %#v", filtered)
	}
}

func TestRunContinuesAfterFailure(t *testing.T) {
	clusters := []inventory.Cluster{{Name: "bad"}, {Name: "good"}}
	results, err := Run(context.Background(), clusters, "drift", false, func(_ context.Context, cluster inventory.Cluster) (Result, error) {
		if cluster.Name == "bad" {
			return Result{}, errors.New("plan failed")
		}
		return Result{Status: "pass"}, nil
	})
	if err != nil || len(results) != 2 || results[0].Status != "fail" || results[1].Status != "pass" {
		t.Fatalf("results=%#v err=%v", results, err)
	}
}

func TestRunFailFast(t *testing.T) {
	results, err := Run(context.Background(), []inventory.Cluster{{Name: "bad"}, {Name: "unreached"}}, "drift", true, func(context.Context, inventory.Cluster) (Result, error) {
		return Result{}, errors.New("plan failed")
	})
	if err == nil || len(results) != 1 {
		t.Fatalf("results=%#v err=%v", results, err)
	}
}
