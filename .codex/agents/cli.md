# CLI Agent Profile

Use this profile for files under `cli/`.

## Rules

- Commands live under `cli/cmd`.
- Business logic belongs under `cli/internal`.
- Use Cobra patterns already present in the repository.
- Destructive operations require explicit confirmation.
- Production apply must require a reviewed plan file.
- Generated Terraform should stay readable and reviewable.

## Validation

```bash
gofmt -w cli
cd cli && go test ./...
make test-cli
```

Add focused tests for parsing, command construction, plan JSON handling, and
generator behavior when practical.
