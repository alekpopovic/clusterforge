package upgradeplanner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type KubernetesInput struct {
	Current, Target, NodeVersion, Root string
	Matrix                             []byte
	LiveAccess                         bool
}
type KubernetesPlan struct {
	Readiness string   `json:"readiness"`
	Current   string   `json:"current"`
	Target    string   `json:"target"`
	Blocking  []string `json:"blocking_issues"`
	Warnings  []string `json:"warnings"`
	Steps     []string `json:"recommended_steps"`
	Rollback  []string `json:"rollback_notes"`
}

var versionRE = regexp.MustCompile(`^(\d+)\.(\d+)(?:\.\d+)?$`)

func KubernetesUpgrade(input KubernetesInput) KubernetesPlan {
	p := KubernetesPlan{Readiness: "ready", Current: input.Current, Target: input.Target, Steps: []string{"review provider and platform compatibility", "upgrade control plane one minor", "upgrade node groups", "validate workloads and add-ons"}, Rollback: []string{"Kubernetes control-plane downgrades are generally unsupported", "preserve state and workload backups; roll workloads forward or restore data using tested runbooks"}}
	currentMajor, currentMinor, err1 := parseVersion(input.Current)
	targetMajor, targetMinor, err2 := parseVersion(input.Target)
	if err1 != nil || err2 != nil {
		p.Blocking = append(p.Blocking, "current and target versions must use major.minor format")
	} else if targetMajor != currentMajor || targetMinor-currentMinor != 1 {
		p.Blocking = append(p.Blocking, "upgrade must advance exactly one Kubernetes minor version")
	}
	if input.NodeVersion != "" && input.NodeVersion != input.Current {
		p.Warnings = append(p.Warnings, fmt.Sprintf("node version %s is not aligned with control plane %s", input.NodeVersion, input.Current))
	}
	if !strings.Contains(string(input.Matrix), "`"+input.Target+"`") {
		p.Warnings = append(p.Warnings, "target version is not listed in local VERSION_MATRIX.md tested versions")
	}
	if !input.LiveAccess {
		p.Warnings = append(p.Warnings, "live cluster access unavailable; result uses config and local files only")
	}
	deprecated := scanDeprecated(input.Root)
	for _, item := range deprecated {
		p.Warnings = append(p.Warnings, "deprecated Kubernetes API detected: "+item)
	}
	if len(p.Blocking) > 0 {
		p.Readiness = "blocked"
	} else if len(p.Warnings) > 0 {
		p.Readiness = "ready-with-warnings"
	}
	return p
}
func parseVersion(value string) (int, int, error) {
	m := versionRE.FindStringSubmatch(strings.TrimPrefix(value, "v"))
	if m == nil {
		return 0, 0, fmt.Errorf("invalid version")
	}
	a, _ := strconv.Atoi(m[1])
	b, _ := strconv.Atoi(m[2])
	return a, b, nil
}
func scanDeprecated(root string) []string {
	var found []string
	patterns := []string{"extensions/v1beta1", "apps/v1beta1", "networking.k8s.io/v1beta1", "policy/v1beta1", "batch/v1beta1"}
	_ = filepath.WalkDir(root, func(path string, e fs.DirEntry, err error) error {
		if err != nil || e.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".yaml" && ext != ".yml" && ext != ".tf" {
			return nil
		}
		data, _ := os.ReadFile(path)
		for _, pattern := range patterns {
			if strings.Contains(string(data), pattern) {
				found = append(found, path+" ("+pattern+")")
			}
		}
		return nil
	})
	return found
}
