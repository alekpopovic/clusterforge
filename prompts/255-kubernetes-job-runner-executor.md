# Prompt 255 — Kubernetes Job runner executor

```text
Implement Kubernetes Job executor for ClusterForge runner.

Goal:
Allow runner jobs to execute as isolated Kubernetes Jobs instead of long-running process execution.

Architecture:
- runner controller receives job from Control Plane
- creates Kubernetes Job per ClusterForge job
- Job runs a runner-worker container
- worker executes validate/plan/policy/drift/cost job
- worker uploads result/artifacts
- Kubernetes Job is deleted or retained based on policy

Config:
runner:
  executor: kubernetes_job
  kubernetes_job:
    namespace: clusterforge-runners
    service_account_name: clusterforge-runner-worker
    image: ghcr.io/example/clusterforge-runner-worker:latest
    cleanup_finished_jobs: true
    ttl_seconds_after_finished: 3600
    resources:
      requests:
        cpu: 250m
        memory: 512Mi
      limits:
        cpu: 1000m
        memory: 2Gi

Features:
- per-job workspace emptyDir
- config mounted from Secret/ConfigMap
- token from Secret
- resource limits
- node selector/tolerations
- labels for tenant/job/environment
- no privileged containers
- run as non-root

Tests:
- job manifest rendering
- labels included
- secrets referenced by name only
- unsupported apply blocked unless enabled
- cleanup policy

Docs:
- docs/control-plane/kubernetes-job-executor.md

Rules:
- Do not include real tokens.
- Do not grant cluster-admin.
- Apply disabled by default.
```
