# Prompt 258 — Job fairness and tenant quotas

```text
Implement job fairness and tenant-level job quotas.

Goal:
Prevent one organization/project from starving others.

Features:
- per-organization active job limit
- per-project active job limit
- per-environment active job limit
- per-job-type limit
- priority classes:
  - low
  - normal
  - high
  - emergency
- fair scheduling across organizations
- FIFO within same priority/scope where practical

Config:
job_scheduling:
  max_active_jobs_per_org: 20
  max_active_jobs_per_project: 10
  max_active_jobs_per_environment: 3
  max_active_apply_jobs_per_environment: 1
  default_priority: normal

Control Plane:
- scheduler picks eligible job based on:
  - priority
  - quotas
  - runner pool
  - required labels
  - lease availability

API:
- GET /api/v1/job-scheduling/status
- GET /api/v1/job-scheduling/queues

CLI:
- cf job queue
- cf job priority set <job-id> --priority high

Tests:
- org active job limit enforced
- environment apply concurrency is 1
- high priority job scheduled first
- fair scheduling across orgs
- runner labels still enforced

Docs:
- docs/control-plane/job-scheduling.md

Rules:
- Emergency priority requires admin.
- Apply jobs remain conservative.
```
