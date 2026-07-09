## Prompt 83 — Existing Kubernetes environment support

```text
Implement first-class existing Kubernetes cluster support.

Goal:
Support users who already have a Kubernetes cluster and only want ClusterForge for platform add-ons and workloads.

CLI:
- cf env create dev --cloud existing --orchestrator kubernetes
- cf generate dev

Generated environment:
live/dev/existing-kubernetes/
  versions.tf
  providers.tf
  main.tf
  variables.tf
  outputs.tf
  terraform.tfvars.example
  README.md

Provider configuration:
- Support kubeconfig path.
- Support kubeconfig context.
- Support in-cluster configuration only as documented future option.
- Helm provider must use the same Kubernetes connection.

Generated main.tf:
- Include optional platform bootstrap.
- Include app rendering support.
- Do not include cloud network or cluster creation modules.

clusterforge.yaml:
- Add support for cloud: existing
- Add support for orchestrator: kubernetes

Docs:
- docs/existing-kubernetes.md

Examples:
- examples/existing-kubernetes-basic-app
- examples/existing-kubernetes-platform-bootstrap

Rules:
- Do not assume EKS/AKS/GKE.
- Do not store kubeconfig content in repo.
- Do not generate secrets.
- Keep this path lightweight.

Tests:
- CLI generator test for existing+kubernetes.
- Config validation test.
- Template rendering test.

Run:
- gofmt
- go test ./...
- terraform fmt -recursive
```

---
