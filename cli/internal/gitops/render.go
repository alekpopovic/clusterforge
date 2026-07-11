package gitops

import (
	"fmt"
	"sort"
	"strings"

	cfapp "github.com/textracta/clusterforge/cli/internal/app"
	"github.com/textracta/clusterforge/cli/internal/config"
	"gopkg.in/yaml.v3"
)

type manifest struct {
	APIVersion string         `yaml:"apiVersion"`
	Kind       string         `yaml:"kind"`
	Metadata   map[string]any `yaml:"metadata"`
	Spec       map[string]any `yaml:"spec"`
}

func Validate(cfg config.GitOps) error {
	if cfg.Provider == "" {
		return fmt.Errorf("gitops.provider is required")
	}
	if cfg.Provider != "argocd" {
		return fmt.Errorf("gitops provider %q is not implemented; use argocd", cfg.Provider)
	}
	if strings.TrimSpace(cfg.RepoURL) == "" {
		return fmt.Errorf("gitops.repo_url is required")
	}
	if strings.Contains(strings.ToLower(cfg.RepoURL), "token@") || strings.Contains(strings.ToLower(cfg.RepoURL), "password") {
		return fmt.Errorf("gitops.repo_url must not contain credentials")
	}
	seen := map[string]bool{}
	for _, cluster := range cfg.Clusters {
		if cluster.Name == "" || cluster.Environment == "" {
			return fmt.Errorf("gitops clusters require name and environment")
		}
		if seen[cluster.Name] {
			return fmt.Errorf("duplicate gitops cluster %q", cluster.Name)
		}
		seen[cluster.Name] = true
	}
	return nil
}

func Render(cfg config.GitOps, apps map[string]cfapp.Manifest, clusterFilter string) ([]byte, error) {
	if err := Validate(cfg); err != nil {
		return nil, err
	}
	clusters := append([]config.GitOpsCluster{}, cfg.Clusters...)
	sort.Slice(clusters, func(i, j int) bool { return clusters[i].Name < clusters[j].Name })
	if clusterFilter != "" {
		found := false
		for _, cluster := range clusters {
			if cluster.Name == clusterFilter {
				found = true
			}
		}
		if !found {
			return nil, fmt.Errorf("gitops cluster %q not found", clusterFilter)
		}
	}
	var documents []manifest
	destinations := make([]map[string]string, 0, len(clusters))
	for _, cluster := range clusters {
		destinations = append(destinations, map[string]string{"name": cluster.Name, "namespace": "*"})
	}
	documents = append(documents, manifest{APIVersion: "argoproj.io/v1alpha1", Kind: "AppProject", Metadata: map[string]any{"name": "clusterforge", "namespace": "argocd"}, Spec: map[string]any{"sourceRepos": []string{cfg.RepoURL}, "destinations": destinations}})
	appNames := make([]string, 0, len(apps))
	for name := range apps {
		appNames = append(appNames, name)
	}
	sort.Strings(appNames)
	for _, cluster := range clusters {
		if clusterFilter != "" && cluster.Name != clusterFilter {
			continue
		}
		documents = append(documents, application("clusterforge-"+cluster.Name, cfg.RepoURL, "clusters/"+cluster.Name+"/apps", cluster.Name, "argocd"))
		for _, appName := range appNames {
			documents = append(documents, application(appName+"-"+cluster.Name, cfg.RepoURL, "apps/"+appName+"/overlays/"+cluster.Environment, cluster.Name, appName))
		}
	}
	var output strings.Builder
	for index, document := range documents {
		data, err := yaml.Marshal(document)
		if err != nil {
			return nil, err
		}
		if index > 0 {
			output.WriteString("---\n")
		}
		output.Write(data)
	}
	return []byte(output.String()), nil
}

func application(name, repoURL, path, cluster, namespace string) manifest {
	return manifest{APIVersion: "argoproj.io/v1alpha1", Kind: "Application", Metadata: map[string]any{"name": name, "namespace": "argocd"}, Spec: map[string]any{"project": "clusterforge", "source": map[string]string{"repoURL": repoURL, "targetRevision": "HEAD", "path": path}, "destination": map[string]string{"name": cluster, "namespace": namespace}, "syncPolicy": map[string]any{"syncOptions": []string{"CreateNamespace=true"}}}}
}
