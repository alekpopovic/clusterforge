# Prompt 211 — Job queue hardening

```text
Harden the Control Plane job queue.

Goal:
Make plan, drift, cost, policy, and apply jobs reliable and observable.

Job model:
- jobs table:
  - id
  - organization_id
  - project_id
  - environment_id
  - type
  - status
  - priority
  - queued_at
  - claimed_at
  - started_at
  - finished_at
  - runner_id
  - attempts
  - max_attempts
  - timeout_seconds
  - lease_expires_at
  - payload_json
  - result_json
  - error_message

Statuses:
- queued
- claimed
- running
- succeeded
- failed
- canceled
- timed_out
- expired

Features:
- job lease
- heartbeat extends lease
- timeout handling
- retry policy
- cancellation
- priority
- per-environment concurrency limit
- per-runner allowed job types

API:
- GET /api/v1/jobs
- GET /api/v1/jobs/{id}
- POST /api/v1/jobs/{id}/cancel
- runner claim/heartbeat/complete/fail endpoints updated

CLI:
- cf job list
- cf job show <id>
- cf job cancel <id>

Tests:
- job claim
- lease expiration
- timeout
- retry
- cancellation
- per-environment concurrency
- runner cannot claim unsupported job type

Docs:
- docs/control-plane/jobs.md

Rules:
- Apply jobs require approval.
- No automatic retry for apply by default.
- All job state transitions audited.
```
