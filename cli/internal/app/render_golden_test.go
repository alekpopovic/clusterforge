package app

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/textracta/clusterforge/cli/internal/config"
)

var updateAppGolden = flag.Bool("update-golden", false, "update golden files")

func TestGoldenKubernetesAppRender(t *testing.T) {
	assertGoldenAppRender(t, "app-kubernetes", config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "eks",
	})
}

func TestGoldenECSAppRender(t *testing.T) {
	assertGoldenAppRender(t, "app-ecs", config.Environment{
		Cloud:        "aws",
		Region:       "eu-central-1",
		Orchestrator: "ecs",
	})
}

func assertGoldenAppRender(t *testing.T, goldenName string, env config.Environment) {
	t.Helper()
	dir := t.TempDir()
	env.Path = filepath.Join(dir, "live", "dev", "aws-"+env.Orchestrator)

	outPath, err := Render(dir, "dev", env, sampleManifest())
	if err != nil {
		t.Fatalf("render %s: %v", goldenName, err)
	}

	goldenPath := filepath.Join("..", "..", "testdata", "golden", goldenName, "api.tf")
	actual := readFile(t, outPath)
	if shouldUpdateAppGolden() {
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
	if actual != string(expectedBytes) {
		t.Fatalf("golden mismatch for %s\nexpected:\n%s\nactual:\n%s", goldenPath, string(expectedBytes), actual)
	}
}

func shouldUpdateAppGolden() bool {
	return *updateAppGolden || os.Getenv("CLUSTERFORGE_UPDATE_GOLDEN") == "true"
}
