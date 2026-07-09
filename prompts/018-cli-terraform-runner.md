## Prompt 18 — CLI Terraform runner

```text
Implement cli/internal/terraform runner.

Purpose:
Provide a safe wrapper around Terraform/OpenTofu CLI execution.

Files:
- runner.go
- runner_test.go if practical

Struct:
Runner {
  Binary string
  WorkDir string
  Env map[string]string
  Stdout io.Writer
  Stderr io.Writer
  Verbose bool
}

Methods:
- Init(ctx context.Context) error
- Validate(ctx context.Context) error
- Plan(ctx context.Context, outFile string, extraArgs []string) error
- Apply(ctx context.Context, planFile string, extraArgs []string) error
- Destroy(ctx context.Context, extraArgs []string) error
- Output(ctx context.Context, json bool) error
- ShowPlanJSON(ctx context.Context, planFile string) ([]byte, error)

Behavior:
- Use os/exec.
- Set working directory.
- Stream stdout/stderr.
- Return errors with command context.
- Do not include secrets in error messages beyond normal command output.
- Check binary exists using exec.LookPath in doctor or constructor.
- For plan:
  - if outFile not empty, call plan -out <outFile>
- For apply:
  - if planFile provided, call apply <planFile>
  - otherwise call apply without auto-approve
- For destroy:
  - never auto-approve by default

Also update CLI commands:
- cf init <env>
- cf plan <env> --out .cf/plans/dev.tfplan
- cf apply <env> --plan-file .cf/plans/dev.tfplan
- cf destroy <env>

Add tests for command construction if you design runner to allow dry-run command building.

Run:
- gofmt
- go test ./...
```

---
