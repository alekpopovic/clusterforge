# Prompt 187 — Runner deployment on Kubernetes

```text
Create Kubernetes deployment support for ClusterForge runner.

Create:
- deploy/runner/kubernetes/
  namespace.yaml or Terraform/Helm template
  deployment.yaml
  serviceaccount.yaml
  configmap.yaml
  secret.example.yaml

Optional Helm chart:
- charts/clusterforge-runner/

Runner deployment should support:
- API URL
- runner token from Kubernetes Secret
- work dir emptyDir
- resource requests/limits
- allowed job types
- concurrency
- node selector/tolerations optional

Security:
- run as non-root
- read-only root filesystem where practical
- no privileged container
- no cloud credentials baked into image
- credentials should be provided through workload identity or external secret

Docs:
- docs/runner-kubernetes.md

Example:
- examples/control-plane-runner-kubernetes

Rules:
- Do not include real token.
- Do not grant cluster-admin.
- Do not enable apply by default.
```
