package terraform

import (
	"fmt"
	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"strconv"
)

func ProfilePlanArgs(p config.ExecutionProfile) []string  { return profileArgs(p, true) }
func ProfileApplyArgs(p config.ExecutionProfile) []string { return profileArgs(p, false) }
func profileArgs(p config.ExecutionProfile, plan bool) []string {
	var args []string
	if p.Parallelism > 0 {
		args = append(args, "-parallelism="+strconv.Itoa(p.Parallelism))
	}
	if plan && p.Refresh != nil {
		args = append(args, "-refresh="+strconv.FormatBool(*p.Refresh))
	}
	if p.LockTimeout != "" {
		args = append(args, "-lock-timeout="+p.LockTimeout)
	}
	if p.Input != nil {
		args = append(args, "-input="+strconv.FormatBool(*p.Input))
	}
	return args
}
func ResolveProfile(profiles map[string]config.ExecutionProfile, name string) (config.ExecutionProfile, error) {
	p, ok := profiles[name]
	if !ok {
		return config.ExecutionProfile{}, fmt.Errorf("execution profile %q not found", name)
	}
	return p, nil
}
