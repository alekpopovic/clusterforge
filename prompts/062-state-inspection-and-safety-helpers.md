## Prompt 62 — State inspection and safety helpers

```text
Add safe Terraform state helper commands to ClusterForge CLI.

Commands:
- cf state list <env> [--stack <stack>]
- cf state show <env> <address> [--stack <stack>]
- cf state pull <env> [--stack <stack>] --output <file>

Goal:
Expose read-only state operations safely.

Rules:
- No state rm.
- No state mv.
- No direct state editing.
- No default printing of full state JSON to terminal unless explicitly requested.
- Warn that state may contain sensitive data.
- For state pull:
  - require --output
  - refuse output path inside git repo unless --allow-repo-output is passed, or print strong warning

Implementation:
- Add methods to terraform runner:
  - StateList
  - StateShow
  - StatePull
- Add policy warning for prod.
- Add tests for command construction.

Docs:
- docs/state.md
- Explain state sensitivity.
- Explain safe operations.
- Explain why direct state editing is discouraged.

Run:
- gofmt
- go test ./...
```

---
