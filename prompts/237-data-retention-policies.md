# Prompt 237 — Data retention policies

```text
Implement data retention policies for Control Plane.

Goal:
Control retention of audit events, artifacts, logs, plan results, policy results, drift results, and cost reports.

Config:
retention:
  audit_events_days: 365
  artifacts_days: 30
  job_logs_days: 30
  policy_results_days: 180
  drift_results_days: 180
  cost_reports_days: 180
  deleted_records_tombstone_days: 90

Behavior:
- scheduled cleanup worker
- dry-run mode
- audit cleanup summary
- do not delete records needed for compliance unless configured
- retention can be disabled per category

API:
- GET /api/v1/retention
- POST /api/v1/retention/cleanup?dry_run=true

CLI:
- cf retention show
- cf retention cleanup --dry-run
- cf retention cleanup --execute

Tests:
- expired policy results cleaned
- audit retention respected
- dry-run deletes nothing
- cleanup summary
- retention config validation

Docs:
- docs/control-plane/data-retention.md

Rules:
- Conservative defaults.
- No silent deletion.
- Production cleanup should be explicit or scheduled with docs.
```
