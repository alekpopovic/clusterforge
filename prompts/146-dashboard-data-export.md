## Prompt 146 — Dashboard data export

```text
Implement dashboard data export.

Goal:
Prepare data for a future ClusterForge dashboard without building the full web app.

CLI:
- cf dashboard export
- cf dashboard export --env prod
- cf dashboard export --fleet
- cf dashboard export --output dashboard-data.json

Export data:
- project metadata
- organization/workspace metadata
- environments
- clusters
- stacks
- apps
- service catalog
- policy results if available
- drift status if available
- cost warnings if available
- runbook index
- module catalog

Output:
dashboard-data.json

Rules:
- No secrets.
- Redact sensitive values.
- Read-only.
- Do not call cloud APIs unless explicitly implemented and safe.
- Should work offline from config and files.

Tests:
- Export minimal project.
- Export project with apps.
- Redaction test.
- JSON schema stability.

Docs:
- docs/dashboard-export.md

Run:
- gofmt
- go test ./...
```


---
