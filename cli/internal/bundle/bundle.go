package bundle

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	cfapp "github.com/textracta/clusterforge/cli/internal/app"
	"github.com/textracta/clusterforge/cli/internal/config"
	"gopkg.in/yaml.v3"
)

const ChecksumFile = "SHA256SUMS"

type Manifest struct {
	SchemaVersion string    `json:"schema_version"`
	CreatedAt     time.Time `json:"created_at"`
	Project       string    `json:"project"`
	Environment   string    `json:"environment,omitempty"`
	ArtifactMode  string    `json:"artifact_mode"`
}

type Summary struct {
	Manifest Manifest `json:"manifest"`
	Files    []string `json:"files"`
}

func Create(root, output, environment string, cfg *config.Config) error {
	if _, err := os.Stat(output); err == nil {
		return fmt.Errorf("bundle output %s already exists", output)
	}
	if environment != "" {
		if _, ok := cfg.Environments[environment]; !ok {
			return fmt.Errorf("environment %q not found", environment)
		}
	}
	if err := os.MkdirAll(output, 0o700); err != nil {
		return err
	}
	manifest := Manifest{SchemaVersion: "1.0", CreatedAt: time.Now().UTC(), Project: cfg.Project.Name, Environment: environment, ArtifactMode: "manifest-only"}
	data, _ := json.MarshalIndent(manifest, "", "  ")
	if err := os.WriteFile(filepath.Join(output, "bundle-manifest.json"), append(data, '\n'), 0o600); err != nil {
		return err
	}
	if err := copyTree(filepath.Join(root, "modules"), filepath.Join(output, "modules")); err != nil {
		return err
	}
	if err := copyTree(filepath.Join(root, "policies"), filepath.Join(output, "policies")); err != nil {
		return err
	}
	if err := copyTree(filepath.Join(root, "template-packs"), filepath.Join(output, "template-packs")); err != nil {
		return err
	}
	if err := copyEnvironment(root, output, environment, cfg); err != nil {
		return err
	}
	images, err := copyApps(root, output)
	if err != nil {
		return err
	}
	if err := writeLines(filepath.Join(output, "images.txt"), images); err != nil {
		return err
	}
	helm, providers, err := discoverDependencies(output)
	if err != nil {
		return err
	}
	if err := writeLines(filepath.Join(output, "helm-charts.txt"), helm); err != nil {
		return err
	}
	if err := writeLines(filepath.Join(output, "providers.txt"), providers); err != nil {
		return err
	}
	if err := writeRunbookIndex(root, output); err != nil {
		return err
	}
	return WriteChecksums(output)
}

func Inspect(path string) (Summary, error) {
	data, err := os.ReadFile(filepath.Join(path, "bundle-manifest.json"))
	if err != nil {
		return Summary{}, fmt.Errorf("read bundle manifest: %w", err)
	}
	var summary Summary
	if err := json.Unmarshal(data, &summary.Manifest); err != nil {
		return Summary{}, fmt.Errorf("parse bundle manifest: %w", err)
	}
	files, err := bundleFiles(path)
	if err != nil {
		return Summary{}, err
	}
	summary.Files = files
	return summary, nil
}

func Verify(path string) error {
	file, err := os.Open(filepath.Join(path, ChecksumFile))
	if err != nil {
		return fmt.Errorf("read checksums: %w", err)
	}
	defer file.Close()
	seen := map[string]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "  ", 2)
		if len(parts) != 2 || parts[1] == "" {
			return fmt.Errorf("invalid checksum line %q", scanner.Text())
		}
		rel := filepath.Clean(parts[1])
		if filepath.IsAbs(rel) || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			return fmt.Errorf("unsafe checksum path %q", parts[1])
		}
		data, err := os.ReadFile(filepath.Join(path, rel))
		if err != nil {
			return fmt.Errorf("verify %s: %w", rel, err)
		}
		sum := sha256.Sum256(data)
		if hex.EncodeToString(sum[:]) != parts[0] {
			return fmt.Errorf("checksum mismatch: %s", rel)
		}
		seen[filepath.ToSlash(rel)] = true
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	files, err := bundleFiles(path)
	if err != nil {
		return err
	}
	for _, rel := range files {
		if rel != ChecksumFile && !seen[rel] {
			return fmt.Errorf("file missing from checksums: %s", rel)
		}
	}
	return nil
}

func WriteChecksums(root string) error {
	files, err := bundleFiles(root)
	if err != nil {
		return err
	}
	var output strings.Builder
	for _, rel := range files {
		if rel == ChecksumFile {
			continue
		}
		data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
		if err != nil {
			return err
		}
		sum := sha256.Sum256(data)
		fmt.Fprintf(&output, "%s  %s\n", hex.EncodeToString(sum[:]), rel)
	}
	return os.WriteFile(filepath.Join(root, ChecksumFile), []byte(output.String()), 0o600)
}

func copyEnvironment(root, output, selected string, cfg *config.Config) error {
	for name, env := range cfg.Environments {
		if selected != "" && selected != name {
			continue
		}
		if err := copyTree(filepath.Join(root, env.Path), filepath.Join(output, "environments", name)); err != nil {
			return err
		}
	}
	return nil
}

func copyApps(root, output string) ([]string, error) {
	names, err := cfapp.List(root)
	if err != nil {
		return nil, err
	}
	var images []string
	for _, name := range names {
		raw, err := os.ReadFile(cfapp.ManifestPath(root, name))
		if err != nil {
			return nil, err
		}
		var manifest cfapp.Manifest
		if err := yaml.Unmarshal(raw, &manifest); err != nil {
			return nil, fmt.Errorf("parse app manifest %s: %w", name, err)
		}
		if manifest.Image != "" {
			images = append(images, manifest.Image)
		}
		manifest.Env = nil
		data, err := yaml.Marshal(manifest)
		if err != nil {
			return nil, err
		}
		path := filepath.Join(output, "apps", name+".yaml")
		if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
			return nil, err
		}
		if err := os.WriteFile(path, data, 0o600); err != nil {
			return nil, err
		}
	}
	return unique(images), nil
}

func copyTree(source, destination string) error {
	return filepath.WalkDir(source, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			if os.IsNotExist(walkErr) {
				return nil
			}
			return walkErr
		}
		if entry.IsDir() {
			if excludedDir(entry.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return nil
		}
		if excludedFile(entry.Name()) {
			return nil
		}
		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		target := filepath.Join(destination, rel)
		if err := os.MkdirAll(filepath.Dir(target), 0o700); err != nil {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o600)
	})
}

func excludedDir(name string) bool {
	return name == ".terraform" || name == ".git" || name == ".cf" || name == "node_modules"
}
func excludedFile(name string) bool {
	lower := strings.ToLower(name)
	return strings.Contains(lower, ".tfstate") || strings.Contains(lower, "kubeconfig") || lower == ".env" || strings.HasSuffix(lower, ".pem") || strings.HasSuffix(lower, ".key") || strings.Contains(lower, "credentials")
}

var helmPattern = regexp.MustCompile(`(?m)^\s*(?:chart|repository)\s*=\s*"([^"]+)"`)
var providerPattern = regexp.MustCompile(`registry\.terraform\.io/[^/\s"]+/[^/\s"]+`)

func discoverDependencies(root string) ([]string, []string, error) {
	var helm, providers []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("bundle contains unsupported symlink: %s", path)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		text := string(data)
		for _, match := range helmPattern.FindAllStringSubmatch(text, -1) {
			helm = append(helm, match[1])
		}
		providers = append(providers, providerPattern.FindAllString(text, -1)...)
		return nil
	})
	return unique(helm), unique(providers), err
}

func writeRunbookIndex(root, output string) error {
	var lines []string
	for _, dir := range []string{"docs/dr", "docs/incident-response"} {
		entries, err := os.ReadDir(filepath.Join(root, dir))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}
		for _, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ".md" {
				lines = append(lines, filepath.ToSlash(filepath.Join(dir, entry.Name())))
			}
		}
	}
	return writeLines(filepath.Join(output, "runbooks.txt"), lines)
}

func writeLines(path string, lines []string) error {
	return os.WriteFile(path, []byte(strings.Join(unique(lines), "\n")+"\n"), 0o600)
}
func unique(values []string) []string {
	set := map[string]bool{}
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			set[value] = true
		}
	}
	result := make([]string, 0, len(set))
	for value := range set {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}
func bundleFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		files = append(files, filepath.ToSlash(rel))
		return nil
	})
	sort.Strings(files)
	return files, err
}
