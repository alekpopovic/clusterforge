## Prompt 110 — Fleet operations CLI

```text
Add read-only fleet operations commands.

Goal:
Allow operators to inspect multiple ClusterForge environments/clusters.

Commands:
- cf fleet status
- cf fleet doctor
- cf fleet drift
- cf fleet policy check
- cf fleet list --json

Behavior:
- Iterate over configured environments/clusters.
- Support filters:
  --environment dev
  --cloud aws
  --orchestrator eks
  --status production
- Default operations must be read-only.
- No apply.
- No destroy.
- Print summary table.
- Support JSON output.

For drift:
- Run drift check per environment/stack if configured.
- Continue on error and report per-cluster failure.
- Do not fail the whole command unless --fail-fast.

Docs:
- docs/fleet-operations.md

Tests:
- Fleet list with filters.
- Fleet doctor aggregates checks.
- Fleet drift handles one failing env.

Rules:
- Read-only only.
- No background daemon.
- No cloud credentials in tests.

Run:
- gofmt
- go test ./...
```

---
