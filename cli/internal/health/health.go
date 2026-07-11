package health

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Status string

const (
	Pass    Status = "pass"
	Warn    Status = "warn"
	Fail    Status = "fail"
	Skipped Status = "skipped"
)

type Check struct {
	Name    string `json:"name"`
	Status  Status `json:"status"`
	Message string `json:"message"`
}
type Report struct {
	Environment string  `json:"environment"`
	Status      Status  `json:"status"`
	SLO         any     `json:"slo,omitempty"`
	Checks      []Check `json:"checks"`
}
type Input struct {
	Environment, Path, Kubeconfig                    string
	SLO                                              any
	CheckNodes, CheckAddons, CheckIngress, CheckApps bool
}

func Evaluate(ctx context.Context, input Input) Report {
	r := Report{Environment: input.Environment, Status: Pass, SLO: input.SLO}
	if info, err := os.Stat(input.Path); err != nil || !info.IsDir() {
		r.Checks = append(r.Checks, Check{"environment.path", Fail, fmt.Sprintf("path %s is unavailable", input.Path)})
		r.Status = Fail
		return r
	}
	r.Checks = append(r.Checks, Check{"environment.path", Pass, input.Path + " exists"})
	if _, err := os.Stat(filepath.Join(input.Path, "terraform.tfstate")); err == nil {
		r.Checks = append(r.Checks, Check{"terraform.state", Pass, "local state file is readable"})
	} else {
		r.Checks = append(r.Checks, Check{"terraform.state", Skipped, "no local state file; remote state was not queried"})
	}
	manifests, _ := filepath.Glob(filepath.Join("apps", "*.yaml"))
	if input.CheckApps || input.CheckIngress {
		if len(manifests) > 0 {
			r.Checks = append(r.Checks, Check{"workloads.manifests", Pass, fmt.Sprintf("%d app manifest(s) present", len(manifests))})
		} else {
			r.Checks = append(r.Checks, Check{"workloads.manifests", Warn, "no app manifests found"})
		}
	}
	if input.Kubeconfig == "" {
		appendSkippedLive(&r, input)
		return finalize(r)
	}
	if input.CheckNodes {
		r.Checks = append(r.Checks, run(ctx, "kubernetes.nodes", "kubectl", []string{"--kubeconfig", input.Kubeconfig, "get", "nodes"}))
	}
	if input.CheckAddons {
		r.Checks = append(r.Checks, run(ctx, "platform.addons", "helm", []string{"--kubeconfig", input.Kubeconfig, "list", "--all-namespaces"}))
	}
	if input.CheckIngress {
		r.Checks = append(r.Checks, run(ctx, "kubernetes.namespaces", "kubectl", []string{"--kubeconfig", input.Kubeconfig, "get", "namespaces"}))
	}
	return finalize(r)
}
func appendSkippedLive(r *Report, input Input) {
	if input.CheckNodes {
		r.Checks = append(r.Checks, Check{"kubernetes.nodes", Skipped, "no kubeconfig configured"})
	}
	if input.CheckAddons {
		r.Checks = append(r.Checks, Check{"platform.addons", Skipped, "no kubeconfig configured"})
	}
	if input.CheckIngress {
		r.Checks = append(r.Checks, Check{"kubernetes.namespaces", Skipped, "no kubeconfig configured"})
	}
}
func run(parent context.Context, name, binary string, args []string) Check {
	if _, err := exec.LookPath(binary); err != nil {
		return Check{name, Skipped, binary + " unavailable"}
	}
	ctx, cancel := context.WithTimeout(parent, 8*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, binary, args...).CombinedOutput()
	if err != nil {
		return Check{name, Warn, fmt.Sprintf("read-only command failed: %v", err)}
	}
	return Check{name, Pass, fmt.Sprintf("read-only check returned %d bytes", len(output))}
}
func finalize(r Report) Report {
	for _, c := range r.Checks {
		if c.Status == Fail {
			r.Status = Fail
			return r
		}
		if c.Status == Warn && r.Status == Pass {
			r.Status = Warn
		}
	}
	return r
}
