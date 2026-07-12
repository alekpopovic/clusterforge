package gitops

import (
	"strings"
	"testing"

	cfapp "github.com/alekpopovic/clusterforge/cli/internal/app"
	"github.com/alekpopovic/clusterforge/cli/internal/config"
)

func TestRenderTwoClustersAndSpecificCluster(t *testing.T) {
	cfg := config.GitOps{Provider: "argocd", RepoURL: "https://github.com/example/gitops", Clusters: []config.GitOpsCluster{{Name: "dev-eks", Environment: "dev"}, {Name: "prod-eks", Environment: "prod"}}}
	apps := map[string]cfapp.Manifest{"api": {Name: "api", SecretEnv: map[string]cfapp.SecretRef{"TOKEN": {SecretName: "private", SecretKey: "token"}}}}
	all, err := Render(cfg, apps, "")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(all), "clusterforge-dev-eks") || !strings.Contains(string(all), "clusterforge-prod-eks") {
		t.Fatalf("missing clusters:\n%s", all)
	}
	one, err := Render(cfg, apps, "dev-eks")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(one), "api-prod-eks") || !strings.Contains(string(one), "api-dev-eks") {
		t.Fatalf("bad filtered output:\n%s", one)
	}
	if strings.Contains(string(all), "private") || strings.Contains(string(all), "token") {
		t.Fatalf("secret reference leaked:\n%s", all)
	}
}

func TestMissingRepoFails(t *testing.T) {
	_, err := Render(config.GitOps{Provider: "argocd"}, nil, "")
	if err == nil {
		t.Fatal("expected missing repo_url error")
	}
}
