package scaffold

import (
	"strings"
	"testing"
)

func TestDefaultsAndSummary(t *testing.T) {
	plan := Defaults()
	if err := plan.Validate(); err != nil {
		t.Fatal(err)
	}
	summary := plan.Summary()
	for _, value := range []string{"clusterforge-demo", "local-kind", "production plan required=true"} {
		if !strings.Contains(summary, value) {
			t.Fatalf("missing %q: %s", value, summary)
		}
	}
}
func TestMissingRequiredInputs(t *testing.T) {
	plan := Defaults()
	plan.Project = ""
	if err := plan.Validate(); err == nil {
		t.Fatal("expected project validation")
	}
}
