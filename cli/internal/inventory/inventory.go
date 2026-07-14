package inventory

import (
	"fmt"
	"sort"
	"strings"

	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"github.com/alekpopovic/clusterforge/cli/internal/environment"
)

type Cluster struct {
	Name           string `json:"name"`
	Environment    string `json:"environment"`
	Cloud          string `json:"cloud"`
	Orchestrator   string `json:"orchestrator"`
	Region         string `json:"region"`
	Path           string `json:"path"`
	Status         string `json:"status"`
	KubeconfigPath string `json:"kubeconfig_path,omitempty"`
	Legacy         bool   `json:"legacy_environment,omitempty"`
}

func List(cfg *config.Config) []Cluster {
	clusters := make([]Cluster, 0)
	if len(cfg.Clusters) > 0 {
		for name, cluster := range cfg.Clusters {
			clusters = append(clusters, Cluster{Name: name, Environment: cluster.Environment, Cloud: cluster.Cloud, Orchestrator: cluster.Orchestrator, Region: cluster.Region, Path: cluster.Path, Status: cluster.Status, KubeconfigPath: cluster.KubeconfigPath})
		}
	} else {
		for name, env := range cfg.Environments {
			status := "development"
			if environment.IsProduction(name) {
				status = "production"
			}
			clusters = append(clusters, Cluster{Name: name, Environment: name, Cloud: env.Cloud, Orchestrator: env.Orchestrator, Region: env.Region, Path: env.Path, Status: status, Legacy: true})
		}
	}
	sort.Slice(clusters, func(i, j int) bool { return clusters[i].Name < clusters[j].Name })
	return clusters
}

func Find(cfg *config.Config, name string) (Cluster, error) {
	for _, cluster := range List(cfg) {
		if cluster.Name == name {
			return cluster, nil
		}
	}
	return Cluster{}, fmt.Errorf("cluster %q not found", name)
}

func IsKubernetes(orchestrator string) bool {
	switch strings.ToLower(orchestrator) {
	case "eks", "aks", "gke", "kubernetes", "k3s", "rke2":
		return true
	default:
		return false
	}
}
