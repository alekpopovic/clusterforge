package migration

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Finding struct {
	File  string `json:"file"`
	Line  int    `json:"line"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Report struct {
	Path              string         `json:"path"`
	Providers         []string       `json:"providers"`
	RootModules       []string       `json:"root_modules"`
	ModuleSources     []string       `json:"module_sources"`
	ResourceCounts    map[string]int `json:"resource_counts"`
	Backends          []string       `json:"backends"`
	TFVarsFiles       []string       `json:"tfvars_files"`
	StateFiles        []string       `json:"state_files"`
	PossibleSecrets   []Finding      `json:"possible_secrets"`
	EnvironmentLayout []string       `json:"environment_layout"`
	ModuleDirectories []string       `json:"module_directories"`
	Architecture      []string       `json:"detected_architecture"`
	EquivalentModules []string       `json:"clusterforge_equivalent_modules"`
	Risks             []string       `json:"risks"`
	SuggestedSteps    []string       `json:"suggested_migration_steps"`
	AdoptionNotes     []string       `json:"import_adoption_notes"`
}

var resourceRE = regexp.MustCompile(`resource\s+"([^"]+)"\s+"[^"]+"`)
var moduleRE = regexp.MustCompile(`(?s)module\s+"[^"]+"\s*\{.*?source\s*=\s*"([^"]+)"`)
var providerBlockRE = regexp.MustCompile(`provider\s+"([^"]+)"`)
var backendRE = regexp.MustCompile(`backend\s+"([^"]+)"`)
var secretRE = regexp.MustCompile(`(?i)^\s*([a-z0-9_-]*(?:password|secret|token|api[_-]?key|access[_-]?key)[a-z0-9_-]*)\s*=\s*.+`)

func Analyze(root string) (Report, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return Report{}, err
	}
	info, err := os.Stat(abs)
	if err != nil {
		return Report{}, fmt.Errorf("inspect migration path: %w", err)
	}
	if !info.IsDir() {
		return Report{}, fmt.Errorf("migration path must be a directory")
	}
	report := Report{Path: abs, ResourceCounts: map[string]int{}}
	providers, roots, sources, backends, tfvars, states, moduleDirs, envs := map[string]bool{}, map[string]bool{}, map[string]bool{}, map[string]bool{}, map[string]bool{}, map[string]bool{}, map[string]bool{}, map[string]bool{}
	err = filepath.WalkDir(abs, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if entry.Name() == ".terraform" || entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		rel, _ := filepath.Rel(abs, path)
		rel = filepath.ToSlash(rel)
		lower := strings.ToLower(entry.Name())
		if strings.Contains(lower, ".tfstate") {
			states[rel] = true
			return nil
		}
		if strings.HasSuffix(lower, ".tfvars") || strings.HasSuffix(lower, ".tfvars.json") {
			tfvars[rel] = true
		}
		if filepath.Ext(lower) != ".tf" && !strings.HasSuffix(lower, ".tfvars") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		text := string(data)
		dir := filepath.ToSlash(filepath.Dir(rel))
		if dir == "." {
			dir = "."
		}
		isModuleDir := strings.Contains("/"+dir+"/", "/modules/") || strings.HasPrefix(dir, "modules/")
		if isModuleDir {
			moduleDirs[dir] = true
		} else {
			roots[dir] = true
		}
		for _, match := range resourceRE.FindAllStringSubmatch(text, -1) {
			typeName := match[1]
			report.ResourceCounts[category(typeName)]++
			if prefix := strings.SplitN(typeName, "_", 2)[0]; prefix != "" {
				providers[prefix] = true
			}
		}
		for _, match := range providerBlockRE.FindAllStringSubmatch(text, -1) {
			providers[match[1]] = true
		}
		for _, match := range moduleRE.FindAllStringSubmatch(text, -1) {
			sources[match[1]] = true
		}
		for _, match := range backendRE.FindAllStringSubmatch(text, -1) {
			backends[match[1]] = true
		}
		for index, line := range strings.Split(text, "\n") {
			if match := secretRE.FindStringSubmatch(line); match != nil {
				report.PossibleSecrets = append(report.PossibleSecrets, Finding{File: rel, Line: index + 1, Key: match[1], Value: "[REDACTED]"})
			}
		}
		parts := strings.Split(rel, "/")
		for index, part := range parts {
			if (part == "live" || part == "environments" || part == "envs") && index+1 < len(parts) {
				envs[parts[index+1]] = true
			}
		}
		return nil
	})
	if err != nil {
		return Report{}, fmt.Errorf("analyze Terraform repository: %w", err)
	}
	report.Providers = keys(providers)
	report.RootModules = keys(roots)
	report.ModuleSources = keys(sources)
	report.Backends = keys(backends)
	report.TFVarsFiles = keys(tfvars)
	report.StateFiles = keys(states)
	report.ModuleDirectories = keys(moduleDirs)
	report.EnvironmentLayout = keys(envs)
	buildAssessment(&report)
	return report, nil
}

func category(resource string) string {
	switch {
	case strings.HasPrefix(resource, "aws_vpc") || strings.HasPrefix(resource, "aws_subnet") || strings.HasPrefix(resource, "aws_route") || strings.HasPrefix(resource, "aws_nat_gateway"):
		return "aws_vpc"
	case strings.HasPrefix(resource, "aws_eks"):
		return "eks"
	case strings.HasPrefix(resource, "aws_ecs"):
		return "ecs"
	case strings.HasPrefix(resource, "kubernetes_"):
		return "kubernetes"
	case resource == "helm_release":
		return "helm"
	default:
		return "other"
	}
}

func buildAssessment(report *Report) {
	if report.ResourceCounts["aws_vpc"] > 0 {
		report.Architecture = append(report.Architecture, "AWS VPC networking")
		report.EquivalentModules = append(report.EquivalentModules, "modules/cloud/aws/network")
	}
	if report.ResourceCounts["eks"] > 0 {
		report.Architecture = append(report.Architecture, "Amazon EKS")
		report.EquivalentModules = append(report.EquivalentModules, "modules/orchestrators/kubernetes/eks")
	}
	if report.ResourceCounts["ecs"] > 0 {
		report.Architecture = append(report.Architecture, "Amazon ECS")
		report.EquivalentModules = append(report.EquivalentModules, "modules/orchestrators/ecs/cluster", "modules/workloads/ecs/service")
	}
	if report.ResourceCounts["kubernetes"] > 0 {
		report.Architecture = append(report.Architecture, "Terraform-managed Kubernetes resources")
		report.EquivalentModules = append(report.EquivalentModules, "modules/platform/kubernetes/*", "modules/workloads/kubernetes/*")
	}
	if report.ResourceCounts["helm"] > 0 {
		report.Architecture = append(report.Architecture, "Terraform-managed Helm releases")
	}
	if len(report.StateFiles) > 0 {
		report.Risks = append(report.Risks, "Terraform state files are present and may contain sensitive values; do not copy or commit them.")
	}
	if len(report.PossibleSecrets) > 0 {
		report.Risks = append(report.Risks, "Possible secret assignments were detected; rotate exposed values and move values to an external secret store.")
	}
	if len(report.Backends) == 0 {
		report.Risks = append(report.Risks, "No backend block was detected; confirm state location, locking, encryption, access and recovery.")
	}
	report.SuggestedSteps = []string{"Inventory owners, environments, state backends, provider versions and deployment pipelines.", "Map existing roots/resources to ClusterForge modules and record intentional gaps.", "Generate a parallel non-production composition and compare plans without changing existing state.", "Choose per-resource adoption: retain, refactor with moved blocks, or import only through a reviewed plan.", "Migrate one bounded environment/stack, validate rollback and repeat incrementally."}
	report.AdoptionNotes = []string{"The analyzer does not prove semantic equivalence or inspect remote state/cloud resources.", "Never replace state addresses by copying state files; use reviewed Terraform import/moved procedures and backups.", "Production adoption requires a saved plan, owner approval, drift review and tested recovery."}
	sort.Strings(report.Architecture)
	report.EquivalentModules = unique(report.EquivalentModules)
}

func (report Report) JSON() ([]byte, error) { return json.MarshalIndent(report, "", "  ") }
func keys(set map[string]bool) []string {
	result := make([]string, 0, len(set))
	for key := range set {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}
func unique(values []string) []string {
	set := map[string]bool{}
	for _, value := range values {
		set[value] = true
	}
	return keys(set)
}
