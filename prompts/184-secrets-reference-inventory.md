# Prompt 184 — Secrets reference inventory

```text
Implement secrets reference inventory.

Goal:
Find where secrets are referenced without reading or storing secret values.

CLI:
- cf secrets inventory
- cf secrets inventory --env prod
- cf secrets inventory --app api
- cf secrets inventory --format json|markdown

Sources:
- app manifests secret_env
- ECS secrets value_from
- External Secrets manifests
- Terraform files referencing secret ARNs
- Kubernetes workload modules
- service catalog dependencies

Output fields:
- reference name
- provider:
  - kubernetes
  - aws-secrets-manager
  - ssm
  - external-secrets
  - ecs
- environment
- app
- path
- key/name
- value is never shown

Control Plane:
- optional API endpoint:
  - POST /api/v1/secret-references
  - GET /api/v1/secret-references

Dashboard:
- optional secrets reference page
- no secret values

Tests:
- Kubernetes secret_env detected
- ECS secret ARN detected
- ExternalSecret detected
- no secret values printed

Docs:
- docs/secrets-inventory.md

Rules:
- Never read secret values.
- Never ask for secret values.
- Redact aggressively.
```
