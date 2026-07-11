package app

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/textracta/clusterforge/cli/internal/bindings"
	"gopkg.in/yaml.v3"
)

const AppsDir = "apps"

var appNamePattern = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)

type Manifest struct {
	Name          string                      `yaml:"name"`
	Type          string                      `yaml:"type"`
	Image         string                      `yaml:"image"`
	Replicas      int                         `yaml:"replicas"`
	Ports         []Port                      `yaml:"ports,omitempty"`
	Env           map[string]string           `yaml:"env,omitempty"`
	SecretEnv     map[string]SecretRef        `yaml:"secret_env,omitempty"`
	Resources     Resources                   `yaml:"resources,omitempty"`
	Ingress       Ingress                     `yaml:"ingress,omitempty"`
	Autoscaling   Autoscaling                 `yaml:"autoscaling,omitempty"`
	CloudIdentity CloudIdentity               `yaml:"cloud_identity,omitempty"`
	Dependencies  map[string]bindings.Request `yaml:"dependencies,omitempty"`
	Backstage     Backstage                   `yaml:"backstage,omitempty"`
	Service       string                      `yaml:"service,omitempty"`
}
type Backstage struct {
	Owner     string `yaml:"owner,omitempty"`
	System    string `yaml:"system,omitempty"`
	Lifecycle string `yaml:"lifecycle,omitempty"`
}

type Port struct {
	Name          string `yaml:"name"`
	ContainerPort int    `yaml:"container_port"`
	Protocol      string `yaml:"protocol"`
}

type SecretRef struct {
	SecretName string `yaml:"secret_name"`
	SecretKey  string `yaml:"secret_key"`
}

type Resources struct {
	CPURequest    string `yaml:"cpu_request,omitempty"`
	MemoryRequest string `yaml:"memory_request,omitempty"`
	CPULimit      string `yaml:"cpu_limit,omitempty"`
	MemoryLimit   string `yaml:"memory_limit,omitempty"`
}

type Ingress struct {
	Enabled bool   `yaml:"enabled"`
	Host    string `yaml:"host,omitempty"`
	Path    string `yaml:"path,omitempty"`
	TLS     bool   `yaml:"tls,omitempty"`
}

type Autoscaling struct {
	Enabled     bool `yaml:"enabled"`
	MinReplicas int  `yaml:"min_replicas,omitempty"`
	MaxReplicas int  `yaml:"max_replicas,omitempty"`
	CPUPercent  int  `yaml:"cpu_percent,omitempty"`
}

// CloudIdentity describes an opt-in workload identity. Policy documents are
// references or JSON policy documents; credentials never belong here.
type CloudIdentity struct {
	Enabled        bool              `yaml:"enabled"`
	Provider       string            `yaml:"provider,omitempty"`
	PolicyARNs     []string          `yaml:"policy_arns,omitempty"`
	InlinePolicies map[string]string `yaml:"inline_policies,omitempty"`
	TaskRoleARN    string            `yaml:"task_role_arn,omitempty"`
}

type AddOptions struct {
	Force       bool
	Image       string
	Port        int
	Replicas    int
	Host        string
	Type        string
	Autoscaling bool
}

func NewManifest(name string, opts AddOptions) Manifest {
	if opts.Type == "" {
		opts.Type = "web"
	}
	if opts.Replicas == 0 {
		opts.Replicas = 1
	}
	manifest := Manifest{
		Name:     name,
		Type:     opts.Type,
		Image:    opts.Image,
		Replicas: opts.Replicas,
		Env:      map[string]string{},
	}
	if opts.Port > 0 {
		manifest.Ports = []Port{{
			Name:          "http",
			ContainerPort: opts.Port,
			Protocol:      "TCP",
		}}
	}
	if opts.Host != "" {
		manifest.Ingress = Ingress{
			Enabled: true,
			Host:    opts.Host,
			Path:    "/",
			TLS:     true,
		}
	}
	if opts.Autoscaling {
		manifest.Autoscaling = Autoscaling{
			Enabled:     true,
			MinReplicas: max(1, opts.Replicas),
			MaxReplicas: max(3, opts.Replicas),
			CPUPercent:  70,
		}
	}
	return manifest
}

func Add(rootDir, name string, opts AddOptions) (string, error) {
	manifest := NewManifest(name, opts)
	if err := manifest.Validate(); err != nil {
		return "", err
	}
	path := ManifestPath(rootDir, name)
	if !opts.Force {
		if _, err := os.Stat(path); err == nil {
			return "", fmt.Errorf("%s already exists; use --force to overwrite it", path)
		}
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", fmt.Errorf("create apps directory: %w", err)
	}
	if err := Save(path, manifest); err != nil {
		return "", err
	}
	return path, nil
}

func Load(path string) (Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Manifest{}, fmt.Errorf("read app manifest %s: %w", path, err)
	}
	if err := validateRawManifestYAML(data); err != nil {
		return Manifest{}, err
	}
	var manifest Manifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return Manifest{}, fmt.Errorf("parse app manifest %s: %w", path, err)
	}
	manifest.ApplyDefaults()
	if err := manifest.Validate(); err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}

func ValidateFile(path string) error {
	_, err := Load(path)
	return err
}

func ImagePolicyWarnings(manifest Manifest, environment string) []string {
	image := strings.TrimSpace(manifest.Image)
	if image == "" {
		return nil
	}
	var warnings []string
	if imageUsesLatestTag(image) {
		warnings = append(warnings, fmt.Sprintf("image %q uses the latest tag; use a version tag or digest for repeatable deploys", image))
	}
	if isProdEnvironment(environment) && !imageIsPinned(image) {
		warnings = append(warnings, fmt.Sprintf("prod image %q is not pinned by tag or digest", image))
	}
	return warnings
}

func Save(path string, manifest Manifest) error {
	manifest.ApplyDefaults()
	if err := manifest.Validate(); err != nil {
		return err
	}
	data, err := yaml.Marshal(manifest)
	if err != nil {
		return fmt.Errorf("encode app manifest: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write app manifest %s: %w", path, err)
	}
	return nil
}

func List(rootDir string) ([]string, error) {
	dir := filepath.Join(rootDir, AppsDir)
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("read apps directory: %w", err)
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".yaml" {
			continue
		}
		names = append(names, strings.TrimSuffix(entry.Name(), ".yaml"))
	}
	sort.Strings(names)
	return names, nil
}

func Remove(rootDir, name string) error {
	path := ManifestPath(rootDir, name)
	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("app %q does not exist", name)
		}
		return fmt.Errorf("remove app manifest %s: %w", path, err)
	}
	return nil
}

func ManifestPath(rootDir, name string) string {
	return filepath.Join(rootDir, AppsDir, name+".yaml")
}

func (m *Manifest) ApplyDefaults() {
	if m.Type == "" {
		m.Type = "web"
	}
	if m.Replicas == 0 {
		m.Replicas = 1
	}
	if m.Env == nil {
		m.Env = map[string]string{}
	}
	if m.SecretEnv == nil {
		m.SecretEnv = map[string]SecretRef{}
	}
	if m.CloudIdentity.InlinePolicies == nil {
		m.CloudIdentity.InlinePolicies = map[string]string{}
	}
	if m.Dependencies == nil {
		m.Dependencies = map[string]bindings.Request{}
	}
	if m.CloudIdentity.Enabled && m.CloudIdentity.Provider == "" {
		m.CloudIdentity.Provider = "aws"
	}
	for i := range m.Ports {
		if m.Ports[i].Protocol == "" {
			m.Ports[i].Protocol = "TCP"
		}
	}
	if m.Ingress.Enabled && m.Ingress.Path == "" {
		m.Ingress.Path = "/"
	}
	if m.Autoscaling.Enabled {
		if m.Autoscaling.MinReplicas == 0 {
			m.Autoscaling.MinReplicas = 1
		}
		if m.Autoscaling.MaxReplicas == 0 {
			m.Autoscaling.MaxReplicas = 3
		}
		if m.Autoscaling.CPUPercent == 0 {
			m.Autoscaling.CPUPercent = 70
		}
	}
}

func (m Manifest) Validate() error {
	var errs ValidationErrors
	if strings.TrimSpace(m.Name) == "" {
		errs.Add("name is required")
	} else if len(m.Name) > 63 || !appNamePattern.MatchString(m.Name) {
		errs.Add("name must be DNS/Kubernetes-name compatible")
	}
	if strings.TrimSpace(m.Type) == "" {
		errs.Add("type is required")
	} else if !isAllowedAppType(m.Type) {
		errs.Add("type must be one of web, worker, cronjob, service")
	}
	if strings.TrimSpace(m.Image) == "" {
		errs.Add("image is required")
	}
	if m.Replicas < 0 {
		errs.Add("replicas must be greater than or equal to 0")
	}
	for index, port := range m.Ports {
		if strings.TrimSpace(port.Name) == "" {
			errs.Add(fmt.Sprintf("ports[%d].name is required", index))
		}
		if port.ContainerPort <= 0 || port.ContainerPort > 65535 {
			errs.Add(fmt.Sprintf("ports[%d].container_port must be between 1 and 65535", index))
		}
		if !isAllowedProtocol(port.Protocol) {
			errs.Add(fmt.Sprintf("ports[%d].protocol must be TCP or UDP", index))
		}
	}
	if m.Ingress.Enabled && strings.TrimSpace(m.Ingress.Host) == "" {
		errs.Add("ingress.host is required when ingress.enabled=true")
	}
	if m.Ingress.Path != "" && !strings.HasPrefix(m.Ingress.Path, "/") {
		errs.Add("ingress.path must start with /")
	}
	if m.Autoscaling.Enabled && m.Autoscaling.MinReplicas > m.Autoscaling.MaxReplicas {
		errs.Add("autoscaling.min_replicas must be less than or equal to max_replicas")
	}
	if m.Autoscaling.Enabled && m.Autoscaling.MaxReplicas <= 0 {
		errs.Add("autoscaling.max_replicas must be greater than 0")
	}
	for key, ref := range m.SecretEnv {
		if strings.TrimSpace(ref.SecretName) == "" {
			errs.Add(fmt.Sprintf("secret_env.%s.secret_name is required", key))
		}
		if strings.TrimSpace(ref.SecretKey) == "" {
			errs.Add(fmt.Sprintf("secret_env.%s.secret_key is required", key))
		}
	}
	if m.CloudIdentity.Enabled {
		if strings.ToLower(strings.TrimSpace(m.CloudIdentity.Provider)) != "aws" {
			errs.Add("cloud_identity.provider must be aws for the current MVP")
		}
		for index, arn := range m.CloudIdentity.PolicyARNs {
			if !strings.HasPrefix(strings.TrimSpace(arn), "arn:aws:iam::") {
				errs.Add(fmt.Sprintf("cloud_identity.policy_arns[%d] must be an AWS IAM policy ARN", index))
			}
			if strings.Contains(arn, "AdministratorAccess") {
				errs.Add(fmt.Sprintf("cloud_identity.policy_arns[%d] must not grant administrator access", index))
			}
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

type ValidationErrors []string

func (v *ValidationErrors) Add(message string) {
	*v = append(*v, message)
}

func (v ValidationErrors) Error() string {
	return strings.Join(v, "\n")
}

func isAllowedAppType(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "web", "worker", "cronjob", "service":
		return true
	default:
		return false
	}
}

func isAllowedProtocol(value string) bool {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "TCP", "UDP":
		return true
	default:
		return false
	}
}

func imageUsesLatestTag(image string) bool {
	tag, ok := imageTag(image)
	return ok && tag == "latest"
}

func imageIsPinned(image string) bool {
	if strings.Contains(image, "@sha256:") {
		return true
	}
	_, ok := imageTag(image)
	return ok
}

func imageTag(image string) (string, bool) {
	if strings.Contains(image, "@") {
		return "", false
	}
	lastSlash := strings.LastIndex(image, "/")
	lastColon := strings.LastIndex(image, ":")
	if lastColon <= lastSlash {
		return "", false
	}
	tag := strings.TrimSpace(image[lastColon+1:])
	return tag, tag != ""
}

func isProdEnvironment(environment string) bool {
	env := strings.ToLower(strings.TrimSpace(environment))
	return env == "prod" || env == "production"
}

func validateRawManifestYAML(data []byte) error {
	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return fmt.Errorf("parse app manifest: %w", err)
	}
	if len(root.Content) == 0 {
		return nil
	}
	doc := root.Content[0]
	if doc.Kind != yaml.MappingNode {
		return ValidationErrors{"manifest must be a YAML mapping"}
	}
	var errs ValidationErrors
	for i := 0; i < len(doc.Content)-1; i += 2 {
		key := doc.Content[i].Value
		value := doc.Content[i+1]
		switch key {
		case "ingress":
			validateIngressNode(value, &errs)
		case "ports":
			validatePortsNode(value, &errs)
		case "resources":
			validateResourceNode(value, &errs)
		case "autoscaling":
			validateAutoscalingNode(value, &errs)
		case "secret_env":
			validateSecretEnvNode(value, &errs)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func validateIngressNode(node *yaml.Node, errs *ValidationErrors) {
	if node.Kind != yaml.MappingNode {
		errs.Add("ingress must be a mapping")
		return
	}
	for i := 0; i < len(node.Content)-1; i += 2 {
		key := node.Content[i].Value
		value := node.Content[i+1]
		if key == "path" && !strings.HasPrefix(value.Value, "/") {
			errs.Add("ingress.path must start with /")
		}
	}
}

func validatePortsNode(node *yaml.Node, errs *ValidationErrors) {
	if node.Kind != yaml.SequenceNode {
		errs.Add("ports must be a list")
		return
	}
	for index, port := range node.Content {
		if port.Kind != yaml.MappingNode {
			errs.Add(fmt.Sprintf("ports[%d] must be a mapping", index))
			continue
		}
		for i := 0; i < len(port.Content)-1; i += 2 {
			key := port.Content[i].Value
			value := port.Content[i+1]
			if key == "protocol" && strings.TrimSpace(value.Value) != "" && !isAllowedProtocol(value.Value) {
				errs.Add(fmt.Sprintf("ports[%d].protocol must be TCP or UDP", index))
			}
		}
	}
}

func validateResourceNode(node *yaml.Node, errs *ValidationErrors) {
	if node.Kind != yaml.MappingNode {
		errs.Add("resources must be a mapping")
		return
	}
	allowed := map[string]bool{
		"cpu_request":    true,
		"memory_request": true,
		"cpu_limit":      true,
		"memory_limit":   true,
	}
	for i := 0; i < len(node.Content)-1; i += 2 {
		key := node.Content[i].Value
		value := node.Content[i+1]
		if !allowed[key] {
			continue
		}
		if strings.TrimSpace(value.Value) == "" {
			errs.Add(fmt.Sprintf("resources.%s must be a non-empty string when provided", key))
		}
	}
}

func validateAutoscalingNode(node *yaml.Node, errs *ValidationErrors) {
	if node.Kind != yaml.MappingNode {
		errs.Add("autoscaling must be a mapping")
		return
	}
	for i := 0; i < len(node.Content)-1; i += 2 {
		key := node.Content[i].Value
		value := node.Content[i+1]
		if key != "max_replicas" {
			continue
		}
		if value.Value == "" {
			errs.Add("autoscaling.max_replicas must be greater than 0")
			continue
		}
		var parsed int
		if err := value.Decode(&parsed); err != nil || parsed <= 0 {
			errs.Add("autoscaling.max_replicas must be greater than 0")
		}
	}
}

func validateSecretEnvNode(node *yaml.Node, errs *ValidationErrors) {
	if node.Kind != yaml.MappingNode {
		errs.Add("secret_env must be a mapping")
		return
	}
	for i := 0; i < len(node.Content)-1; i += 2 {
		envName := node.Content[i].Value
		ref := node.Content[i+1]
		if ref.Kind != yaml.MappingNode {
			errs.Add(fmt.Sprintf("secret_env.%s must reference secret_name and secret_key", envName))
			continue
		}
		for j := 0; j < len(ref.Content)-1; j += 2 {
			key := ref.Content[j].Value
			value := ref.Content[j+1]
			switch key {
			case "secret_name", "secret_key":
				if strings.TrimSpace(value.Value) == "" {
					errs.Add(fmt.Sprintf("secret_env.%s.%s is required", envName, key))
				}
			default:
				errs.Add(fmt.Sprintf("secret_env.%s.%s is not allowed; use secret_name and secret_key references only", envName, key))
			}
		}
	}
}
