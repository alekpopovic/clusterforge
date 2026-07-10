package app

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/textracta/clusterforge/cli/internal/bindings"
	"github.com/textracta/clusterforge/cli/internal/config"
)

func TestAddAppManifest(t *testing.T) {
	dir := t.TempDir()
	path, err := Add(dir, "api", AddOptions{
		Image:    "ghcr.io/company/api:1.0.0",
		Port:     8080,
		Replicas: 2,
		Host:     "api.dev.example.com",
		Type:     "web",
	})
	if err != nil {
		t.Fatalf("add app: %v", err)
	}
	manifest, err := Load(path)
	if err != nil {
		t.Fatalf("load app: %v", err)
	}
	if manifest.Name != "api" || manifest.Image != "ghcr.io/company/api:1.0.0" {
		t.Fatalf("manifest = %#v", manifest)
	}
	if !manifest.Ingress.Enabled || manifest.Ingress.Host != "api.dev.example.com" {
		t.Fatalf("ingress = %#v", manifest.Ingress)
	}
}

func TestNewManifestWithAutoscaling(t *testing.T) {
	manifest := NewManifest("api", AddOptions{
		Image:       "nginx:1.27",
		Replicas:    2,
		Autoscaling: true,
	})
	if !manifest.Autoscaling.Enabled {
		t.Fatal("autoscaling should be enabled")
	}
	if manifest.Autoscaling.MinReplicas != 2 || manifest.Autoscaling.MaxReplicas != 3 {
		t.Fatalf("autoscaling = %#v", manifest.Autoscaling)
	}
}

func TestValidManifestPassesValidation(t *testing.T) {
	if err := sampleManifest().Validate(); err != nil {
		t.Fatalf("valid manifest failed validation: %v", err)
	}
}

func TestMissingImageFailsValidation(t *testing.T) {
	manifest := sampleManifest()
	manifest.Image = ""
	assertValidationError(t, manifest.Validate(), "image is required")
}

func TestImagePolicyWarnsForLatestTag(t *testing.T) {
	manifest := sampleManifest()
	manifest.Image = "nginx:latest"

	warnings := ImagePolicyWarnings(manifest, "dev")
	if len(warnings) != 1 || !strings.Contains(warnings[0], "latest tag") {
		t.Fatalf("warnings = %#v", warnings)
	}
}

func TestImagePolicyAcceptsDigest(t *testing.T) {
	manifest := sampleManifest()
	manifest.Image = "nginx@sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	if warnings := ImagePolicyWarnings(manifest, "prod"); len(warnings) != 0 {
		t.Fatalf("warnings = %#v", warnings)
	}
}

func TestImagePolicyAcceptsVersionTag(t *testing.T) {
	manifest := sampleManifest()
	manifest.Image = "nginx:1.27"

	if warnings := ImagePolicyWarnings(manifest, "prod"); len(warnings) != 0 {
		t.Fatalf("warnings = %#v", warnings)
	}
}

func TestImagePolicyWarnsForUnpinnedProdImage(t *testing.T) {
	manifest := sampleManifest()
	manifest.Image = "nginx"

	warnings := ImagePolicyWarnings(manifest, "prod")
	if len(warnings) != 1 || !strings.Contains(warnings[0], "not pinned") {
		t.Fatalf("warnings = %#v", warnings)
	}
}

func TestBadPortFailsValidation(t *testing.T) {
	manifest := sampleManifest()
	manifest.Ports[0].ContainerPort = 70000
	assertValidationError(t, manifest.Validate(), "ports[0].container_port must be between 1 and 65535")
}

func TestIngressEnabledWithoutHostFailsValidation(t *testing.T) {
	manifest := sampleManifest()
	manifest.Ingress.Host = ""
	assertValidationError(t, manifest.Validate(), "ingress.host is required when ingress.enabled=true")
}

func TestSecretValueFailsValidation(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "api.yaml")
	data := []byte(`
name: api
type: web
image: nginx:1.27
replicas: 1
secret_env:
  DATABASE_URL:
    value: postgres://example
`)
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("write manifest: %v", err)
	}
	assertValidationError(t, ValidateFile(path), "secret_env.DATABASE_URL.value is not allowed")
}

func TestListAppManifests(t *testing.T) {
	dir := t.TempDir()
	if _, err := Add(dir, "worker", AddOptions{Image: "busybox:1.36"}); err != nil {
		t.Fatalf("add worker: %v", err)
	}
	if _, err := Add(dir, "api", AddOptions{Image: "nginx:1.27"}); err != nil {
		t.Fatalf("add api: %v", err)
	}
	apps, err := List(dir)
	if err != nil {
		t.Fatalf("list apps: %v", err)
	}
	if strings.Join(apps, ",") != "api,worker" {
		t.Fatalf("apps = %#v", apps)
	}
}

func TestRenderKubernetesModuleCall(t *testing.T) {
	dir := t.TempDir()
	manifest := sampleManifest()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
		Path:         filepath.Join(dir, "live", "dev", "aws-eks"),
	}
	outPath, err := Render(dir, "dev", env, manifest)
	if err != nil {
		t.Fatalf("render kubernetes app: %v", err)
	}
	rendered := readFile(t, outPath)
	for _, want := range []string{
		`source = "../../../../modules/workloads/kubernetes/app"`,
		`name      = "api"`,
		`namespace = "dev"`,
		`secret_env = {`,
		`secret_name = "api-secrets"`,
		`autoscaling = {`,
	} {
		if !strings.Contains(rendered, want) {
			t.Fatalf("rendered Kubernetes app missing %q:\n%s", want, rendered)
		}
	}
}

func TestRenderECSModuleCall(t *testing.T) {
	dir := t.TempDir()
	manifest := sampleManifest()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "ecs",
		Path:         filepath.Join(dir, "live", "dev", "aws-ecs"),
	}
	outPath, err := Render(dir, "dev", env, manifest)
	if err != nil {
		t.Fatalf("render ecs app: %v", err)
	}
	rendered := readFile(t, outPath)
	for _, want := range []string{
		`source = "../../../../modules/workloads/ecs/service"`,
		`cluster_arn        = module.ecs_cluster.cluster_arn`,
		`security_group_ids = var.app_security_group_ids`,
		`container_port = 8080`,
		`Replace value_from with existing SSM Parameter Store or Secrets Manager ARNs`,
	} {
		if !strings.Contains(rendered, want) {
			t.Fatalf("rendered ECS app missing %q:\n%s", want, rendered)
		}
	}
}

func TestRenderEKSCloudIdentity(t *testing.T) {
	dir := t.TempDir()
	manifest := sampleManifest()
	manifest.CloudIdentity = CloudIdentity{
		Enabled:    true,
		Provider:   "aws",
		PolicyARNs: []string{"arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"},
	}
	env := config.Environment{Cloud: "aws", Orchestrator: "eks", Path: filepath.Join(dir, "live", "dev", "aws-eks")}
	outPath, err := Render(dir, "dev", env, manifest)
	if err != nil {
		t.Fatalf("render EKS cloud identity: %v", err)
	}
	rendered := readFile(t, outPath)
	for _, want := range []string{
		`module "api_irsa"`,
		`source = "../../../../modules/cloud/aws/irsa-role"`,
		`"eks.amazonaws.com/role-arn" = module.api_irsa.role_arn`,
		`arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess`,
	} {
		if !strings.Contains(rendered, want) {
			t.Fatalf("rendered EKS identity missing %q:\n%s", want, rendered)
		}
	}
}

func TestRenderECSCloudIdentity(t *testing.T) {
	dir := t.TempDir()
	manifest := sampleManifest()
	manifest.CloudIdentity = CloudIdentity{
		Enabled:    true,
		Provider:   "aws",
		PolicyARNs: []string{"arn:aws:iam::aws:policy/AmazonSQSReadOnlyAccess"},
	}
	env := config.Environment{Cloud: "aws", Orchestrator: "ecs", Path: filepath.Join(dir, "live", "dev", "aws-ecs")}
	outPath, err := Render(dir, "dev", env, manifest)
	if err != nil {
		t.Fatalf("render ECS cloud identity: %v", err)
	}
	if rendered := readFile(t, outPath); !strings.Contains(rendered, "task_role_policy_arns") || !strings.Contains(rendered, "AmazonSQSReadOnlyAccess") {
		t.Fatalf("rendered ECS task role is incomplete:\n%s", rendered)
	}
}

func TestCloudIdentityUnsupportedTargetFails(t *testing.T) {
	manifest := sampleManifest()
	manifest.CloudIdentity = CloudIdentity{Enabled: true, Provider: "aws"}
	dir := t.TempDir()
	env := config.Environment{Cloud: "gcp", Orchestrator: "gke", Path: filepath.Join(dir, "live", "dev", "gke")}
	_, err := Render(dir, "dev", env, manifest)
	assertValidationError(t, err, "not supported for orchestrator")
}

func TestCloudIdentityRejectsAdministratorPolicy(t *testing.T) {
	manifest := sampleManifest()
	manifest.CloudIdentity = CloudIdentity{Enabled: true, Provider: "aws", PolicyARNs: []string{"arn:aws:iam::aws:policy/AdministratorAccess"}}
	assertValidationError(t, manifest.Validate(), "must not grant administrator access")
}

func TestRenderServiceBindings(t *testing.T) {
	dir := t.TempDir()
	envPath := filepath.Join(dir, "live", "dev", "aws-eks")
	if err := os.MkdirAll(envPath, 0o755); err != nil {
		t.Fatal(err)
	}
	registry := []byte(`dependencies:
  main:
    type: rds-postgres
    module: database
  jobs:
    type: sqs
    module: queues
    key: jobs
`)
	if err := os.WriteFile(filepath.Join(envPath, "dependencies.yaml"), registry, 0o644); err != nil {
		t.Fatal(err)
	}
	manifest := sampleManifest()
	manifest.Dependencies = map[string]bindings.Request{
		"database": {Type: "rds-postgres", Reference: "main", Env: map[string]string{"DATABASE_HOST": "endpoint"}},
		"queue":    {Type: "sqs", Reference: "jobs", Env: map[string]string{"QUEUE_URL": "queue_url"}},
	}
	outPath, err := Render(dir, "dev", config.Environment{Cloud: "aws", Orchestrator: "eks", Path: envPath}, manifest)
	if err != nil {
		t.Fatalf("render bindings: %v", err)
	}
	rendered := readFile(t, outPath)
	for _, want := range []string{`"DATABASE_HOST" = module.database.endpoint`, `"QUEUE_URL" = module.queues.queue_urls["jobs"]`} {
		if !strings.Contains(rendered, want) {
			t.Fatalf("rendered binding missing %q:\n%s", want, rendered)
		}
	}
}

func TestRenderUnknownServiceBindingFails(t *testing.T) {
	dir := t.TempDir()
	manifest := sampleManifest()
	manifest.Dependencies = map[string]bindings.Request{
		"database": {Type: "rds-postgres", Reference: "missing", Env: map[string]string{"DATABASE_HOST": "endpoint"}},
	}
	_, err := Render(dir, "dev", config.Environment{Cloud: "aws", Orchestrator: "eks", Path: filepath.Join(dir, "live", "dev")}, manifest)
	assertValidationError(t, err, "unknown registry entry")
}

func TestRenderSecretBindingDoesNotRenderValue(t *testing.T) {
	dir := t.TempDir()
	envPath := filepath.Join(dir, "live", "dev")
	if err := os.MkdirAll(envPath, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(envPath, "dependencies.yaml"), []byte("dependencies:\n  main:\n    type: rds-postgres\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	manifest := sampleManifest()
	manifest.Dependencies = map[string]bindings.Request{
		"database": {Type: "rds-postgres", Reference: "main", Secrets: map[string]string{"DATABASE_PASSWORD": "password"}},
	}
	outPath, err := Render(dir, "dev", config.Environment{Cloud: "aws", Orchestrator: "eks", Path: envPath}, manifest)
	if err != nil {
		t.Fatalf("render secret binding: %v", err)
	}
	rendered := readFile(t, outPath)
	if strings.Contains(rendered, `DATABASE_PASSWORD =`) || !strings.Contains(rendered, "requires an external secret reference") {
		t.Fatalf("secret binding was not safely rendered:\n%s", rendered)
	}
}

func TestRenderUnsupportedOrchestrator(t *testing.T) {
	dir := t.TempDir()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "nomad",
		Path:         filepath.Join(dir, "live", "dev", "nomad"),
	}
	_, err := Render(dir, "dev", env, sampleManifest())
	if err == nil {
		t.Fatal("expected unsupported orchestrator error")
	}
	if !strings.Contains(err.Error(), "unsupported orchestrator") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func sampleManifest() Manifest {
	return Manifest{
		Name:     "api",
		Type:     "web",
		Image:    "ghcr.io/company/api:1.0.0",
		Replicas: 2,
		Ports: []Port{{
			Name:          "http",
			ContainerPort: 8080,
			Protocol:      "TCP",
		}},
		Env: map[string]string{
			"NODE_ENV": "production",
		},
		SecretEnv: map[string]SecretRef{
			"DATABASE_URL": {
				SecretName: "api-secrets",
				SecretKey:  "database-url",
			},
		},
		Resources: Resources{
			CPURequest:    "100m",
			MemoryRequest: "128Mi",
			CPULimit:      "500m",
			MemoryLimit:   "512Mi",
		},
		Ingress: Ingress{
			Enabled: true,
			Host:    "api.dev.example.com",
			Path:    "/",
			TLS:     true,
		},
		Autoscaling: Autoscaling{
			Enabled:     true,
			MinReplicas: 2,
			MaxReplicas: 5,
			CPUPercent:  70,
		},
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return string(data)
}

func assertValidationError(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected validation error containing %q", want)
	}
	if !strings.Contains(err.Error(), want) {
		t.Fatalf("expected validation error containing %q, got %v", want, err)
	}
}
