package secrets

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFindKubernetesAndECSReferencesWithoutValues(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "apps"), 0o755); err != nil {
		t.Fatal(err)
	}
	manifest := "name: api\ntype: web\nimage: example/api:1\nsecret_env:\n  DATABASE_URL:\n    secret_name: api-secret\n    secret_key: database-url\n"
	if err := os.WriteFile(filepath.Join(root, "apps", "api.yaml"), []byte(manifest), 0o600); err != nil {
		t.Fatal(err)
	}
	env := filepath.Join(root, "live", "prod")
	if err := os.MkdirAll(env, 0o755); err != nil {
		t.Fatal(err)
	}
	terraform := `secrets = [{ name = "TOKEN", value_from = "arn:aws:secretsmanager:region:account:secret:api" }]`
	if err := os.WriteFile(filepath.Join(env, "main.tf"), []byte(terraform), 0o600); err != nil {
		t.Fatal(err)
	}
	refs, err := Discover(root, env, "api")
	if err != nil {
		t.Fatal(err)
	}
	if len(refs) != 2 {
		t.Fatalf("refs=%#v", refs)
	}
	joined := ""
	for _, ref := range refs {
		joined += ref.Name + ref.Key
	}
	if !strings.Contains(joined, "api-secretdatabase-url") || !strings.Contains(joined, "arn:aws:secretsmanager") {
		t.Fatalf("refs=%#v", refs)
	}
	if strings.Contains(joined, "super-secret-value") {
		t.Fatal("secret value was printed")
	}
}

func TestRotationPlan(t *testing.T) {
	plan := strings.Join(RotationPlan("prod", []Reference{{Name: "api-secret"}}), "\n")
	for _, expected := range []string{"prod", "external secret store", "Restart Kubernetes pods", "redeploy ECS tasks", "Revoke"} {
		if !strings.Contains(plan, expected) {
			t.Fatalf("plan missing %q: %s", expected, plan)
		}
	}
}
