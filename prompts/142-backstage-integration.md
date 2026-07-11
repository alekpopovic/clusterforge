## Prompt 142 — Backstage integration

```text
Add Backstage integration support.

Goal:
Allow ClusterForge projects, apps, and environments to generate Backstage catalog entities.

Create:
- docs/backstage.md
- examples/backstage/catalog-info.yaml

CLI:
- cf backstage generate
- cf backstage generate --app api
- cf backstage generate --env prod
- cf backstage generate --output catalog-info.yaml

Generated entities:
- Component for app
- Resource for cluster/environment
- System for project
- Group/team mapping if teams exist

Config:
clusterforge.yaml:
  backstage:
    enabled: true
    owner: platform-team
    system: payments-platform
    lifecycle: production

App manifest extension:
backstage:
  owner: payments-team
  system: payments-platform
  lifecycle: production

Rules:
- Generate YAML only.
- Do not call Backstage API.
- Do not include secrets.
- Keep fields configurable.

Tests:
- Generate project catalog.
- Generate app catalog.
- Owner fallback works.
- YAML validates structurally.

Docs:
- Explain how to commit catalog-info.yaml.
- Explain entity model.
- Explain ownership.

Run:
- gofmt
- go test ./...
```


---
