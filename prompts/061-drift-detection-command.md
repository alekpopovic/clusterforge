## Prompt 61 — Drift detection command

```text
Add drift detection support to ClusterForge CLI.

Command:
- cf drift check <env>
- cf drift check <env> --stack <stack>
- cf drift check <env> --json

Goal:
Detect infrastructure drift using Terraform/OpenTofu plan.

Behavior:
- Run plan with detailed exit code.
- Exit codes:
  - 0: no drift
  - 2: drift/changes detected
  - 1: error
- Print human summary.
- Support JSON output.
- Do not apply changes.
- For prod, print strong warning that drift must be reviewed.

Implementation:
- Extend terraform runner to support plan -detailed-exitcode.
- Parse plan JSON if --json or --summary is requested.
- Reuse policy/risk parser.
- Add tests for exit code handling.

Docs:
- docs/drift-detection.md
- Explain:
  - what drift is
  - why plan detects it
  - limitations
  - CI scheduled drift check pattern

GitHub Actions:
- Add example workflow:
  .github/workflows/drift-check-example.yml
- It should be disabled/commented or documented as requiring credentials.

Rules:
- No apply.
- No auto-remediation.
- No cloud credentials in repo.

Run:
- gofmt
- go test ./...
```

---
