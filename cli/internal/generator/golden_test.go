package generator

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/textracta/clusterforge/cli/internal/config"
)

var updateGeneratorGolden = flag.Bool("update-golden", false, "update golden files")

func TestGoldenAWSEKSSimple(t *testing.T) {
	assertGoldenEnvironment(t, "aws-eks-simple", config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
	}, Options{Project: "demo"})
}

func TestGoldenAWSECSSimple(t *testing.T) {
	assertGoldenEnvironment(t, "aws-ecs-simple", config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "ecs",
	}, Options{Project: "demo"})
}

func TestGoldenExistingKubernetes(t *testing.T) {
	assertGoldenEnvironment(t, "existing-kubernetes", config.Environment{
		Cloud:        "existing",
		Region:       "local",
		Orchestrator: "kubernetes",
	}, Options{Project: "demo"})
}

func TestGoldenS3Backend(t *testing.T) {
	assertGoldenEnvironment(t, "backend-s3", config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
	}, Options{
		Project: "demo",
		Backend: config.Backend{
			Type:          "s3",
			Bucket:        "clusterforge-test-state",
			Region:        "eu-central-1",
			DynamoDBTable: "clusterforge-test-locks",
			KeyPrefix:     "clusterforge/dev",
		},
	})
}

func TestGoldenTemplatePackOverride(t *testing.T) {
	dir := t.TempDir()
	templatesDir := filepath.Join(dir, "templates")
	targetDir := filepath.Join(templatesDir, "env", "aws-eks")
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		t.Fatalf("mkdir template pack: %v", err)
	}
	for _, file := range environmentFiles {
		templateName := file + ".tmpl"
		if file == "README.md" {
			templateName = "README.md.tmpl"
		}
		body := "{{ .Header }}\n\n# template-pack-override {{ .Environment }} {{ .Project }} {{ .Name }} {{ .ModulesPath }}\n"
		if err := os.WriteFile(filepath.Join(targetDir, templateName), []byte(body), 0o644); err != nil {
			t.Fatalf("write template %s: %v", templateName, err)
		}
	}
	assertGoldenEnvironment(t, "template-pack-override", config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
	}, Options{
		Project:      "demo",
		TemplatesDir: templatesDir,
	})
}

func assertGoldenEnvironment(t *testing.T, goldenName string, env config.Environment, opts Options) {
	t.Helper()
	dir := t.TempDir()
	env.Path = filepath.Join(dir, "live", "dev", env.Cloud+"-"+env.Orchestrator)
	opts.RootDir = dir

	if _, err := Generate("dev", env, opts); err != nil {
		t.Fatalf("generate %s: %v", goldenName, err)
	}

	goldenDir := filepath.Join("..", "..", "testdata", "golden", goldenName)
	for _, file := range environmentFiles {
		actualPath := filepath.Join(env.Path, file)
		goldenPath := filepath.Join(goldenDir, file)
		assertGoldenFile(t, goldenPath, actualPath)
	}
}

func assertGoldenFile(t *testing.T, goldenPath, actualPath string) {
	t.Helper()
	actual := readFile(t, actualPath)
	if shouldUpdateGolden() {
		if err := os.MkdirAll(filepath.Dir(goldenPath), 0o755); err != nil {
			t.Fatalf("mkdir golden dir: %v", err)
		}
		if err := os.WriteFile(goldenPath, []byte(actual), 0o644); err != nil {
			t.Fatalf("update golden %s: %v", goldenPath, err)
		}
		return
	}
	expectedBytes, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("read golden %s: %v", goldenPath, err)
	}
	expected := string(expectedBytes)
	if actual != expected {
		t.Fatalf("golden mismatch for %s\nexpected:\n%s\nactual:\n%s", goldenPath, expected, actual)
	}
	if strings.Contains(actual, os.TempDir()) {
		t.Fatalf("generated file contains a machine-specific temp path: %s", actualPath)
	}
}

func shouldUpdateGolden() bool {
	return *updateGeneratorGolden || os.Getenv("CLUSTERFORGE_UPDATE_GOLDEN") == "true"
}
