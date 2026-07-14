package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/alekpopovic/clusterforge/cli/internal/policyengine"
)

func TestPolicyCheckTreatsProductionLikeEnvironmentAsProduction(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "clusterforge.yaml")
	data := `project:
  name: demo
audit:
  enabled: false
policies:
  require_plan_file_for_apply: false
  block_destroy_in_prod: false
environments:
  prod-eu:
    path: live/prod-eu
`
	if err := os.WriteFile(configPath, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	oldPath, oldPlan, oldJSON, oldFormat, oldPack := opts.ConfigPath, policyPlanFile, policyCheckJSON, policyFormat, policyPack
	opts.ConfigPath, policyPlanFile, policyCheckJSON, policyFormat, policyPack = configPath, "", true, "", ""
	t.Cleanup(func() {
		opts.ConfigPath, policyPlanFile, policyCheckJSON, policyFormat, policyPack = oldPath, oldPlan, oldJSON, oldFormat, oldPack
	})
	var out bytes.Buffer
	policyCheckCmd.SetOut(&out)
	t.Cleanup(func() { policyCheckCmd.SetOut(nil) })
	if err := policyCheckCmd.RunE(policyCheckCmd, []string{"prod-eu"}); err != nil {
		t.Fatalf("policy check: %v", err)
	}

	var result policyengine.Result
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		t.Fatalf("json: %v\n%s", err, out.String())
	}
	for _, finding := range result.Findings {
		if finding.PolicyID == "CF-PROD-001" {
			return
		}
	}
	t.Fatalf("expected production finding, got %#v", result.Findings)
}
