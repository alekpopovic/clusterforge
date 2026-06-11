package terraform

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

type Runner struct {
	Binary  string
	WorkDir string
	Verbose bool
}

func NewRunner(binary, workDir string, verbose bool) Runner {
	return Runner{
		Binary:  binary,
		WorkDir: workDir,
		Verbose: verbose,
	}
}

func (r Runner) Run(ctx context.Context, args ...string) error {
	if r.Binary == "" {
		return fmt.Errorf("engine binary is required")
	}
	if r.WorkDir == "" {
		return fmt.Errorf("work directory is required")
	}
	if r.Verbose {
		fmt.Fprintf(os.Stderr, "running: %s %v in %s\n", r.Binary, args, r.WorkDir)
	}
	cmd := exec.CommandContext(ctx, r.Binary, args...)
	cmd.Dir = r.WorkDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
