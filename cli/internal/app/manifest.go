package app

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const AppsDir = "apps"

type Manifest struct {
	Name        string               `yaml:"name"`
	Type        string               `yaml:"type"`
	Image       string               `yaml:"image"`
	Replicas    int                  `yaml:"replicas"`
	Ports       []Port               `yaml:"ports"`
	Env         map[string]string    `yaml:"env"`
	SecretEnv   map[string]SecretRef `yaml:"secret_env"`
	Resources   Resources            `yaml:"resources"`
	Ingress     Ingress              `yaml:"ingress"`
	Autoscaling Autoscaling          `yaml:"autoscaling"`
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
	CPURequest    string `yaml:"cpu_request"`
	MemoryRequest string `yaml:"memory_request"`
	CPULimit      string `yaml:"cpu_limit"`
	MemoryLimit   string `yaml:"memory_limit"`
}

type Ingress struct {
	Enabled bool   `yaml:"enabled"`
	Host    string `yaml:"host"`
	Path    string `yaml:"path"`
	TLS     bool   `yaml:"tls"`
}

type Autoscaling struct {
	Enabled     bool `yaml:"enabled"`
	MinReplicas int  `yaml:"min_replicas"`
	MaxReplicas int  `yaml:"max_replicas"`
	CPUPercent  int  `yaml:"cpu_percent"`
}

type AddOptions struct {
	Force    bool
	Image    string
	Port     int
	Replicas int
	Host     string
	Type     string
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
	if strings.TrimSpace(m.Name) == "" {
		return fmt.Errorf("app name is required")
	}
	if strings.TrimSpace(m.Type) == "" {
		return fmt.Errorf("app type is required")
	}
	if strings.TrimSpace(m.Image) == "" {
		return fmt.Errorf("app image is required")
	}
	if m.Replicas < 0 {
		return fmt.Errorf("app replicas must be greater than or equal to 0")
	}
	for _, port := range m.Ports {
		if strings.TrimSpace(port.Name) == "" {
			return fmt.Errorf("app port name is required")
		}
		if port.ContainerPort <= 0 || port.ContainerPort > 65535 {
			return fmt.Errorf("app port %q container_port must be between 1 and 65535", port.Name)
		}
	}
	if m.Ingress.Enabled && strings.TrimSpace(m.Ingress.Host) == "" {
		return fmt.Errorf("ingress.host is required when ingress is enabled")
	}
	if m.Autoscaling.Enabled && m.Autoscaling.MinReplicas > m.Autoscaling.MaxReplicas {
		return fmt.Errorf("autoscaling.min_replicas must be less than or equal to max_replicas")
	}
	return nil
}
