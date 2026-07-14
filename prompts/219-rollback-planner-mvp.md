# Prompt 219 — Rollback planner MVP

```text
Implement rollback planner MVP.

Goal:
Generate rollback guidance based on change history and app manifests.

CLI:
- cf rollback history <env>
- cf rollback plan <env> --to-change <change-id>
- cf rollback app <app> --env <env> --to-image <image>

Control Plane:
- API endpoint:
  - GET /api/v1/environments/{id}/rollback-options
  - POST /api/v1/rollback-plans

Rollback plan should include:
- target environment
- selected previous change
- resources likely affected
- recommended Git revert or manifest changes
- Terraform plan command to run
- warnings for data resources
- approval requirement
- related runbooks

For app rollback:
- update app manifest in generated output or show patch only
- do not apply automatically
- support --write-patch to file
- support --dry-run

Rules:
- Do not perform rollback automatically.
- Do not mutate prod unless explicitly writing a patch and user confirms.
- Warn that infrastructure rollback can be unsafe.
- Data resources require manual review.

Tests:
- rollback history from change records
- app image rollback patch generated
- infrastructure rollback plan warns
- no auto-apply

Docs:
- docs/control-plane/rollback.md
```
