package config

import "fmt"

const DefaultPath = "clusterforge.yaml"

type Config struct {
	Project      Project                `yaml:"project"`
	Engines      map[string]Engine      `yaml:"engines"`
	Defaults     Defaults               `yaml:"defaults"`
	Environments map[string]Environment `yaml:"environments"`
	Policies     Policies               `yaml:"policies"`
}

type Project struct {
	Name          string `yaml:"name"`
	DefaultEngine string `yaml:"default_engine"`
}

type Engine struct {
	Binary string `yaml:"binary"`
}

type Defaults struct {
	Cloud        string `yaml:"cloud"`
	Region       string `yaml:"region"`
	Orchestrator string `yaml:"orchestrator"`
}

type Environment struct {
	Cloud        string `yaml:"cloud"`
	Region       string `yaml:"region"`
	Orchestrator string `yaml:"orchestrator"`
	Path         string `yaml:"path"`
}

type Policies struct {
	RequirePlanFileForApply      bool `yaml:"require_plan_file_for_apply"`
	BlockDestroyInProd           bool `yaml:"block_destroy_in_prod"`
	RequireManualApprovalForProd bool `yaml:"require_manual_approval_for_prod"`
}

func DefaultConfig(name string) *Config {
	return &Config{
		Project: Project{
			Name:          name,
			DefaultEngine: "terraform",
		},
		Engines: map[string]Engine{
			"terraform": {Binary: "terraform"},
			"opentofu":  {Binary: "tofu"},
		},
		Defaults: Defaults{
			Cloud:        "aws",
			Region:       "eu-central-1",
			Orchestrator: "eks",
		},
		Environments: map[string]Environment{
			"dev": {
				Cloud:        "aws",
				Region:       "eu-central-1",
				Orchestrator: "eks",
				Path:         "live/dev/aws-eks",
			},
		},
		Policies: Policies{
			RequirePlanFileForApply:      true,
			BlockDestroyInProd:           true,
			RequireManualApprovalForProd: true,
		},
	}
}

func (c *Config) Validate() error {
	if c.Project.Name == "" {
		return fmt.Errorf("project.name is required")
	}
	if c.Project.DefaultEngine == "" {
		c.Project.DefaultEngine = "terraform"
	}
	if c.Engines == nil {
		c.Engines = map[string]Engine{}
	}
	if c.Environments == nil {
		c.Environments = map[string]Environment{}
	}
	for name, env := range c.Environments {
		if env.Path == "" {
			return fmt.Errorf("environment %q path is required", name)
		}
	}
	return nil
}

func (c *Config) EngineBinary(name string) (string, error) {
	switch name {
	case "terraform":
		name = "terraform"
	case "tofu", "opentofu":
		name = "opentofu"
	}
	engine, ok := c.Engines[name]
	if !ok {
		return "", fmt.Errorf("engine %q is not configured", name)
	}
	if engine.Binary == "" {
		return "", fmt.Errorf("engine %q binary is empty", name)
	}
	return engine.Binary, nil
}
