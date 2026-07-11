package secrets

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	cfapp "github.com/textracta/clusterforge/cli/internal/app"
)

type Reference struct {
	Source   string `json:"source"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Key      string `json:"key,omitempty"`
	Variable string `json:"variable,omitempty"`
	App      string `json:"app,omitempty"`
}

var (
	secretNamePattern  = regexp.MustCompile(`(?m)secret_name\s*=\s*"([^"]+)"`)
	secretKeyPattern   = regexp.MustCompile(`(?m)secret_key\s*=\s*"([^"]+)"`)
	valueFromPattern   = regexp.MustCompile(`(?m)value_from\s*=\s*"([^"]+)"`)
	externalKeyPattern = regexp.MustCompile(`(?m)^\s*(?:key|remoteRef):\s*([^\s#]+)`)
)

func FromApps(root, appFilter string) ([]Reference, error) {
	names, err := cfapp.List(root)
	if err != nil {
		return nil, err
	}
	var refs []Reference
	for _, appName := range names {
		if appFilter != "" && appName != appFilter {
			continue
		}
		manifest, err := cfapp.Load(cfapp.ManifestPath(root, appName))
		if err != nil {
			return nil, err
		}
		for variable, ref := range manifest.SecretEnv {
			refs = append(refs, Reference{Source: filepath.ToSlash(cfapp.ManifestPath(root, appName)), Kind: "kubernetes-secret", Name: ref.SecretName, Key: ref.SecretKey, Variable: variable, App: appName})
		}
	}
	return refs, nil
}

func FromFiles(root string) ([]Reference, error) {
	if root == "" {
		return nil, nil
	}
	var refs []Reference
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			if os.IsNotExist(walkErr) {
				return nil
			}
			return walkErr
		}
		if entry.IsDir() {
			if entry.Name() == ".terraform" || entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".tf" && ext != ".yaml" && ext != ".yml" && ext != ".json" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(".", path)
		if err != nil {
			rel = path
		}
		source := filepath.ToSlash(rel)
		text := string(data)
		for _, match := range valueFromPattern.FindAllStringSubmatch(text, -1) {
			refs = append(refs, Reference{Source: source, Kind: "ecs-value-from", Name: match[1]})
		}
		if strings.Contains(text, "secret_env") {
			names := secretNamePattern.FindAllStringSubmatch(text, -1)
			keys := secretKeyPattern.FindAllStringSubmatch(text, -1)
			for index, name := range names {
				ref := Reference{Source: source, Kind: "kubernetes-secret", Name: name[1]}
				if index < len(keys) {
					ref.Key = keys[index][1]
				}
				refs = append(refs, ref)
			}
		}
		if strings.Contains(text, "ExternalSecret") {
			for _, match := range externalKeyPattern.FindAllStringSubmatch(text, -1) {
				refs = append(refs, Reference{Source: source, Kind: "external-secret", Name: strings.Trim(match[1], `"'`)})
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("discover secret references: %w", err)
	}
	return refs, nil
}

func Discover(root, environmentPath, appFilter string) ([]Reference, error) {
	refs, err := FromApps(root, appFilter)
	if err != nil {
		return nil, err
	}
	fileRefs, err := FromFiles(environmentPath)
	if err != nil {
		return nil, err
	}
	refs = append(refs, fileRefs...)
	sort.Slice(refs, func(i, j int) bool {
		left := refs[i].Source + refs[i].Kind + refs[i].Name + refs[i].Key
		right := refs[j].Source + refs[j].Kind + refs[j].Name + refs[j].Key
		return left < right
	})
	return refs, nil
}

func RotationPlan(environment string, refs []Reference) []string {
	steps := []string{
		"Confirm owners, consumers, maintenance window, rollback limits, and audit record for " + environment + ".",
		"Update or rotate each value in its external secret store; never pass values to ClusterForge.",
		"Wait for External Secrets or the platform integration to reconcile references.",
		"Restart Kubernetes pods or redeploy ECS tasks gradually, then verify health and authentication.",
		"Revoke the previous credential only after all consumers use the replacement; record evidence.",
	}
	if len(refs) == 0 {
		steps = append([]string{"No local secret references were discovered; confirm manually before proceeding."}, steps...)
	}
	return steps
}
