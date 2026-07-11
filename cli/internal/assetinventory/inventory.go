package assetinventory

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Asset struct {
	Address     string            `json:"address"`
	Type        string            `json:"type"`
	Name        string            `json:"name"`
	Provider    string            `json:"provider,omitempty"`
	Module      string            `json:"module,omitempty"`
	Environment string            `json:"environment"`
	Stack       string            `json:"stack,omitempty"`
	Cloud       string            `json:"cloud"`
	Region      string            `json:"region"`
	Tags        map[string]string `json:"tags,omitempty"`
	Sensitive   bool              `json:"sensitive"`
	Source      string            `json:"source"`
}
type state struct {
	Values struct {
		RootModule module `json:"root_module"`
	} `json:"values"`
}
type module struct {
	Address      string     `json:"address"`
	Resources    []resource `json:"resources"`
	ChildModules []module   `json:"child_modules"`
}
type resource struct {
	Address         string         `json:"address"`
	Mode            string         `json:"mode"`
	Type            string         `json:"type"`
	Name            string         `json:"name"`
	ProviderName    string         `json:"provider_name"`
	Values          map[string]any `json:"values"`
	SensitiveValues any            `json:"sensitive_values"`
}

func ParseState(data []byte, environment, stack, cloud, region string) ([]Asset, error) {
	var s state
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("parse Terraform state JSON: %w", err)
	}
	var assets []Asset
	walkModule(s.Values.RootModule, func(m module, r resource) {
		tags := map[string]string{}
		for _, key := range []string{"tags", "labels", "tags_all"} {
			if values, ok := r.Values[key].(map[string]any); ok {
				for k, v := range values {
					if secretKey(k) {
						tags[k] = "[REDACTED]"
					} else if text, ok := v.(string); ok {
						tags[k] = text
					}
				}
			}
		}
		assets = append(assets, Asset{Address: r.Address, Type: r.Type, Name: r.Name, Provider: r.ProviderName, Module: m.Address, Environment: environment, Stack: stack, Cloud: cloud, Region: region, Tags: tags, Sensitive: hasSensitive(r.SensitiveValues), Source: "terraform-state"})
	})
	sort.Slice(assets, func(i, j int) bool { return assets[i].Address < assets[j].Address })
	return assets, nil
}
func walkModule(m module, fn func(module, resource)) {
	for _, r := range m.Resources {
		fn(m, r)
	}
	for _, child := range m.ChildModules {
		walkModule(child, fn)
	}
}
func hasSensitive(v any) bool {
	switch value := v.(type) {
	case bool:
		return value
	case map[string]any:
		for _, child := range value {
			if hasSensitive(child) {
				return true
			}
		}
	case []any:
		for _, child := range value {
			if hasSensitive(child) {
				return true
			}
		}
	}
	return false
}
func secretKey(key string) bool {
	key = strings.ToLower(key)
	return strings.Contains(key, "secret") || strings.Contains(key, "password") || strings.Contains(key, "token") || strings.Contains(key, "key")
}
func WriteCSV(out io.Writer, assets []Asset) error {
	w := csv.NewWriter(out)
	defer w.Flush()
	_ = w.Write([]string{"address", "type", "name", "provider", "module", "environment", "stack", "cloud", "region", "sensitive", "source"})
	for _, a := range assets {
		_ = w.Write([]string{a.Address, a.Type, a.Name, a.Provider, a.Module, a.Environment, a.Stack, a.Cloud, a.Region, fmt.Sprint(a.Sensitive), a.Source})
	}
	return w.Error()
}
func WriteMarkdown(out io.Writer, assets []Asset) error {
	fmt.Fprintln(out, "| Address | Type | Environment | Stack | Cloud | Region | Sensitive | Source |\n| --- | --- | --- | --- | --- | --- | --- | --- |")
	for _, a := range assets {
		fmt.Fprintf(out, "| %s | %s | %s | %s | %s | %s | %t | %s |\n", a.Address, a.Type, a.Environment, a.Stack, a.Cloud, a.Region, a.Sensitive, a.Source)
	}
	return nil
}
