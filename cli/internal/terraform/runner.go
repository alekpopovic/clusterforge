package terraform

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Runner struct {
	Binary  string
	WorkDir string
	Env     map[string]string
	Stdout  io.Writer
	Stderr  io.Writer
	Verbose bool
}

func NewRunner(binary, workDir string, verbose bool) Runner {
	return Runner{
		Binary:  binary,
		WorkDir: workDir,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
		Verbose: verbose,
	}
}

func (r Runner) Init(ctx context.Context) error {
	return r.run(ctx, initArgs())
}

func (r Runner) Validate(ctx context.Context) error {
	return r.run(ctx, validateArgs())
}

func (r Runner) Plan(ctx context.Context, outFile string, extraArgs []string) error {
	return r.run(ctx, planArgs(outFile, extraArgs))
}

func (r Runner) Apply(ctx context.Context, planFile string, extraArgs []string) error {
	return r.run(ctx, applyArgs(planFile, extraArgs))
}

func (r Runner) Destroy(ctx context.Context, extraArgs []string) error {
	return r.run(ctx, destroyArgs(extraArgs))
}

func (r Runner) Output(ctx context.Context, json bool) error {
	return r.run(ctx, outputArgs(json))
}

func (r Runner) ShowPlanJSON(ctx context.Context, planFile string) ([]byte, error) {
	if planFile == "" {
		return nil, fmt.Errorf("plan file is required")
	}

	var stdout bytes.Buffer
	copy := r
	copy.Stdout = &stdout
	if err := copy.run(ctx, showPlanJSONArgs(planFile)); err != nil {
		return nil, err
	}
	return stdout.Bytes(), nil
}

func (r Runner) Run(ctx context.Context, args ...string) error {
	return r.run(ctx, args)
}

func (r Runner) run(ctx context.Context, args []string) error {
	if r.Binary == "" {
		return fmt.Errorf("engine binary is required")
	}
	if r.WorkDir == "" {
		return fmt.Errorf("work directory is required")
	}
	if len(args) == 0 {
		return fmt.Errorf("command arguments are required")
	}
	binary, err := exec.LookPath(r.Binary)
	if err != nil {
		return fmt.Errorf("find engine binary %q: %w", r.Binary, err)
	}

	stdout := r.Stdout
	if stdout == nil {
		stdout = os.Stdout
	}
	stderr := r.Stderr
	if stderr == nil {
		stderr = os.Stderr
	}

	if r.Verbose {
		fmt.Fprintf(stderr, "running %s %s in %s\n", r.Binary, args[0], r.WorkDir)
	}
	cmd := exec.CommandContext(ctx, binary, args...)
	cmd.Dir = r.WorkDir
	cmd.Env = mergeEnv(os.Environ(), r.Env)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s failed in %s: %w", r.Binary, args[0], r.WorkDir, err)
	}
	return nil
}

func initArgs() []string {
	return []string{"init"}
}

func validateArgs() []string {
	return []string{"validate"}
}

func planArgs(outFile string, extraArgs []string) []string {
	args := []string{"plan"}
	if outFile != "" {
		args = append(args, "-out", outFile)
	}
	return append(args, extraArgs...)
}

func applyArgs(planFile string, extraArgs []string) []string {
	args := []string{"apply"}
	if planFile != "" {
		args = append(args, planFile)
	}
	return append(args, extraArgs...)
}

func destroyArgs(extraArgs []string) []string {
	return append([]string{"destroy"}, extraArgs...)
}

func outputArgs(json bool) []string {
	args := []string{"output"}
	if json {
		args = append(args, "-json")
	}
	return args
}

func showPlanJSONArgs(planFile string) []string {
	return []string{"show", "-json", planFile}
}

func mergeEnv(base []string, overrides map[string]string) []string {
	if len(overrides) == 0 {
		return base
	}
	env := append([]string{}, base...)
	for key, value := range overrides {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	return env
}
