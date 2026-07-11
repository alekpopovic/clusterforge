## Prompt 141 — Cloud asset inventory export

```text
Add cloud asset inventory export support.

Goal:
Export ClusterForge-managed infrastructure inventory for audit, reporting, and CMDB integration.

CLI:
- cf inventory export <env>
- cf inventory export <env> --stack <stack>
- cf inventory export --fleet
- cf inventory export --format json|csv|markdown
- cf inventory export --output inventory.json

Data sources:
- Terraform state when available
- Terraform plan JSON when available
- ClusterForge config
- Module catalog
- App manifests

Inventory fields:
- address
- type
- name
- provider
- module
- environment
- stack
- cloud
- region
- tags/labels
- sensitive flag if known
- source

Rules:
- Do not print sensitive values.
- State may contain secrets, so redact aggressively.
- Read-only.
- No cloud API calls for MVP unless explicit.

Docs:
- docs/inventory-export.md

Tests:
- State fixture parses resources.
- CSV output.
- JSON output.
- Sensitive attributes redacted.

Run:
- gofmt
- go test ./...
```


---
