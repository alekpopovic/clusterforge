package servicecatalog

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"net/url"
	"os"
	"sort"
	"strings"
)

const DefaultPath = "service-catalog.yaml"

type Catalog struct {
	Services map[string]Service `yaml:"services" json:"services"`
}
type Service struct {
	Owner        string                 `yaml:"owner" json:"owner"`
	Tier         string                 `yaml:"tier" json:"tier"`
	Lifecycle    string                 `yaml:"lifecycle" json:"lifecycle"`
	Repositories Repositories           `yaml:"repositories,omitempty" json:"repositories,omitempty"`
	Environments map[string]Environment `yaml:"environments,omitempty" json:"environments,omitempty"`
	Dependencies []string               `yaml:"dependencies,omitempty" json:"dependencies,omitempty"`
	Alerts       Alerts                 `yaml:"alerts,omitempty" json:"alerts,omitempty"`
}
type Repositories struct {
	Source string `yaml:"source,omitempty" json:"source,omitempty"`
	Image  string `yaml:"image,omitempty" json:"image,omitempty"`
}
type Environment struct {
	URL string `yaml:"url,omitempty" json:"url,omitempty"`
}
type Alerts struct {
	Dashboard string `yaml:"dashboard,omitempty" json:"dashboard,omitempty"`
	Runbook   string `yaml:"runbook,omitempty" json:"runbook,omitempty"`
}

func Load(path string) (Catalog, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Catalog{}, fmt.Errorf("read service catalog: %w", err)
	}
	var c Catalog
	if err := yaml.Unmarshal(data, &c); err != nil {
		return Catalog{}, fmt.Errorf("parse service catalog: %w", err)
	}
	if err := c.Validate(); err != nil {
		return Catalog{}, err
	}
	return c, nil
}
func (c Catalog) Validate() error {
	if len(c.Services) == 0 {
		return fmt.Errorf("services must not be empty")
	}
	for name, s := range c.Services {
		if strings.TrimSpace(name) == "" || strings.TrimSpace(s.Owner) == "" || strings.TrimSpace(s.Lifecycle) == "" {
			return fmt.Errorf("service %q requires name, owner, and lifecycle", name)
		}
		for env, e := range s.Environments {
			if e.URL != "" {
				u, err := url.ParseRequestURI(e.URL)
				if err != nil || u.Scheme == "" || u.Host == "" {
					return fmt.Errorf("service %q environment %q URL is invalid", name, env)
				}
			}
		}
		for _, dep := range s.Dependencies {
			if dep == name {
				return fmt.Errorf("service %q depends on itself", name)
			}
		}
		raw, _ := yaml.Marshal(s)
		lower := strings.ToLower(string(raw))
		for _, word := range []string{"password:", "secret:", "token:", "api_key:"} {
			if strings.Contains(lower, word) {
				return fmt.Errorf("service %q contains prohibited secret-like metadata", name)
			}
		}
	}
	for name, service := range c.Services {
		for _, dependency := range service.Dependencies {
			if _, ok := c.Services[dependency]; !ok {
				return fmt.Errorf("service %q references unknown dependency %q", name, dependency)
			}
		}
	}
	return nil
}
func (c Catalog) Names() []string {
	names := make([]string, 0, len(c.Services))
	for name := range c.Services {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
func (c Catalog) JSON() (byteData []byte, err error) { return json.MarshalIndent(c, "", "  ") }
func (c Catalog) WriteMarkdown(out io.Writer) error {
	fmt.Fprintln(out, "| Service | Owner | Tier | Lifecycle | Dependencies | Runbook |\n| --- | --- | --- | --- | --- | --- |")
	for _, name := range c.Names() {
		s := c.Services[name]
		fmt.Fprintf(out, "| %s | %s | %s | %s | %s | %s |\n", name, s.Owner, s.Tier, s.Lifecycle, strings.Join(s.Dependencies, ", "), s.Alerts.Runbook)
	}
	return nil
}
func (c Catalog) DOT() string {
	var b strings.Builder
	b.WriteString("digraph services {\n")
	for _, name := range c.Names() {
		fmt.Fprintf(&b, "  %q;\n", name)
		deps := append([]string{}, c.Services[name].Dependencies...)
		sort.Strings(deps)
		for _, dep := range deps {
			fmt.Fprintf(&b, "  %q -> %q;\n", name, dep)
		}
	}
	b.WriteString("}\n")
	return b.String()
}
