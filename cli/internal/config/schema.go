package config

import (
	"fmt"
	"strings"
)

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
	Layout       string `yaml:"layout,omitempty"`
	Stacks       Stacks `yaml:"stacks,omitempty"`
}

type Stacks map[string]Stack

type Stack struct {
	Path string `yaml:"path"`
}

type Policies struct {
	RequirePlanFileForApply      bool `yaml:"require_plan_file_for_apply"`
	BlockDestroyInProd           bool `yaml:"block_destroy_in_prod"`
	RequireManualApprovalForProd bool `yaml:"require_manual_approval_for_prod"`
}

func (c *Config) Validate() error {
	if strings.TrimSpace(c.Project.Name) == "" {
		return fmt.Errorf("project.name is required")
	}
	if _, ok := c.Engines[c.Project.DefaultEngine]; !ok {
		return fmt.Errorf("project.default_engine %q must exist in engines", c.Project.DefaultEngine)
	}
	if err := validateCloud("defaults.cloud", c.Defaults.Cloud); err != nil {
		return err
	}
	if err := validateOrchestrator("defaults.orchestrator", c.Defaults.Orchestrator); err != nil {
		return err
	}
	for name, env := range c.Environments {
		if strings.TrimSpace(name) == "" {
			return fmt.Errorf("environment name must not be empty")
		}
		if strings.TrimSpace(env.Path) == "" {
			return fmt.Errorf("environment %q path is required", name)
		}
		layout := env.Layout
		if layout == "" {
			layout = "simple"
		}
		if layout != "simple" && layout != "stacked" {
			return fmt.Errorf("environment %q layout must be simple or stacked", name)
		}
		if layout == "stacked" {
			for _, stackName := range StackOrder() {
				stack, ok := env.Stacks[stackName]
				if !ok || strings.TrimSpace(stack.Path) == "" {
					return fmt.Errorf("environment %q stack %q path is required", name, stackName)
				}
			}
		}
		if err := validateCloud(fmt.Sprintf("environment %q cloud", name), env.Cloud); err != nil {
			return err
		}
		if err := validateOrchestrator(fmt.Sprintf("environment %q orchestrator", name), env.Orchestrator); err != nil {
			return err
		}
	}
	return nil
}

func StackOrder() []string {
	return []string{"network", "cluster", "platform", "apps"}
}

func (e Environment) EffectiveLayout() string {
	if e.Layout == "" {
		return "simple"
	}
	return e.Layout
}

func (e Environment) StackPaths(stack string) ([]string, error) {
	if e.EffectiveLayout() == "simple" {
		if stack != "" {
			return nil, fmt.Errorf("environment uses simple layout; stack %q is not available", stack)
		}
		return []string{e.Path}, nil
	}
	if stack != "" {
		resolved, ok := e.Stacks[stack]
		if !ok || strings.TrimSpace(resolved.Path) == "" {
			return nil, fmt.Errorf("unknown stack %q; expected one of network, cluster, platform, apps", stack)
		}
		return []string{resolved.Path}, nil
	}
	paths := make([]string, 0, len(StackOrder()))
	for _, name := range StackOrder() {
		resolved, ok := e.Stacks[name]
		if !ok || strings.TrimSpace(resolved.Path) == "" {
			return nil, fmt.Errorf("stack %q path is required", name)
		}
		paths = append(paths, resolved.Path)
	}
	return paths, nil
}

func (e Environment) StackPath(stack string) (string, error) {
	paths, err := e.StackPaths(stack)
	if err != nil {
		return "", err
	}
	if len(paths) != 1 {
		return "", fmt.Errorf("stack must be specified")
	}
	return paths[0], nil
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

func validateCloud(field, value string) error {
	if allowedClouds[strings.ToLower(strings.TrimSpace(value))] {
		return nil
	}
	return fmt.Errorf("%s must be one of aws, azure, gcp, hetzner, local", field)
}

func validateOrchestrator(field, value string) error {
	if allowedOrchestrators[strings.ToLower(strings.TrimSpace(value))] {
		return nil
	}
	return fmt.Errorf("%s must be one of eks, ecs, kubernetes, nomad, docker, swarm, k3s, rke2, aks, gke", field)
}
