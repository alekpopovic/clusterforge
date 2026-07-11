package policyengine

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type Policy struct {
	ID            string   `json:"id" yaml:"id"`
	Title         string   `json:"title" yaml:"title"`
	Description   string   `json:"description" yaml:"description"`
	Severity      string   `json:"severity" yaml:"severity"`
	Category      string   `json:"category" yaml:"category"`
	Scope         string   `json:"scope" yaml:"scope"`
	DefaultAction string   `json:"default_action" yaml:"default_action"`
	Remediation   string   `json:"remediation" yaml:"remediation"`
	References    []string `json:"references,omitempty" yaml:"references,omitempty"`
	Enabled       bool     `json:"enabled" yaml:"enabled"`
}

type Finding struct {
	PolicyID string `json:"policy_id"`
	Title    string `json:"title"`
	Severity string `json:"severity"`
	Action   string `json:"action"`
	Scope    string `json:"scope"`
	Message  string `json:"message"`
	Location string `json:"location,omitempty"`
}

type Input struct {
	Environment      string
	Production       bool
	BackendType      string
	TrackedFiles     []string
	AppPath          string
	AppYAML          []byte
	Image            string
	IngressEnabled   bool
	TerraformFiles   map[string][]byte
	RequirePlanFile  bool
	BlockProdDestroy bool
	ExternalFindings []Finding
}

type Options struct {
	Pack      string
	Overrides map[string]string
}

type Result struct {
	Pack     string    `json:"pack"`
	Findings []Finding `json:"findings"`
	Blocked  bool      `json:"blocked"`
}

func BuiltIns() []Policy {
	return []Policy{
		{ID: "CF-REPO-001", Title: "Terraform state must not be committed", Description: "Tracked state can expose sensitive infrastructure data.", Severity: "critical", Category: "secrets", Scope: "repository", DefaultAction: "block", Remediation: "Remove state from Git and rotate exposed values.", Enabled: true},
		{ID: "CF-REPO-002", Title: ".env files must not be committed", Description: "Environment files commonly contain credentials.", Severity: "high", Category: "secrets", Scope: "repository", DefaultAction: "block", Remediation: "Remove the file and use a secret store.", Enabled: true},
		{ID: "CF-APP-001", Title: "App manifests must not contain plaintext secrets", Description: "Secret-looking literal values belong in external secret stores.", Severity: "high", Category: "secrets", Scope: "app", DefaultAction: "block", Remediation: "Use secret_env references.", Enabled: true},
		{ID: "CF-APP-002", Title: "Production images must not use latest", Description: "Mutable image tags make deployments non-repeatable.", Severity: "high", Category: "supply-chain", Scope: "image", DefaultAction: "block", Remediation: "Use a version tag or digest.", Enabled: true},
		{ID: "CF-PROD-001", Title: "Production apply requires a plan", Description: "Production applies must consume a reviewed plan file.", Severity: "critical", Category: "safety", Scope: "environment", DefaultAction: "block", Remediation: "Enable require_plan_file_for_apply.", Enabled: true},
		{ID: "CF-PROD-002", Title: "Production destroy is blocked", Description: "Production destroy must be blocked by default.", Severity: "critical", Category: "safety", Scope: "environment", DefaultAction: "block", Remediation: "Enable block_destroy_in_prod.", Enabled: true},
		{ID: "CF-PROD-003", Title: "Production backend must be remote", Description: "Local state lacks shared locking and managed durability.", Severity: "high", Category: "state", Scope: "state", DefaultAction: "block", Remediation: "Configure an encrypted remote backend.", Enabled: true},
		{ID: "CF-MOD-001", Title: "Production module source uses a mutable branch", Description: "A main/master ref can change without review.", Severity: "high", Category: "supply-chain", Scope: "module", DefaultAction: "warn", Remediation: "Pin a release tag or commit SHA.", Enabled: true},
		{ID: "CF-NET-001", Title: "Public ingress requires approval", Description: "Public exposure must be explicitly approved.", Severity: "high", Category: "network", Scope: "app", DefaultAction: "warn", Remediation: "Add clusterforge.io/public-ingress-approved: true after review.", Enabled: true},
		{ID: "CF-K8S-001", Title: "Production LoadBalancer requires approval", Description: "LoadBalancer services may create public endpoints.", Severity: "high", Category: "network", Scope: "plan", DefaultAction: "warn", Remediation: "Add an approval annotation or use an internal service.", Enabled: true},
		{ID: "CF-IAM-001", Title: "Wildcard IAM permissions require review", Description: "Wildcard actions or resources can create broad privilege.", Severity: "high", Category: "iam", Scope: "plan", DefaultAction: "warn", Remediation: "Scope actions and resources to required values.", Enabled: true},
	}
}

var secretPattern = regexp.MustCompile(`(?i)(password|secret|token|api[_-]?key|access[_-]?key)\s*:\s*["']?[A-Za-z0-9+/=_-]{8,}`)
var mutableModuleRef = regexp.MustCompile(`(?i)(ref=|\?ref=)(main|master)([^A-Za-z0-9_-]|$)`)

func Evaluate(input Input, options Options) Result {
	policies := map[string]Policy{}
	for _, policy := range BuiltIns() {
		policies[policy.ID] = policy
	}
	var findings []Finding
	add := func(id, message, location string) {
		policy := policies[id]
		action := policy.DefaultAction
		if options.Pack != "production" && action == "block" && id != "CF-REPO-001" && id != "CF-REPO-002" && id != "CF-APP-001" {
			action = "warn"
		}
		if override := options.Overrides[id]; override != "" {
			action = override
		}
		findings = append(findings, Finding{PolicyID: id, Title: policy.Title, Severity: policy.Severity, Action: action, Scope: policy.Scope, Message: message, Location: location})
	}
	for _, file := range input.TrackedFiles {
		lower := strings.ToLower(filepathSlash(file))
		if strings.HasSuffix(lower, ".tfstate") || strings.Contains(lower, ".tfstate.") {
			add("CF-REPO-001", "tracked Terraform state file", file)
		}
		base := lower[strings.LastIndex(lower, "/")+1:]
		if base == ".env" || strings.HasPrefix(base, ".env.") {
			add("CF-REPO-002", "tracked environment file", file)
		}
	}
	if len(input.AppYAML) > 0 && secretPattern.Match(input.AppYAML) {
		add("CF-APP-001", "secret-looking literal found in app manifest", input.AppPath)
	}
	if input.Production && imageLatest(input.Image) {
		add("CF-APP-002", "production image uses latest tag", input.AppPath)
	}
	if input.Production && !input.RequirePlanFile {
		add("CF-PROD-001", "production plan-file policy is disabled", "clusterforge.yaml")
	}
	if input.Production && !input.BlockProdDestroy {
		add("CF-PROD-002", "production destroy protection is disabled", "clusterforge.yaml")
	}
	if input.Production && (input.BackendType == "" || input.BackendType == "local") {
		add("CF-PROD-003", "production environment uses local state", "clusterforge.yaml")
	}
	approved := strings.Contains(string(input.AppYAML), "clusterforge.io/public-ingress-approved: true")
	if input.Production && input.IngressEnabled && !approved {
		add("CF-NET-001", "production app enables ingress without approval annotation", input.AppPath)
	}
	for path, content := range input.TerraformFiles {
		text := string(content)
		if input.Production && mutableModuleRef.MatchString(text) {
			add("CF-MOD-001", "module source uses main/master ref", path)
		}
		if input.Production && strings.Contains(text, `type = "LoadBalancer"`) && !strings.Contains(text, "clusterforge.io/load-balancer-approved") {
			add("CF-K8S-001", "LoadBalancer service lacks approval", path)
		}
		if strings.Contains(text, `"Action":"*"`) || strings.Contains(text, `"Resource":"*"`) || strings.Contains(text, `actions = ["*"]`) || strings.Contains(text, `resources = ["*"]`) {
			add("CF-IAM-001", "wildcard IAM action or resource found", path)
		}
	}
	findings = append(findings, input.ExternalFindings...)
	sort.Slice(findings, func(i, j int) bool {
		if findings[i].PolicyID == findings[j].PolicyID {
			return findings[i].Location < findings[j].Location
		}
		return findings[i].PolicyID < findings[j].PolicyID
	})
	result := Result{Pack: options.Pack, Findings: findings}
	for _, finding := range findings {
		if finding.Action == "block" {
			result.Blocked = true
		}
	}
	return result
}

func imageLatest(image string) bool {
	if strings.TrimSpace(image) == "" {
		return false
	}
	withoutDigest := strings.Split(image, "@")[0]
	lastSlash := strings.LastIndex(withoutDigest, "/")
	lastColon := strings.LastIndex(withoutDigest, ":")
	return lastColon <= lastSlash || strings.EqualFold(withoutDigest[lastColon+1:], "latest")
}

func filepathSlash(path string) string { return strings.ReplaceAll(path, "\\", "/") }

func SARIF(result Result) ([]byte, error) {
	type sarifResult struct {
		RuleID    string            `json:"ruleId"`
		Level     string            `json:"level"`
		Message   map[string]string `json:"message"`
		Locations []map[string]any  `json:"locations,omitempty"`
	}
	results := make([]sarifResult, 0, len(result.Findings))
	for _, finding := range result.Findings {
		level := "warning"
		if finding.Severity == "critical" || finding.Action == "block" {
			level = "error"
		}
		item := sarifResult{RuleID: finding.PolicyID, Level: level, Message: map[string]string{"text": finding.Message}}
		if finding.Location != "" {
			item.Locations = []map[string]any{{"physicalLocation": map[string]any{"artifactLocation": map[string]string{"uri": finding.Location}}}}
		}
		results = append(results, item)
	}
	document := map[string]any{"version": "2.1.0", "$schema": "https://json.schemastore.org/sarif-2.1.0.json", "runs": []any{map[string]any{"tool": map[string]any{"driver": map[string]string{"name": "ClusterForge policy engine"}}, "results": results}}}
	return json.MarshalIndent(document, "", "  ")
}

func ValidateAction(action string) error {
	if action != "advisory" && action != "warn" && action != "block" {
		return fmt.Errorf("invalid policy action %q", action)
	}
	return nil
}

type policyOverrideFile struct {
	Policies []struct {
		ID      string `yaml:"id"`
		Action  string `yaml:"action"`
		Enabled *bool  `yaml:"enabled,omitempty"`
	} `yaml:"policies"`
}

func LoadOverrides(directories ...string) (map[string]string, error) {
	result := map[string]string{}
	for _, directory := range directories {
		entries, err := os.ReadDir(directory)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		for _, entry := range entries {
			if entry.IsDir() || (filepath.Ext(entry.Name()) != ".yaml" && filepath.Ext(entry.Name()) != ".yml") {
				continue
			}
			data, err := os.ReadFile(filepath.Join(directory, entry.Name()))
			if err != nil {
				return nil, err
			}
			var file policyOverrideFile
			if err := yaml.Unmarshal(data, &file); err != nil {
				return nil, fmt.Errorf("parse policy overrides %s: %w", entry.Name(), err)
			}
			for _, override := range file.Policies {
				if override.Enabled != nil && !*override.Enabled {
					result[override.ID] = "advisory"
					continue
				}
				if err := ValidateAction(override.Action); err != nil {
					return nil, fmt.Errorf("policy %s: %w", override.ID, err)
				}
				result[override.ID] = override.Action
			}
		}
	}
	return result, nil
}
