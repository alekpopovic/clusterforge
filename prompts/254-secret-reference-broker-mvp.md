# Prompt 254 — Secret reference broker MVP

```text
Implement secret reference broker MVP.

Goal:
Centralize secret references and validation without reading secret values.

Control Plane:
- secret_references table:
  - id
  - organization_id
  - project_id
  - environment_id nullable
  - app_id nullable
  - provider
  - reference
  - key_name nullable
  - usage_type
  - source_path
  - metadata_json
  - created_at
  - updated_at

Providers:
- kubernetes_secret
- aws_secrets_manager
- aws_ssm_parameter
- external_secrets
- vault
- ecs_secret

API:
- GET /api/v1/secret-references
- POST /api/v1/secret-references/import
- GET /api/v1/secret-references/{id}

CLI:
- cf secrets inventory
- cf secrets sync
- cf secrets validate-references
- cf secrets report --format markdown|json

Validation:
- reference format validation only
- no secret value reads
- optional provider existence check only when explicitly enabled

Dashboard:
- secret references page
- app secret references
- environment secret references

Tests:
- import from app manifests
- import ECS value_from
- import ExternalSecret manifest
- Vault path reference
- no values stored
- cross-tenant access denied

Docs:
- docs/control-plane/secret-reference-broker.md

Rules:
- Never read secret values.
- Never display secret values.
- This is inventory and validation only.
```
