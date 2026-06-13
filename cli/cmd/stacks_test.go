package cmd

import (
	"strings"
	"testing"

	"github.com/textracta/clusterforge/cli/internal/config"
)

func TestPlanStackPathResolution(t *testing.T) {
	env := stackedTestEnvironment()

	paths, err := resolveStackPaths(env, "network")
	if err != nil {
		t.Fatalf("resolve stack paths: %v", err)
	}
	if len(paths) != 1 || paths[0] != "live/dev/aws-eks/network" {
		t.Fatalf("paths = %#v", paths)
	}
}

func TestResolveStackPathsReturnsDependencyOrder(t *testing.T) {
	env := stackedTestEnvironment()

	paths, err := resolveStackPaths(env, "")
	if err != nil {
		t.Fatalf("resolve all stack paths: %v", err)
	}
	want := []string{
		"live/dev/aws-eks/network",
		"live/dev/aws-eks/cluster",
		"live/dev/aws-eks/platform",
		"live/dev/aws-eks/apps",
	}
	for i := range want {
		if paths[i] != want[i] {
			t.Fatalf("paths[%d] = %q, want %q", i, paths[i], want[i])
		}
	}
}

func TestResolveStackPathsUnknownStack(t *testing.T) {
	env := stackedTestEnvironment()

	_, err := resolveStackPaths(env, "database")
	if err == nil {
		t.Fatal("expected unknown stack to fail")
	}
	if !strings.Contains(err.Error(), "unknown stack") {
		t.Fatalf("error = %v", err)
	}
}

func stackedTestEnvironment() config.Environment {
	return config.Environment{
		Path:   "live/dev/aws-eks",
		Layout: "stacked",
		Stacks: config.Stacks{
			"network":  {Path: "live/dev/aws-eks/network"},
			"cluster":  {Path: "live/dev/aws-eks/cluster"},
			"platform": {Path: "live/dev/aws-eks/platform"},
			"apps":     {Path: "live/dev/aws-eks/apps"},
		},
	}
}
