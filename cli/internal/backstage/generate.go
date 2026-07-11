package backstage

import (
	"fmt"
	cfapp "github.com/textracta/clusterforge/cli/internal/app"
	"github.com/textracta/clusterforge/cli/internal/config"
	"gopkg.in/yaml.v3"
	"sort"
)

type Entity struct {
	APIVersion string         `yaml:"apiVersion"`
	Kind       string         `yaml:"kind"`
	Metadata   Metadata       `yaml:"metadata"`
	Spec       map[string]any `yaml:"spec"`
}
type Metadata struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
}

func Generate(cfg *config.Config, apps map[string]cfapp.Manifest, appFilter, envFilter string) ([]byte, error) {
	if !cfg.Backstage.Enabled {
		return nil, fmt.Errorf("Backstage integration is disabled")
	}
	owner := fallback(cfg.Backstage.Owner, cfg.Organization.Owner, "platform-team")
	system := fallback(cfg.Backstage.System, cfg.Project.Name)
	lifecycle := fallback(cfg.Backstage.Lifecycle, "experimental")
	entities := []Entity{{"backstage.io/v1alpha1", "System", Metadata{system, "ClusterForge project " + cfg.Project.Name}, map[string]any{"owner": owner}}}
	names := keys(cfg.Teams)
	for _, name := range names {
		entities = append(entities, Entity{"backstage.io/v1alpha1", "Group", Metadata{name, ""}, map[string]any{"type": "team", "children": []string{}}})
	}
	envNames := keys(cfg.Environments)
	for _, name := range envNames {
		if envFilter != "" && envFilter != name {
			continue
		}
		env := cfg.Environments[name]
		entities = append(entities, Entity{"backstage.io/v1alpha1", "Resource", Metadata{name, env.Cloud + " " + env.Orchestrator + " environment"}, map[string]any{"type": "infrastructure", "owner": owner, "system": system, "lifecycle": lifecycle}})
	}
	appNames := make([]string, 0, len(apps))
	for name := range apps {
		appNames = append(appNames, name)
	}
	sort.Strings(appNames)
	for _, name := range appNames {
		if appFilter != "" && appFilter != name {
			continue
		}
		m := apps[name]
		entities = append(entities, Entity{"backstage.io/v1alpha1", "Component", Metadata{name, ""}, map[string]any{"type": "service", "owner": fallback(m.Backstage.Owner, owner), "system": fallback(m.Backstage.System, system), "lifecycle": fallback(m.Backstage.Lifecycle, lifecycle)}})
	}
	var out []byte
	for i, e := range entities {
		data, err := yaml.Marshal(e)
		if err != nil {
			return nil, err
		}
		if i > 0 {
			out = append(out, []byte("---\n")...)
		}
		out = append(out, data...)
	}
	return out, nil
}
func fallback(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
func keys[T any](values map[string]T) []string {
	result := make([]string, 0, len(values))
	for key := range values {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}
