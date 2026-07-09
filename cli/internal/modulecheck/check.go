package modulecheck

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Severity string

const (
	Pass Severity = "pass"
	Warn Severity = "warn"
	Fail Severity = "fail"
)

type Finding struct {
	Severity Severity `json:"severity"`
	Check    string   `json:"check"`
	Message  string   `json:"message"`
}

type ModuleResult struct {
	Path     string    `json:"path"`
	Status   Severity  `json:"status"`
	Findings []Finding `json:"findings"`
}

type Report struct {
	Status  Severity       `json:"status"`
	Modules []ModuleResult `json:"modules"`
}

type Options struct {
	Path string
	Root string
}

var requiredFiles = []string{"main.tf", "variables.tf", "outputs.tf", "versions.tf", "README.md"}

func Check(opts Options) (Report, error) {
	root := opts.Root
	if root == "" {
		root = "."
	}
	path := opts.Path
	if path == "" {
		path = "modules"
	}
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return Report{}, fmt.Errorf("resolve root: %w", err)
	}
	absPath := path
	if !filepath.IsAbs(absPath) {
		absPath = filepath.Join(absRoot, path)
	}
	moduleDirs, err := moduleDirectories(absPath)
	if err != nil {
		return Report{}, err
	}
	catalog := readOptional(filepath.Join(absRoot, "MODULE_CATALOG.md"))
	terraformDocsEnabled := fileExists(filepath.Join(absRoot, ".terraform-docs.yml"))

	report := Report{Status: Pass, Modules: make([]ModuleResult, 0, len(moduleDirs))}
	for _, dir := range moduleDirs {
		rel, err := filepath.Rel(absRoot, dir)
		if err != nil {
			rel = dir
		}
		result := checkModule(filepath.ToSlash(rel), dir, catalog, terraformDocsEnabled)
		if result.Status == Fail {
			report.Status = Fail
		} else if result.Status == Warn && report.Status == Pass {
			report.Status = Warn
		}
		report.Modules = append(report.Modules, result)
	}
	return report, nil
}

func moduleDirectories(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("stat %s: %w", path, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", path)
	}
	if fileExists(filepath.Join(path, "versions.tf")) {
		return []string{path}, nil
	}

	var dirs []string
	err = filepath.WalkDir(path, func(current string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			return nil
		}
		if strings.Contains(filepath.ToSlash(current), "/.terraform") {
			return filepath.SkipDir
		}
		if fileExists(filepath.Join(current, "versions.tf")) {
			dirs = append(dirs, current)
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk %s: %w", path, err)
	}
	sort.Strings(dirs)
	return dirs, nil
}

func checkModule(relPath, dir, catalog string, terraformDocsEnabled bool) ModuleResult {
	result := ModuleResult{Path: relPath, Status: Pass}
	add := func(severity Severity, check, message string) {
		result.Findings = append(result.Findings, Finding{Severity: severity, Check: check, Message: message})
		if severity == Fail {
			result.Status = Fail
		} else if severity == Warn && result.Status == Pass {
			result.Status = Warn
		}
	}

	for _, file := range requiredFiles {
		if !fileExists(filepath.Join(dir, file)) {
			add(Fail, "required-file", fmt.Sprintf("missing %s", file))
		}
	}
	if result.Status == Fail {
		return result
	}

	readme := readOptional(filepath.Join(dir, "README.md"))
	if !hasUsage(readme) {
		add(Fail, "readme-usage", "README.md must include a usage section or module example")
	}
	if terraformDocsEnabled && (!strings.Contains(readme, "BEGIN_TF_DOCS") || !strings.Contains(readme, "END_TF_DOCS")) {
		add(Warn, "terraform-docs", "README.md is missing terraform-docs markers")
	}
	if !hasStatus(relPath, readme, catalog) {
		add(Warn, "module-status", "module status is not declared in README.md or MODULE_CATALOG.md")
	}

	variables := readOptional(filepath.Join(dir, "variables.tf"))
	for _, name := range blocksMissingDescription(variables, "variable") {
		add(Warn, "variable-description", fmt.Sprintf("variable %q is missing a description", name))
	}
	outputs := readOptional(filepath.Join(dir, "outputs.tf"))
	for _, name := range blocksMissingDescription(outputs, "output") {
		add(Warn, "output-description", fmt.Sprintf("output %q is missing a description", name))
	}

	versions := readOptional(filepath.Join(dir, "versions.tf"))
	if !strings.Contains(versions, "required_version") {
		add(Fail, "required-version", "versions.tf must declare terraform.required_version")
	}
	checkProviders(dir, versions, add)
	checkSecrets(relPath, dir, add)
	return result
}

func hasUsage(readme string) bool {
	lower := strings.ToLower(readme)
	return strings.Contains(lower, "## usage") || strings.Contains(lower, "# usage") || strings.Contains(readme, `module "`)
}

func hasStatus(relPath, readme, catalog string) bool {
	if strings.Contains(strings.ToLower(readme), "status") {
		return true
	}
	return strings.Contains(catalog, relPath)
}

func blocksMissingDescription(content, kind string) []string {
	re := regexp.MustCompile(`(?s)` + kind + `\s+"([^"]+)"\s*\{(.*?)\}`)
	matches := re.FindAllStringSubmatch(content, -1)
	var missing []string
	for _, match := range matches {
		if !strings.Contains(match[2], "description") {
			missing = append(missing, match[1])
		}
	}
	sort.Strings(missing)
	return missing
}

func checkProviders(dir, versions string, add func(Severity, string, string)) {
	resourceProvider := map[string]string{
		"aws_":        "aws",
		"azurerm_":    "azurerm",
		"google_":     "google",
		"kubernetes_": "kubernetes",
		"helm_":       "helm",
		"nomad_":      "nomad",
		"docker_":     "docker",
	}
	body := readOptional(filepath.Join(dir, "main.tf")) + "\n" + readOptional(filepath.Join(dir, "variables.tf")) + "\n" + readOptional(filepath.Join(dir, "outputs.tf"))
	for prefix, provider := range resourceProvider {
		if strings.Contains(body, `resource "`+prefix) && !strings.Contains(versions, provider+" =") {
			add(Warn, "required-provider", fmt.Sprintf("resources with prefix %s are present but required provider %q was not detected", prefix, provider))
		}
	}
	if regexp.MustCompile(`(?m)^\s*provider\s+"`).FindStringIndex(body) != nil {
		add(Fail, "provider-configuration", "reusable modules must not configure providers")
	}
}

func checkSecrets(relPath, dir string, add func(Severity, string, string)) {
	secretPattern := regexp.MustCompile(`(?i)(aws_access_key_id|aws_secret_access_key|BEGIN (RSA|OPENSSH) PRIVATE KEY|password\s*=\s*"[^"]+"|token\s*=\s*"[^"]+")`)
	for _, name := range []string{"README.md", "main.tf", "variables.tf", "outputs.tf", "terraform.tfvars.example"} {
		path := filepath.Join(dir, name)
		content := readOptional(path)
		if content == "" {
			continue
		}
		if secretPattern.MatchString(content) {
			add(Fail, "obvious-secret", fmt.Sprintf("%s contains an obvious secret-like value", filepath.ToSlash(filepath.Join(relPath, name))))
		}
	}
}

func readOptional(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(data)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
