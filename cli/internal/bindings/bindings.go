// Package bindings resolves app dependency declarations into Terraform output
// references. It deliberately handles references, not secret values.
package bindings

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const RegistryFile = "dependencies.yaml"

var identifierPattern = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

type Request struct {
	Type      string            `yaml:"type"`
	Reference string            `yaml:"reference"`
	Env       map[string]string `yaml:"env,omitempty"`
	Secrets   map[string]string `yaml:"secrets,omitempty"`
}

type Registry struct {
	Dependencies map[string]Entry `yaml:"dependencies"`
}

type Entry struct {
	Type          string            `yaml:"type"`
	Module        string            `yaml:"module,omitempty"`
	Key           string            `yaml:"key,omitempty"`
	RemoteState   string            `yaml:"remote_state,omitempty"`
	RemoteOutputs map[string]string `yaml:"remote_outputs,omitempty"`
}

type Result struct {
	Environment map[string]string
	Comments    []string
}

func Load(environmentPath string) (Registry, error) {
	path := filepath.Join(environmentPath, RegistryFile)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Registry{Dependencies: map[string]Entry{}}, nil
		}
		return Registry{}, fmt.Errorf("read dependency registry %s: %w", path, err)
	}
	var registry Registry
	if err := yaml.Unmarshal(data, &registry); err != nil {
		return Registry{}, fmt.Errorf("parse dependency registry %s: %w", path, err)
	}
	if registry.Dependencies == nil {
		registry.Dependencies = map[string]Entry{}
	}
	return registry, nil
}

func Resolve(registry Registry, requests map[string]Request) (Result, error) {
	result := Result{Environment: map[string]string{}}
	aliases := make([]string, 0, len(requests))
	for alias := range requests {
		aliases = append(aliases, alias)
	}
	sort.Strings(aliases)
	for _, alias := range aliases {
		request := requests[alias]
		entry, ok := registry.Dependencies[request.Reference]
		if !ok {
			return Result{}, fmt.Errorf("dependency %q references unknown registry entry %q", alias, request.Reference)
		}
		if request.Type == "" || request.Type != entry.Type {
			return Result{}, fmt.Errorf("dependency %q type %q does not match registry type %q", alias, request.Type, entry.Type)
		}
		if entry.RemoteState != "" {
			result.Comments = append(result.Comments, fmt.Sprintf("Cross-stack dependency %s uses data.terraform_remote_state.%s; configure that data source with a reviewed backend.", request.Reference, entry.RemoteState))
		}
		for envName, output := range request.Env {
			if looksSecret(output) {
				return Result{}, fmt.Errorf("dependency %q output %q looks secret; declare it under secrets so plaintext is never rendered", alias, output)
			}
			expression, err := outputExpression(request.Reference, output, entry)
			if err != nil {
				return Result{}, fmt.Errorf("dependency %q: %w", alias, err)
			}
			result.Environment[envName] = expression
		}
		for envName, output := range request.Secrets {
			result.Comments = append(result.Comments, fmt.Sprintf("Secret binding %s from %s.%s requires an external secret reference; no value was rendered.", envName, request.Reference, output))
		}
	}
	sort.Strings(result.Comments)
	return result, nil
}

func outputExpression(reference, output string, entry Entry) (string, error) {
	terraformOutput, indexed, ok := supportedOutput(entry.Type, output)
	if !ok {
		return "", fmt.Errorf("unsupported output %q for dependency type %q", output, entry.Type)
	}
	if entry.RemoteState != "" {
		if !identifierPattern.MatchString(entry.RemoteState) {
			return "", fmt.Errorf("remote_state %q must be a Terraform identifier", entry.RemoteState)
		}
		remoteOutput := entry.RemoteOutputs[output]
		if remoteOutput == "" {
			return "", fmt.Errorf("remote output mapping for %q is required", output)
		}
		if !identifierPattern.MatchString(remoteOutput) {
			return "", fmt.Errorf("remote output %q must be a Terraform identifier", remoteOutput)
		}
		return fmt.Sprintf("data.terraform_remote_state.%s.outputs.%s", entry.RemoteState, remoteOutput), nil
	}
	module := entry.Module
	if module == "" {
		module = reference
	}
	if !identifierPattern.MatchString(module) {
		return "", fmt.Errorf("module %q must be a Terraform identifier", module)
	}
	expression := fmt.Sprintf("module.%s.%s", module, terraformOutput)
	if indexed {
		key := entry.Key
		if key == "" {
			key = reference
		}
		expression += fmt.Sprintf("[%q]", key)
	}
	return expression, nil
}

func supportedOutput(dependencyType, output string) (string, bool, bool) {
	supported := map[string]map[string]struct {
		name    string
		indexed bool
	}{
		"rds-postgres": {
			"endpoint": {name: "endpoint"}, "port": {name: "port"}, "db_name": {name: "db_name"},
		},
		"elasticache-redis": {
			"primary_endpoint_address": {name: "primary_endpoint_address"}, "reader_endpoint_address": {name: "reader_endpoint_address"}, "port": {name: "port"},
		},
		"sqs": {
			"queue_url": {name: "queue_urls", indexed: true}, "queue_arn": {name: "queue_arns", indexed: true},
		},
	}
	definition, ok := supported[dependencyType][output]
	return definition.name, definition.indexed, ok
}

func looksSecret(value string) bool {
	lower := strings.ToLower(value)
	return strings.Contains(lower, "password") || strings.Contains(lower, "secret_value") || strings.Contains(lower, "private_key") || strings.Contains(lower, "token")
}
