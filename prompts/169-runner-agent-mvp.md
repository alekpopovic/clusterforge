# Prompt 169 — Runner agent MVP

```text
Implement ClusterForge runner agent MVP.

Location:
- runner/

Language:
- Go

Binary:
- clusterforge-runner

Runner commands:
- clusterforge-runner run
- clusterforge-runner once
- clusterforge-runner version

Config:
runner:
  name: local-runner
  api_url: http://localhost:8080
  token_env: CLUSTERFORGE_RUNNER_TOKEN
  work_dir: .cf/runner-work
  poll_interval: 10s
  max_concurrent_jobs: 1
  allowed_job_types:
    - validate
    - policy_check
    - plan
    - drift_check
    - cost_scan

MVP behavior:
- authenticate as runner
- heartbeat to API
- poll for pending jobs
- claim job
- execute supported job type
- upload result
- cleanup workspace

For MVP execution:
- implement validate and policy_check first
- plan job may be implemented if safe
- apply job should remain disabled until approval workflow exists

Control plane endpoints:
- POST /api/v1/runners/heartbeat
- GET /api/v1/runner/jobs/next
- POST /api/v1/runner/jobs/{id}/claim
- POST /api/v1/runner/jobs/{id}/complete
- POST /api/v1/runner/jobs/{id}/fail

Tests:
- runner config loading
- heartbeat request
- job polling
- job completion
- unsupported job type fails safely

Rules:
- Do not run arbitrary shell commands.
- Do not support apply yet.
- Do not store tokens in logs.
- Clean workspace after job.
- No cloud credentials required for tests.

Run:
- cd runner && gofmt -w .
- cd runner && go test ./...
- cd runner && go build -o clusterforge-runner .
```
