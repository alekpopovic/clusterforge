package upgradeplanner

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var knownComponents = map[string]bool{"ingress-nginx": true, "cert-manager": true, "external-dns": true, "external-secrets": true, "metrics-server": true, "prometheus-stack": true, "loki": true, "argocd": true, "kyverno": true, "gatekeeper": true, "velero": true, "argo-rollouts": true}
var crdComponents = map[string]bool{"cert-manager": true, "external-secrets": true, "argocd": true, "kyverno": true, "gatekeeper": true, "velero": true, "argo-rollouts": true}

type AddonChange struct {
	Component string `json:"component"`
	Current   string `json:"current,omitempty"`
	Desired   string `json:"desired"`
	Change    bool   `json:"change"`
	CRD       bool   `json:"crd"`
	Warning   string `json:"warning,omitempty"`
}
type PlatformPlan struct {
	Changes  []AddonChange `json:"changes"`
	Warnings []string      `json:"warnings,omitempty"`
}

var chartVersionRE = regexp.MustCompile(`(?m)version\s*=\s*"([^"]+)"`)

func PlatformUpgrade(root string, desired map[string]string) PlatformPlan {
	current := map[string]string{}
	_ = filepath.WalkDir(root, func(path string, e fs.DirEntry, err error) error {
		if err != nil || e.IsDir() || filepath.Ext(path) != ".tf" {
			return nil
		}
		data, _ := os.ReadFile(path)
		for component := range knownComponents {
			if strings.Contains(string(data), component) {
				if m := chartVersionRE.FindSubmatch(data); m != nil {
					current[component] = string(m[1])
				}
			}
		}
		return nil
	})
	names := make([]string, 0, len(desired))
	for name := range desired {
		names = append(names, name)
	}
	sort.Strings(names)
	p := PlatformPlan{}
	for _, name := range names {
		change := AddonChange{Component: name, Current: current[name], Desired: desired[name], Change: current[name] != desired[name], CRD: crdComponents[name]}
		if !knownComponents[name] {
			change.Warning = "unknown platform component"
			p.Warnings = append(p.Warnings, name+": unknown platform component")
		}
		if change.CRD && change.Change {
			change.Warning = "CRD upgrade requires manual compatibility and rollback review"
		}
		p.Changes = append(p.Changes, change)
	}
	return p
}
