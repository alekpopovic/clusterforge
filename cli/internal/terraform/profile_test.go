package terraform

import (
	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"reflect"
	"testing"
)

func boolp(v bool) *bool { return &v }
func TestProfilePlanArgs(t *testing.T) {
	p := config.ExecutionProfile{Parallelism: 5, Refresh: boolp(true), LockTimeout: "10m", Input: boolp(false)}
	want := []string{"-parallelism=5", "-refresh=true", "-lock-timeout=10m", "-input=false"}
	if got := ProfilePlanArgs(p); !reflect.DeepEqual(got, want) {
		t.Fatalf("%#v", got)
	}
}
func TestUnknownProfile(t *testing.T) {
	if _, err := ResolveProfile(nil, "missing"); err == nil {
		t.Fatal("expected error")
	}
}
