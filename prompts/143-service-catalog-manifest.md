## Prompt 143 — Service catalog manifest

```text
Add ClusterForge service catalog manifest support.

Goal:
Provide an internal source of truth for applications, owners, dependencies, environments, and operational metadata.

Create:
- service-catalog.yaml schema
- docs/service-catalog.md

Example:
services:
  api:
    owner: payments-team
    tier: backend
    lifecycle: production
    repositories:
      source: https://github.com/example/api
      image: ghcr.io/example/api
    environments:
      dev:
        url: https://api.dev.example.com
      prod:
        url: https://api.example.com
    dependencies:
      - postgres
      - redis
      - sqs-jobs
    alerts:
      dashboard: ""
      runbook: failed-deployment

CLI:
- cf service list
- cf service show <name>
- cf service validate
- cf service export --format json|markdown
- cf service graph --format dot

Integration:
- app manifests may reference service catalog entries.
- Backstage generation can use service catalog data.
- Runbook links can use service catalog.

Rules:
- Metadata only.
- No secrets.
- No remote API.
- Keep schema simple.

Tests:
- Validate service catalog.
- List services.
- Export JSON.
- Service graph.

Run:
- gofmt
- go test ./...
```


---
