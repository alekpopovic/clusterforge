package generator

import (
	"strings"
	"testing"
)

func TestRenderTerraformCloudBackend(t *testing.T) {
	got, err := RenderTerraformCloudBackend("example-org", "clusterforge-dev")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`cloud {`, `organization = "example-org"`, `name = "clusterforge-dev"`} {
		if !strings.Contains(got, want) {
			t.Fatalf("missing %q: %s", want, got)
		}
	}
}
func TestRenderTerraformCloudBackendMissingOrganization(t *testing.T) {
	if _, err := RenderTerraformCloudBackend("", "workspace"); err == nil {
		t.Fatal("expected error")
	}
}
