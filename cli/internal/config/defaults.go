package config

var allowedClouds = map[string]bool{
	"aws":      true,
	"azure":    true,
	"gcp":      true,
	"hetzner":  true,
	"local":    true,
	"existing": true,
}

var allowedOrchestrators = map[string]bool{
	"eks":        true,
	"ecs":        true,
	"kubernetes": true,
	"nomad":      true,
	"docker":     true,
	"swarm":      true,
	"k3s":        true,
	"rke2":       true,
	"aks":        true,
	"gke":        true,
}

var allowedClusterStatuses = map[string]bool{
	"experimental": true,
	"development":  true,
	"staging":      true,
	"production":   true,
	"deprecated":   true,
}

func DefaultConfig(name string) *Config {
	cfg := &Config{
		ClusterForgeVersion: "0.1.0",
		Project: Project{
			Name: name,
		},
		Environments: map[string]Environment{},
		Audit: Audit{
			Enabled: true,
			Path:    ".cf/audit.log",
		},
		Policies: Policies{
			RequirePlanFileForApply:      true,
			BlockDestroyInProd:           true,
			RequireManualApprovalForProd: true,
		},
		Plugins: Plugins{
			Enabled:          false,
			AllowPathPlugins: true,
		},
	}
	cfg.ApplyDefaults()
	return cfg
}

func (c *Config) ApplyDefaults() {
	if c.ClusterForgeVersion == "" {
		c.ClusterForgeVersion = "0.1.0"
	}
	if c.Project.DefaultEngine == "" {
		c.Project.DefaultEngine = "terraform"
	}
	if c.Engines == nil {
		c.Engines = map[string]Engine{}
	}
	terraform := c.Engines["terraform"]
	if terraform.Binary == "" {
		terraform.Binary = "terraform"
	}
	c.Engines["terraform"] = terraform

	opentofu := c.Engines["opentofu"]
	if opentofu.Binary == "" {
		opentofu.Binary = "tofu"
	}
	c.Engines["opentofu"] = opentofu

	if c.Defaults.Cloud == "" {
		c.Defaults.Cloud = "aws"
	}
	if c.Defaults.Region == "" {
		c.Defaults.Region = "eu-central-1"
	}
	if c.Defaults.Orchestrator == "" {
		c.Defaults.Orchestrator = "eks"
	}
	if c.Audit.Enabled && c.Audit.Path == "" {
		c.Audit.Path = ".cf/audit.log"
	}
	if c.Environments == nil {
		c.Environments = map[string]Environment{}
	}
	if c.Clusters == nil {
		c.Clusters = map[string]Cluster{}
	}
	if c.Backends == nil {
		c.Backends = map[string]Backend{}
	}
	if c.Workspaces == nil {
		c.Workspaces = map[string]Workspace{}
	}
	if c.Teams == nil {
		c.Teams = map[string]Team{}
	}
	if c.AWSAccounts == nil {
		c.AWSAccounts = map[string]AWSAccount{}
	}
	if c.Regions == nil {
		c.Regions = map[string]string{}
	}
	if c.PlatformVersions == nil {
		c.PlatformVersions = map[string]string{}
	}
	if c.ExecutionProfiles == nil {
		c.ExecutionProfiles = map[string]ExecutionProfile{}
	}
	for name, backend := range c.Backends {
		if backend.Type == "" {
			backend.Type = "local"
		}
		c.Backends[name] = backend
	}
	for name, env := range c.Environments {
		if env.Cloud == "" {
			env.Cloud = c.Defaults.Cloud
		}
		if env.Region == "" {
			env.Region = c.Defaults.Region
		}
		if env.Orchestrator == "" {
			env.Orchestrator = c.Defaults.Orchestrator
		}
		if env.Layout == "" {
			env.Layout = "simple"
		}
		if env.Layout == "stacked" && env.Stacks == nil {
			env.Stacks = Stacks{}
			for _, stack := range StackOrder() {
				env.Stacks[stack] = Stack{Path: env.Path + "/" + stack}
			}
		}
		c.Environments[name] = env
	}
	for name, cluster := range c.Clusters {
		env := c.Environments[cluster.Environment]
		if cluster.Cloud == "" {
			cluster.Cloud = env.Cloud
		}
		if cluster.Region == "" {
			cluster.Region = env.Region
		}
		if cluster.Orchestrator == "" {
			cluster.Orchestrator = env.Orchestrator
		}
		if cluster.Path == "" {
			cluster.Path = env.Path
		}
		if cluster.Status == "" {
			cluster.Status = "experimental"
		}
		c.Clusters[name] = cluster
	}
}
