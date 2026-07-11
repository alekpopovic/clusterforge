package dashboard

import (
	"time"

	cfapp "github.com/textracta/clusterforge/cli/internal/app"
	"github.com/textracta/clusterforge/cli/internal/config"
	"github.com/textracta/clusterforge/cli/internal/inventory"
	"github.com/textracta/clusterforge/cli/internal/runbooks"
	"github.com/textracta/clusterforge/cli/internal/servicecatalog"
)

const SchemaVersion = "1.0"

type Export struct {
	SchemaVersion string                            `json:"schema_version"`
	GeneratedAt   time.Time                         `json:"generated_at"`
	Project       Project                           `json:"project"`
	Organization  config.Organization               `json:"organization,omitempty"`
	Workspaces    map[string]config.Workspace       `json:"workspaces,omitempty"`
	Environments  []Environment                     `json:"environments"`
	Clusters      []Cluster                         `json:"clusters"`
	Apps          []App                             `json:"apps"`
	Services      map[string]servicecatalog.Service `json:"services,omitempty"`
	Runbooks      []Runbook                         `json:"runbooks"`
	ModuleCatalog string                            `json:"module_catalog,omitempty"`
	Policy        Evidence                          `json:"policy"`
	Drift         Evidence                          `json:"drift"`
	Cost          Evidence                          `json:"cost"`
}

type Project struct {
	Name          string `json:"name"`
	DefaultEngine string `json:"default_engine"`
}

type Environment struct {
	Name         string            `json:"name"`
	Cloud        string            `json:"cloud"`
	Region       string            `json:"region"`
	Orchestrator string            `json:"orchestrator"`
	Path         string            `json:"path"`
	Layout       string            `json:"layout"`
	Stacks       map[string]string `json:"stacks,omitempty"`
}

type Cluster struct {
	Name         string `json:"name"`
	Environment  string `json:"environment"`
	Cloud        string `json:"cloud"`
	Region       string `json:"region"`
	Orchestrator string `json:"orchestrator"`
	Status       string `json:"status,omitempty"`
}

type App struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Image     string `json:"image,omitempty"`
	Service   string `json:"service,omitempty"`
	Owner     string `json:"owner,omitempty"`
	System    string `json:"system,omitempty"`
	Lifecycle string `json:"lifecycle,omitempty"`
}

type Runbook struct {
	Name     string   `json:"name"`
	Title    string   `json:"title"`
	Category string   `json:"category,omitempty"`
	Severity string   `json:"severity,omitempty"`
	Path     string   `json:"path"`
	Tags     []string `json:"tags,omitempty"`
}

type Evidence struct {
	Available bool   `json:"available"`
	Path      string `json:"path,omitempty"`
}

func Build(cfg *config.Config, apps map[string]cfapp.Manifest, catalog servicecatalog.Catalog, books []runbooks.Runbook, envFilter, moduleCatalog string, evidence map[string]bool) Export {
	out := Export{
		SchemaVersion: SchemaVersion,
		GeneratedAt:   time.Now().UTC(),
		Project:       Project{Name: cfg.Project.Name, DefaultEngine: cfg.Project.DefaultEngine},
		Organization:  cfg.Organization,
		Workspaces:    cfg.Workspaces,
		Services:      catalog.Services,
		ModuleCatalog: moduleCatalog,
		Policy:        Evidence{Available: evidence["policy"], Path: evidencePath(evidence["policy"], ".cf/dashboard/policy.json")},
		Drift:         Evidence{Available: evidence["drift"], Path: evidencePath(evidence["drift"], ".cf/dashboard/drift.json")},
		Cost:          Evidence{Available: evidence["cost"], Path: evidencePath(evidence["cost"], ".cf/dashboard/cost.json")},
	}
	for name, env := range cfg.Environments {
		if envFilter != "" && envFilter != name {
			continue
		}
		stacks := make(map[string]string, len(env.Stacks))
		for stackName, stack := range env.Stacks {
			stacks[stackName] = stack.Path
		}
		out.Environments = append(out.Environments, Environment{Name: name, Cloud: env.Cloud, Region: env.Region, Orchestrator: env.Orchestrator, Path: env.Path, Layout: env.EffectiveLayout(), Stacks: stacks})
	}
	for _, cluster := range inventory.List(cfg) {
		if envFilter != "" && cluster.Environment != envFilter {
			continue
		}
		out.Clusters = append(out.Clusters, Cluster{Name: cluster.Name, Environment: cluster.Environment, Cloud: cluster.Cloud, Region: cluster.Region, Orchestrator: cluster.Orchestrator, Status: cluster.Status})
	}
	for name, manifest := range apps {
		out.Apps = append(out.Apps, App{Name: name, Type: manifest.Type, Image: manifest.Image, Service: manifest.Service, Owner: manifest.Backstage.Owner, System: manifest.Backstage.System, Lifecycle: manifest.Backstage.Lifecycle})
	}
	for _, book := range books {
		out.Runbooks = append(out.Runbooks, Runbook{Name: book.Name, Title: book.Title, Category: book.Category, Severity: book.Severity, Path: book.Path, Tags: book.Tags})
	}
	sortExport(&out)
	return out
}

func evidencePath(available bool, path string) string {
	if available {
		return path
	}
	return ""
}
