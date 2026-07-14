# Prompt 256 — Ephemeral runner pools

```text
Implement ephemeral runner pools design and MVP.

Goal:
Support temporary runners for isolated job execution.

Use cases:
- high-risk prod plan
- short-lived preview environment jobs
- per-tenant isolation
- air-gapped job execution
- burst capacity

Control Plane:
- runner_pool type:
  - persistent
  - ephemeral
- runner lifecycle:
  - requested
  - starting
  - active
  - draining
  - terminated
  - failed

API:
- POST /api/v1/runner-pools/{id}/scale-up
- POST /api/v1/runner-pools/{id}/drain
- GET /api/v1/runner-pools/{id}/capacity

Runner:
- supports drain mode
- stops claiming new jobs
- finishes active jobs
- exits when idle if ephemeral

CLI:
- cf runner pool scale-up <pool>
- cf runner pool drain <pool>
- cf runner drain <runner>

Docs:
- docs/control-plane/ephemeral-runners.md

Tests:
- draining runner stops claiming jobs
- active job can finish
- ephemeral runner exits when idle
- pool capacity reported

Rules:
- Do not auto-scale cloud resources yet unless explicit.
- No apply on untrusted ephemeral runners.
- Token expiration should be short for ephemeral runners.
```
