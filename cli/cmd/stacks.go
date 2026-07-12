package cmd

import (
	"fmt"

	"github.com/alekpopovic/clusterforge/cli/internal/config"
)

func resolveStackPaths(env config.Environment, stack string) ([]string, error) {
	return env.StackPaths(stack)
}

func resolveSingleStackPath(env config.Environment, stack string) (string, error) {
	return env.StackPath(stack)
}

func stackLabel(path string, total int) string {
	if total <= 1 {
		return ""
	}
	return fmt.Sprintf("==> %s", path)
}
