# Prompt 225 — Preview environments MVP for Kubernetes

```text
Implement preview environments MVP for Kubernetes.

Goal:
Allow app preview deployments into existing Kubernetes cluster namespace-per-PR.

CLI:
- cf preview create --app api --pr 123 --env dev
- cf preview delete --app api --pr 123 --env dev
- cf preview list --env dev
- cf preview cleanup --env dev

Behavior:
- create namespace:
  preview-<app>-pr-<number>
- render workload module with unique name/namespace
- optional ingress host:
  pr-123-api.dev.example.com
- label all resources:
  clusterforge.io/preview=true
  clusterforge.io/pr=123
  clusterforge.io/app=api
- TTL annotation
- cleanup deletes generated preview files or Terraform stack

Config:
preview:
  enabled: true
  ttl_hours: 48
  base_domain: dev.example.com
  namespace_prefix: preview

Rules:
- Kubernetes only for MVP.
- No production previews by default.
- Do not copy production secrets.
- Require explicit preview secret strategy.
- Cleanup must be safe and scoped by labels/names.

Tests:
- render preview
- list preview metadata
- delete preview files
- cleanup expired preview
- prod preview blocked

Docs:
- docs/preview-environments.md

Run:
- gofmt
- go test ./...
```
