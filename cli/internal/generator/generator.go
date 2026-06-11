package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/textracta/clusterforge/cli/internal/config"
)

func Generate(name string, env config.Environment) error {
	if env.Path == "" {
		return fmt.Errorf("environment %q has no path", name)
	}
	if err := os.MkdirAll(env.Path, 0o755); err != nil {
		return fmt.Errorf("create environment path %s: %w", env.Path, err)
	}

	readme := filepath.Join(env.Path, "README.md")
	if _, err := os.Stat(readme); err == nil {
		return nil
	}
	content := fmt.Sprintf(`# %s

Generated ClusterForge environment skeleton.

- Cloud: %s
- Region: %s
- Orchestrator: %s

Terraform/OpenTofu logic should remain visible in this directory.
`, name, env.Cloud, env.Region, env.Orchestrator)
	if err := os.WriteFile(readme, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", readme, err)
	}
	return nil
}
