## Prompt 65 — Environment promotion workflow

```text
Implement environment promotion workflow documentation and CLI scaffolding.

Goal:
Support controlled promotion from dev to staging to prod.

Create docs:
- docs/promotion.md

Define:
- dev: fast iteration
- staging: production-like validation
- prod: controlled apply

CLI commands:
- cf promote plan --from dev --to staging
- cf promote diff --from staging --to prod

For this task:
- Implement only safe read-only comparison if practical.
- Compare:
  - app manifests
  - environment config
  - generated Terraform files
- Do not apply changes.
- Do not copy secrets.
- Do not mutate prod.

Behavior:
- Show differences between app image tags, replicas, ingress hosts, autoscaling config.
- Warn when prod differs from staging.
- Support --json.

Tests:
- Compare two sample app manifests.
- Detect image difference.
- Detect missing app.
- Detect ingress host difference.

Docs:
- Explain Git-based promotion.
- Explain tag pinning.
- Explain approval process.
- Explain why direct prod edits are discouraged.

Run:
- gofmt
- go test ./...
```

---
