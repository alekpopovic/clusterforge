package generator

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/textracta/clusterforge/cli/internal/config"
)

func TestGenerateAWSEKS(t *testing.T) {
	dir := t.TempDir()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
		Path:         filepath.Join(dir, "live", "dev", "aws-eks"),
	}

	result, err := Generate("dev", env, Options{
		RootDir: dir,
		Project: "demo",
	})
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if result.Target != "aws-eks" {
		t.Fatalf("target = %q", result.Target)
	}
	for _, file := range environmentFiles {
		path := filepath.Join(env.Path, file)
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected %s to exist: %v", file, err)
		}
	}
	mainTF := readFile(t, filepath.Join(env.Path, "main.tf"))
	if !strings.Contains(mainTF, `module "eks"`) {
		t.Fatal("main.tf does not include EKS module")
	}
	if !strings.Contains(mainTF, "../../../modules/orchestrators/kubernetes/eks") {
		t.Fatalf("main.tf has wrong relative module path:\n%s", mainTF)
	}
	backendTF := readFile(t, filepath.Join(env.Path, "backend.tf"))
	if !strings.Contains(backendTF, `backend "local"`) {
		t.Fatalf("backend.tf does not include local backend:\n%s", backendTF)
	}
}

func TestGenerateS3Backend(t *testing.T) {
	dir := t.TempDir()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
		Path:         filepath.Join(dir, "live", "prod", "aws-eks"),
	}

	_, err := Generate("prod", env, Options{
		RootDir: dir,
		Backend: config.Backend{
			Type:          "s3",
			Bucket:        "example-terraform-state",
			Region:        "eu-central-1",
			DynamoDBTable: "example-terraform-locks",
			KeyPrefix:     "clusterforge/prod",
		},
	})
	if err != nil {
		t.Fatalf("generate s3 backend: %v", err)
	}
	backendTF := readFile(t, filepath.Join(env.Path, "backend.tf"))
	for _, expected := range []string{
		`backend "s3"`,
		`bucket         = "example-terraform-state"`,
		`key            = "clusterforge/prod/terraform.tfstate"`,
		`region         = "eu-central-1"`,
		`dynamodb_table = "example-terraform-locks"`,
		`encrypt        = true`,
	} {
		if !strings.Contains(backendTF, expected) {
			t.Fatalf("backend.tf missing %q:\n%s", expected, backendTF)
		}
	}
}

func TestGenerateMissingS3BucketFails(t *testing.T) {
	dir := t.TempDir()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
		Path:         filepath.Join(dir, "live", "prod", "aws-eks"),
	}

	_, err := Generate("prod", env, Options{
		RootDir: dir,
		Backend: config.Backend{
			Type:   "s3",
			Region: "eu-central-1",
		},
	})
	if err == nil {
		t.Fatal("expected missing s3 bucket to fail")
	}
}

func TestGenerateProdLocalBackendWarns(t *testing.T) {
	dir := t.TempDir()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
		Path:         filepath.Join(dir, "live", "prod", "aws-eks"),
	}
	var out bytes.Buffer

	if _, err := Generate("prod", env, Options{
		RootDir: dir,
		Stdout:  &out,
	}); err != nil {
		t.Fatalf("generate prod local backend: %v", err)
	}
	if !strings.Contains(out.String(), "local backend") {
		t.Fatalf("expected local backend warning, got %q", out.String())
	}
}

func TestGenerateStackedAWSEKS(t *testing.T) {
	dir := t.TempDir()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
		Path:         filepath.Join(dir, "live", "dev", "aws-eks"),
		Layout:       "stacked",
	}

	result, err := Generate("dev", env, Options{
		RootDir: dir,
		Project: "demo",
	})
	if err != nil {
		t.Fatalf("generate stacked: %v", err)
	}
	if result.Target != "aws-eks-stacked" {
		t.Fatalf("target = %q", result.Target)
	}
	for _, stack := range config.StackOrder() {
		for _, file := range environmentFiles {
			path := filepath.Join(env.Path, stack, file)
			if _, err := os.Stat(path); err != nil {
				t.Fatalf("expected %s to exist: %v", path, err)
			}
		}
	}
	mainTF := readFile(t, filepath.Join(env.Path, "network", "main.tf"))
	if !strings.Contains(mainTF, "Network stack") {
		t.Fatalf("network stack main.tf missing stack title:\n%s", mainTF)
	}
	if !strings.Contains(mainTF, "../../../../modules") {
		t.Fatalf("network stack main.tf has wrong relative module path:\n%s", mainTF)
	}
	backendTF := readFile(t, filepath.Join(env.Path, "network", "backend.tf"))
	if !strings.Contains(backendTF, `backend "local"`) {
		t.Fatalf("stack backend.tf does not include local backend:\n%s", backendTF)
	}
}

func TestGenerateRefusesOverwriteWithoutForce(t *testing.T) {
	dir := t.TempDir()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "ecs",
		Path:         filepath.Join(dir, "live", "dev", "aws-ecs"),
	}
	if err := os.MkdirAll(env.Path, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(env.Path, "main.tf"), []byte("existing"), 0o644); err != nil {
		t.Fatalf("write existing file: %v", err)
	}

	if _, err := Generate("dev", env, Options{RootDir: dir}); err == nil {
		t.Fatal("expected overwrite refusal")
	}
}

func TestGenerateDryRunWritesNothing(t *testing.T) {
	dir := t.TempDir()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
		Path:         filepath.Join(dir, "live", "dev", "aws-eks"),
	}
	var out bytes.Buffer

	if _, err := Generate("dev", env, Options{
		RootDir: dir,
		DryRun:  true,
		Stdout:  &out,
	}); err != nil {
		t.Fatalf("dry-run generate: %v", err)
	}
	if _, err := os.Stat(env.Path); !os.IsNotExist(err) {
		t.Fatalf("dry-run wrote environment path, stat err: %v", err)
	}
	if !strings.Contains(out.String(), "create") {
		t.Fatalf("dry-run output = %q", out.String())
	}
}

func TestGenerateDryRunReportsExistingFilesWithoutForce(t *testing.T) {
	dir := t.TempDir()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "ecs",
		Path:         filepath.Join(dir, "live", "dev", "aws-ecs"),
	}
	if err := os.MkdirAll(env.Path, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(env.Path, "main.tf"), []byte("existing"), 0o644); err != nil {
		t.Fatalf("write existing file: %v", err)
	}
	var out bytes.Buffer

	if _, err := Generate("dev", env, Options{
		RootDir: dir,
		DryRun:  true,
		Stdout:  &out,
	}); err != nil {
		t.Fatalf("dry-run generate: %v", err)
	}
	if !strings.Contains(out.String(), "update "+filepath.Join(env.Path, "main.tf")) {
		t.Fatalf("dry-run output = %q", out.String())
	}
	if got := readFile(t, filepath.Join(env.Path, "main.tf")); got != "existing" {
		t.Fatalf("dry-run changed existing file: %q", got)
	}
}

func TestGenerateUnknownTarget(t *testing.T) {
	dir := t.TempDir()
	env := config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "nomad",
		Path:         filepath.Join(dir, "live", "dev", "aws-nomad"),
	}

	_, err := Generate("dev", env, Options{RootDir: dir})
	if err == nil {
		t.Fatal("expected unsupported target error")
	}
	if !strings.Contains(err.Error(), "unsupported generate target") {
		t.Fatalf("unexpected error: %v", err)
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
