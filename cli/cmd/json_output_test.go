package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/textracta/clusterforge/cli/internal/terraform/planjson"
)

func TestEnvListJSONOutputCanBeUnmarshaled(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "clusterforge.yaml")
	if err := os.WriteFile(configPath, []byte(`project:
  name: demo
environments:
  dev:
    cloud: aws
    region: eu-central-1
    orchestrator: eks
    path: live/dev/aws-eks
`), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

	oldConfigPath := opts.ConfigPath
	oldJSON := envListJSON
	opts.ConfigPath = configPath
	envListJSON = true
	t.Cleanup(func() {
		opts.ConfigPath = oldConfigPath
		envListJSON = oldJSON
	})

	var out bytes.Buffer
	envListCmd.SetOut(&out)
	t.Cleanup(func() { envListCmd.SetOut(nil) })

	if err := envListCmd.RunE(envListCmd, nil); err != nil {
		t.Fatalf("env list json: %v", err)
	}

	var response envListResponse
	if err := json.Unmarshal(out.Bytes(), &response); err != nil {
		t.Fatalf("unmarshal env list json: %v\n%s", err, out.String())
	}
	if len(response.Environments) != 1 || response.Environments[0].Name != "dev" {
		t.Fatalf("response = %#v", response)
	}
}

func TestEnvListHumanOutputStillWorks(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "clusterforge.yaml")
	if err := os.WriteFile(configPath, []byte(`project:
  name: demo
environments:
  dev:
    path: live/dev/aws-eks
`), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

	oldConfigPath := opts.ConfigPath
	oldJSON := envListJSON
	opts.ConfigPath = configPath
	envListJSON = false
	t.Cleanup(func() {
		opts.ConfigPath = oldConfigPath
		envListJSON = oldJSON
	})

	var out bytes.Buffer
	envListCmd.SetOut(&out)
	t.Cleanup(func() { envListCmd.SetOut(nil) })

	if err := envListCmd.RunE(envListCmd, nil); err != nil {
		t.Fatalf("env list human: %v", err)
	}
	if !strings.Contains(out.String(), "dev") || strings.HasPrefix(strings.TrimSpace(out.String()), "{") {
		t.Fatalf("unexpected human output: %q", out.String())
	}
}

func TestAppListJSONOutputCanBeUnmarshaled(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	if err := os.MkdirAll("apps", 0o755); err != nil {
		t.Fatalf("mkdir apps: %v", err)
	}
	if err := os.WriteFile(filepath.Join("apps", "api.yaml"), []byte("name: api\n"), 0o644); err != nil {
		t.Fatalf("write app: %v", err)
	}

	oldJSON := appListJSON
	appListJSON = true
	t.Cleanup(func() { appListJSON = oldJSON })

	var out bytes.Buffer
	appListCmd.SetOut(&out)
	t.Cleanup(func() { appListCmd.SetOut(nil) })

	if err := appListCmd.RunE(appListCmd, nil); err != nil {
		t.Fatalf("app list json: %v", err)
	}

	var response appListResponse
	if err := json.Unmarshal(out.Bytes(), &response); err != nil {
		t.Fatalf("unmarshal app list json: %v\n%s", err, out.String())
	}
	if len(response.Apps) != 1 || response.Apps[0] != "api" {
		t.Fatalf("response = %#v", response)
	}
	if strings.Contains(strings.ToLower(out.String()), "secret") {
		t.Fatalf("app list json should not include secret fields: %s", out.String())
	}
}

func TestDoctorJSONOutputShape(t *testing.T) {
	report := doctorReport{
		Status: doctorWarn,
		Checks: []doctorCheck{{
			Name:    "terraform binary",
			Status:  doctorPass,
			Message: "terraform found",
		}},
	}
	data, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("marshal doctor report: %v", err)
	}
	var decoded doctorReport
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal doctor report: %v", err)
	}
	if decoded.Status != doctorWarn || decoded.Checks[0].Status != doctorPass {
		t.Fatalf("decoded = %#v", decoded)
	}
}

func TestPolicyAndPlanJSONDoNotContainSecrets(t *testing.T) {
	policyResponse := policyCheckResponse{
		Environment: "dev",
		Messages:    []string{"non-production changes should still be reviewed before apply"},
		Summary:     summaryResponse(testSummary()),
	}
	planResponse := planRiskResponse{
		Environment: "dev",
		Stacks: []planStackSummary{{
			Path:    "live/dev/aws-eks",
			Risk:    "LOW",
			Summary: summaryResponse(testSummary()),
		}},
	}
	for name, value := range map[string]any{
		"policy": policyResponse,
		"plan":   planResponse,
	} {
		data, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal %s: %v", name, err)
		}
		var decoded map[string]any
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal %s: %v", name, err)
		}
		if strings.Contains(strings.ToLower(string(data)), "secret") {
			t.Fatalf("%s json should not include secret fields: %s", name, string(data))
		}
	}
}

func testSummary() planjson.Summary {
	return planjson.Summary{
		Creates:   1,
		Updates:   1,
		Addresses: []string{"module.example"},
	}
}
