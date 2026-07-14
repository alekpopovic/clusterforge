# Prompt 207 — Artifact storage RFC

```text
Create artifact storage RFC for ClusterForge Control Plane.

Create:
- docs/rfcs/020-artifact-storage.md
- docs/control-plane/artifacts.md

Goal:
Design secure storage for plan summaries, redacted logs, policy results, cost reports, and optional Terraform plan files.

Artifacts:
- plan summary JSON
- policy result JSON/SARIF
- cost report JSON
- drift result JSON
- redacted logs
- raw Terraform plan file optional and disabled by default
- apply logs
- generated inventory
- bundle manifests

Storage backends:
1. database for small JSON artifacts
2. local filesystem for dev
3. S3-compatible object storage
4. Azure Blob future
5. GCS future

Security:
- raw plan files may contain sensitive values
- artifact encryption
- signed URLs
- retention policy
- per-project scoping
- access controlled by RBAC
- audit downloads
- redaction before upload
- size limits

Config:
artifacts:
  backend: filesystem
  filesystem:
    path: .cf/artifacts
  retention_days: 30
  allow_raw_plan_upload: false
  max_artifact_size_mb: 50

Do not implement code in this prompt.
Define:
- data model
- artifact lifecycle
- retention policy
- security model
- API endpoints
- runner upload flow
- CLI download flow

Update:
- ROADMAP_V0.6.md
```
