# Prompt 208 — Artifact storage implementation

```text
Implement artifact storage MVP.

Goal:
Store and retrieve sanitized Control Plane artifacts.

Backends:
- database for small artifacts
- local filesystem for larger artifacts

Do not implement S3 yet unless straightforward.

Database:
- artifacts table:
  - id
  - organization_id
  - workspace_id
  - project_id
  - environment_id nullable
  - related_type
  - related_id
  - artifact_type
  - storage_backend
  - storage_key
  - content_type
  - size_bytes
  - checksum_sha256
  - sensitive
  - created_by
  - created_at
  - expires_at

API:
- POST /api/v1/artifacts
- GET /api/v1/artifacts
- GET /api/v1/artifacts/{id}
- GET /api/v1/artifacts/{id}/download
- DELETE /api/v1/artifacts/{id}

CLI:
- cf artifact list
- cf artifact download <id> --output file
- cf artifact delete <id>

Runner:
- upload redacted logs
- upload plan summary
- upload policy results
- upload cost report
- do not upload raw plan file unless allow_raw_plan_upload=true

Security:
- RBAC protected
- audit downloads and deletes
- enforce size limits
- content checksum
- redaction required before upload
- raw plan upload blocked by default

Tests:
- upload artifact
- download artifact
- checksum verified
- size limit enforced
- RBAC read blocked
- artifact download audited

Docs:
- docs/control-plane/artifacts.md

Run:
- go test ./... in control-plane, cli, runner
```
