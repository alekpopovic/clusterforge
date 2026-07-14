# Prompt 277 — GitOps reconciliation status ingestion

```text
Implement GitOps reconciliation status ingestion.

Goal:
Allow Control Plane to display Argo CD or Flux sync/health status without becoming the GitOps controller.

Supported MVP:
- Argo CD Application status from Kubernetes API or exported JSON
- Flux Kustomization/HelmRelease status optional

CLI:
- cf gitops status --env prod
- cf gitops export-status --format json
- cf api push-gitops-status

Control Plane:
- gitops_app_status table:
  - id
  - project_id
  - environment_id
  - cluster_id
  - provider
  - app_name
  - sync_status
  - health_status
  - revision
  - message
  - observed_at
  - metadata_json

API:
- GET /api/v1/gitops/status
- POST /api/v1/gitops/status/import

Dashboard:
- GitOps status page
- environment detail shows GitOps status

Tests:
- import Argo CD app fixture
- import Flux fixture if implemented
- status visible by environment
- cross-tenant denied

Docs:
- docs/gitops-status.md

Rules:
- Read-only.
- Do not trigger sync.
- Do not store Git credentials.
- Do not replace Argo CD/Flux.
```
