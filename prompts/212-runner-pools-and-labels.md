# Prompt 212 — Runner pools and labels

```text
Implement runner pools and labels.

Goal:
Route jobs to appropriate runners based on environment, cloud, capabilities, and trust level.

Config:
runner:
  name: prod-aws-runner-1
  pool: prod-aws
  labels:
    cloud: aws
    region: eu-central-1
    network: private
    trust: production
  allowed_job_types:
    - validate
    - policy_check
    - plan
    - drift_check
    - cost_scan
    - apply

Control Plane:
- runner_pools table
- runner_labels
- job required_labels
- job runner_pool

API:
- GET /api/v1/runner-pools
- POST /api/v1/runner-pools
- GET /api/v1/runners
- PATCH /api/v1/runners/{id}/labels

CLI:
- cf runner pool list
- cf runner pool create <name>
- cf runner labels <runner>
- cf runner assign <runner> --pool prod-aws

Scheduler:
- runner can only claim jobs from its pool
- required labels must match
- prod jobs require production-trust runner if policy enabled

Tests:
- matching runner claims job
- non-matching runner cannot claim job
- prod job requires prod runner
- labels update audited

Docs:
- docs/control-plane/runner-pools.md

Rules:
- Do not let dev runner claim prod jobs by default.
- Runner pool assignment requires admin.
```
