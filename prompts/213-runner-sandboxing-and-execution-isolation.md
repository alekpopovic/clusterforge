# Prompt 213 — Runner sandboxing and execution isolation

```text
Harden runner execution isolation.

Goal:
Reduce risk when runners execute Terraform and ClusterForge jobs.

Runner config:
execution:
  workspace_root: /var/lib/clusterforge-runner/work
  cleanup_after_job: true
  max_workspace_size_mb: 2048
  allowed_binaries:
    - terraform
    - tofu
    - cf
    - git
  network_policy: documented
  shell_enabled: false

Features:
- unique workspace per job
- path traversal prevention
- workspace cleanup
- max log size
- max artifact size
- command allowlist
- environment variable allowlist/blocklist
- secret redaction
- job timeout
- no arbitrary shell execution

Optional:
- containerized execution design doc
- Kubernetes job-per-plan design doc

Tests:
- workspace path isolation
- cleanup after success/failure
- blocked unknown binary
- blocked path traversal
- log truncation
- timeout

Docs:
- docs/control-plane/runner-security.md

Rules:
- Do not run arbitrary user-provided commands.
- Do not execute plugins unless explicitly allowed.
- Do not log environment secrets.
```
