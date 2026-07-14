# Prompt 264 — Compliance evidence collection

```text
Implement compliance evidence collection.

Goal:
Collect non-sensitive evidence artifacts showing how ClusterForge controls are configured and operating.

Evidence types:
- policy check results
- approval records
- change records
- audit event exports
- RBAC bindings
- token rotation status
- artifact retention config
- backup evidence files
- DR drill evidence
- module conformance reports
- security scan summaries
- release gate reports
- smoke test matrix
- compliance mapping reports

Database:
- evidence_records
  - id
  - organization_id
  - project_id nullable
  - environment_id nullable
  - evidence_type
  - title
  - description
  - source_type
  - source_id nullable
  - artifact_id nullable
  - created_by
  - created_at
  - metadata_json

API:
- GET /api/v1/evidence
- POST /api/v1/evidence/collect
- GET /api/v1/evidence/{id}

CLI:
- cf evidence collect
- cf evidence list
- cf evidence show <id>
- cf evidence export --format markdown|json

Dashboard:
- evidence page

Tests:
- collect policy evidence
- collect approval evidence
- evidence export
- RBAC protects evidence
- no secrets in evidence metadata

Docs:
- docs/control-plane/evidence-collection.md

Rules:
- Do not claim compliance certification.
- Evidence is support material only.
- Redact sensitive data.
```
