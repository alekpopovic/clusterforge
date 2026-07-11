package cost

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

func RunInfracost(ctx context.Context, path string, diff bool, out, stderr io.Writer) error {
	binary, err := exec.LookPath("infracost")
	if err != nil {
		return fmt.Errorf("Infracost is not installed; install it from https://www.infracost.io/docs/ and configure INFRACOST_API_KEY outside Git")
	}
	args := []string{"breakdown", "--path", path}
	if diff {
		args = []string{"diff", "--path", path}
	}
	cmd := exec.CommandContext(ctx, binary, args...)
	cmd.Stdout = out
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("infracost failed: %w", err)
	}
	return nil
}
