# Prompt 170 — Plan request workflow

```text
Implement plan request workflow in Control Plane, CLI, and Runner.

Goal:
Allow users to request Terraform plans through the Control Plane.

Control Plane:
- POST /api/v1/plan-requests
- GET /api/v1/plan-requests
- GET /api/v1/plan-requests/{id}
- POST /api/v1/plan-requests/{id}/cancel

PlanRequest fields:
- project_id
- environment_id
- stack
- git_ref
- status:
  - pending
  - claimed
  - running
  - succeeded
  - failed
  - canceled
- requested_by
- summary_json
- logs_redacted
- created_at
- updated_at

CLI:
- cf plan request <env>
- cf plan status <request-id>
- cf plan logs <request-id>
- cf plan list

Runner:
- support plan job
- clone repo or use local path if configured
- run cf plan or terraform plan
- generate plan summary
- upload sanitized summary
- do not upload raw plan file unless artifact storage policy is implemented

Security:
- plan artifacts can contain sensitive data
- raw plan file upload disabled by default
- logs must be redacted

Tests:
- create plan request
- runner processes plan request with fake executor
- plan status
- failed plan result
- redaction works

Docs:
- docs/control-plane-plan-workflow.md

Rules:
- No apply.
- No auto-approve.
- No raw secrets in logs.
- Make local fake executor available for tests.

Run:
- go test ./... in control-plane, cli, runner
```
