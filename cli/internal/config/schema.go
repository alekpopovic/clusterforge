package config

import (
	"fmt"
	"strings"
)

const DefaultPath = "clusterforge.yaml"

type Config struct {
	ClusterForgeVersion       string                 `yaml:"clusterforge_version,omitempty"`
	Project                   Project                `yaml:"project"`
	Engines                   map[string]Engine      `yaml:"engines"`
	Defaults                  Defaults               `yaml:"defaults"`
	Environments              map[string]Environment `yaml:"environments"`
	Clusters                  map[string]Cluster     `yaml:"clusters,omitempty"`
	Audit                     Audit                  `yaml:"audit,omitempty"`
	Backends                  map[string]Backend     `yaml:"backends,omitempty"`
	Policies                  Policies               `yaml:"policies"`
	TemplatePacks             []TemplatePack         `yaml:"template_packs,omitempty"`
	Plugins                   Plugins                `yaml:"plugins,omitempty"`
	Organization              Organization           `yaml:"organization,omitempty"`
	Workspaces                map[string]Workspace   `yaml:"workspaces,omitempty"`
	Teams                     map[string]Team        `yaml:"teams,omitempty"`
	AWSAccounts               map[string]AWSAccount  `yaml:"aws_accounts,omitempty"`
	AllowSharedProdAWSAccount bool                   `yaml:"allow_shared_prod_aws_account,omitempty"`
	Regions                   map[string]string      `yaml:"regions,omitempty"`
	PlatformVersions          map[string]string      `yaml:"platform_versions,omitempty"`
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
	Cloud             string `yaml:"cloud"`
	Region            string `yaml:"region"`
	Orchestrator      string `yaml:"orchestrator"`
	Path              string `yaml:"path"`
	Layout            string `yaml:"layout,omitempty"`
	Stacks            Stacks `yaml:"stacks,omitempty"`
	Account           string `yaml:"account,omitempty"`
	KubernetesVersion string `yaml:"kubernetes_version,omitempty"`
}

type AWSAccount struct {
	AccountID string `yaml:"account_id"`
	Region    string `yaml:"region"`
	Profile   string `yaml:"profile,omitempty"`
	RoleARN   string `yaml:"role_arn,omitempty"`
}

type Cluster struct {
	Environment       string `yaml:"environment"`
	Cloud             string `yaml:"cloud"`
	Orchestrator      string `yaml:"orchestrator"`
	Region            string `yaml:"region"`
	Path              string `yaml:"path"`
	Status            string `yaml:"status,omitempty"`
	KubeconfigPath    string `yaml:"kubeconfig_path,omitempty"`
	KubernetesVersion string `yaml:"kubernetes_version,omitempty"`
	NodeVersion       string `yaml:"node_version,omitempty"`
}

type Audit struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path,omitempty"`
}

type Stacks map[string]Stack

type Stack struct {
	Path string `yaml:"path"`
}

type Backend struct {
	Type          string `yaml:"type"`
	Bucket        string `yaml:"bucket,omitempty"`
	Region        string `yaml:"region,omitempty"`
	DynamoDBTable string `yaml:"dynamodb_table,omitempty"`
	KeyPrefix     string `yaml:"key_prefix,omitempty"`
}

type Policies struct {
	RequirePlanFileForApply      bool `yaml:"require_plan_file_for_apply"`
	BlockDestroyInProd           bool `yaml:"block_destroy_in_prod"`
	RequireManualApprovalForProd bool `yaml:"require_manual_approval_for_prod"`
}

type TemplatePack struct {
	Name    string `yaml:"name"`
	Path    string `yaml:"path,omitempty"`
	Source  string `yaml:"source,omitempty"`
	Version string `yaml:"version,omitempty"`
	Enabled *bool  `yaml:"enabled,omitempty"`
}

func (p TemplatePack) IsEnabled() bool { return p.Enabled == nil || *p.Enabled }

type Plugins struct {
	Enabled          bool     `yaml:"enabled"`
	Directories      []string `yaml:"directories,omitempty"`
	AllowPathPlugins bool     `yaml:"allow_path_plugins"`
	Disabled         []string `yaml:"disabled,omitempty"`
}

type Organization struct {
	Name    string `yaml:"name,omitempty"`
	Owner   string `yaml:"owner,omitempty"`
	Contact string `yaml:"contact,omitempty"`
}
type Workspace struct {
	Description  string   `yaml:"description,omitempty"`
	Environments []string `yaml:"environments,omitempty"`
}
type Team struct {
	Owners     []string `yaml:"owners,omitempty"`
	Namespaces []string `yaml:"namespaces,omitempty"`
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
		if env.Cloud == "aws" && env.Account != "" {
			if _, ok := c.AWSAccounts[env.Account]; !ok {
				return fmt.Errorf("environment %q references unknown AWS account %q", name, env.Account)
			}
		}
	}
	for name, account := range c.AWSAccounts {
		if strings.TrimSpace(name) == "" || len(account.AccountID) != 12 {
			return fmt.Errorf("AWS account %q requires a 12-digit account_id", name)
		}
		for _, char := range account.AccountID {
			if char < '0' || char > '9' {
				return fmt.Errorf("AWS account %q requires a 12-digit account_id", name)
			}
		}
		if account.AccountID == "000000000000" {
			return fmt.Errorf("AWS account %q must not use a placeholder/root account ID", name)
		}
	}
	seenRegions := map[string]string{}
	for alias, region := range c.Regions {
		if strings.TrimSpace(alias) == "" || strings.TrimSpace(region) == "" {
			return fmt.Errorf("region alias and value must not be empty")
		}
		if previous := seenRegions[region]; previous != "" {
			return fmt.Errorf("region aliases %q and %q both map to %q", previous, alias, region)
		}
		seenRegions[region] = alias
	}
	for name, cluster := range c.Clusters {
		if strings.TrimSpace(name) == "" {
			return fmt.Errorf("cluster name must not be empty")
		}
		if strings.TrimSpace(cluster.Environment) == "" {
			return fmt.Errorf("cluster %q environment is required", name)
		}
		if _, ok := c.Environments[cluster.Environment]; !ok {
			return fmt.Errorf("cluster %q references unknown environment %q", name, cluster.Environment)
		}
		if strings.TrimSpace(cluster.Path) == "" {
			return fmt.Errorf("cluster %q path is required", name)
		}
		if err := validateCloud(fmt.Sprintf("cluster %q cloud", name), cluster.Cloud); err != nil {
			return err
		}
		if err := validateOrchestrator(fmt.Sprintf("cluster %q orchestrator", name), cluster.Orchestrator); err != nil {
			return err
		}
		if !allowedClusterStatuses[strings.ToLower(strings.TrimSpace(cluster.Status))] {
			return fmt.Errorf("cluster %q status must be one of experimental, development, staging, production, deprecated", name)
		}
	}
	for name, workspace := range c.Workspaces {
		if strings.TrimSpace(name) == "" {
			return fmt.Errorf("workspace name must not be empty")
		}
		seen := map[string]bool{}
		for _, environment := range workspace.Environments {
			if seen[environment] {
				return fmt.Errorf("workspace %q contains duplicate environment %q", name, environment)
			}
			if _, ok := c.Environments[environment]; !ok {
				return fmt.Errorf("workspace %q references unknown environment %q", name, environment)
			}
			seen[environment] = true
		}
	}
	for name, team := range c.Teams {
		if strings.TrimSpace(name) == "" {
			return fmt.Errorf("team name must not be empty")
		}
		if len(team.Owners) == 0 {
			return fmt.Errorf("team %q requires at least one owner", name)
		}
	}
	for name, backend := range c.Backends {
		if strings.TrimSpace(name) == "" {
			return fmt.Errorf("backend environment name must not be empty")
		}
		if err := backend.Validate(name); err != nil {
			return err
		}
	}
	seenPacks := map[string]bool{}
	for _, pack := range c.TemplatePacks {
		if strings.TrimSpace(pack.Name) == "" {
			return fmt.Errorf("template pack name is required")
		}
		if strings.TrimSpace(pack.Path) == "" && strings.TrimSpace(pack.Source) == "" {
			return fmt.Errorf("template pack %q requires path or source", pack.Name)
		}
		if strings.TrimSpace(pack.Source) != "" && strings.TrimSpace(pack.Version) == "" {
			return fmt.Errorf("template pack %q version is required for source", pack.Name)
		}
		if seenPacks[pack.Name] {
			return fmt.Errorf("duplicate template pack %q", pack.Name)
		}
		seenPacks[pack.Name] = true
	}
	return nil
}

func (b Backend) Validate(environment string) error {
	backendType := strings.ToLower(strings.TrimSpace(b.Type))
	if backendType == "" {
		backendType = "local"
	}
	switch backendType {
	case "local", "azurerm", "gcs":
		return nil
	case "s3":
		if strings.TrimSpace(b.Bucket) == "" {
			return fmt.Errorf("backends.%s.bucket is required for s3 backend", environment)
		}
		if strings.TrimSpace(b.Region) == "" {
			return fmt.Errorf("backends.%s.region is required for s3 backend", environment)
		}
		return nil
	default:
		return fmt.Errorf("backends.%s.type must be one of local, s3, azurerm, gcs", environment)
	}
}

func (b Backend) EffectiveType() string {
	if strings.TrimSpace(b.Type) == "" {
		return "local"
	}
	return strings.ToLower(strings.TrimSpace(b.Type))
}

func (c *Config) BackendFor(environment string) Backend {
	if c.Backends == nil {
		return Backend{Type: "local"}
	}
	backend, ok := c.Backends[environment]
	if !ok {
		return Backend{Type: "local"}
	}
	if backend.Type == "" {
		backend.Type = "local"
	}
	return backend
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
	return fmt.Errorf("%s must be one of aws, azure, gcp, hetzner, local, existing", field)
}

func validateOrchestrator(field, value string) error {
	if allowedOrchestrators[strings.ToLower(strings.TrimSpace(value))] {
		return nil
	}
	return fmt.Errorf("%s must be one of eks, ecs, kubernetes, nomad, docker, swarm, k3s, rke2, aks, gke", field)
}
