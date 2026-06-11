package config

var allowedClouds = map[string]bool{
	"aws":     true,
	"azure":   true,
	"gcp":     true,
	"hetzner": true,
	"local":   true,
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

func DefaultConfig(name string) *Config {
	cfg := &Config{
		Project: Project{
			Name: name,
		},
		Environments: map[string]Environment{},
		Policies: Policies{
			RequirePlanFileForApply:      true,
			BlockDestroyInProd:           true,
			RequireManualApprovalForProd: true,
		},
	}
	cfg.ApplyDefaults()
	return cfg
}

func (c *Config) ApplyDefaults() {
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
	if c.Environments == nil {
		c.Environments = map[string]Environment{}
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
		c.Environments[name] = env
	}
}
