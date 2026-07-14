# Prompt 214 — Remote execution profiles

```text
Add remote execution profiles.

Goal:
Allow different execution behavior for local, CI, dev, staging, prod, and runner jobs.

Config:
execution_profiles:
  dev-runner:
    engine: terraform
    parallelism: 10
    lock_timeout: 5m
    allowed_job_types:
      - validate
      - plan
      - drift_check
  prod-runner:
    engine: terraform
    parallelism: 3
    lock_timeout: 20m
    require_plan_file: true
    require_approval: true
    block_destroy: true
    allowed_job_types:
      - validate
      - plan
      - drift_check
      - apply

Control Plane:
- environments reference execution profile
- jobs inherit execution profile
- runner enforces profile settings

CLI:
- cf profile list
- cf profile show <name>
- cf profile validate <name>

Tests:
- prod profile requires approval
- dev profile does not allow apply if not listed
- parallelism args passed
- unknown profile fails
- runner enforces job type restrictions

Docs:
- docs/control-plane/execution-profiles.md

Rules:
- Local execution profiles remain supported.
- Remote execution must be stricter by default.
```
