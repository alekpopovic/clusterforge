## Prompt 82 — Local Kubernetes development target with Kind or K3d

```text
Add local Kubernetes development target support.

Goal:
Allow users to test ClusterForge workloads locally without cloud infrastructure.

Supported local targets:
- kind
- k3d

Create docs:
- docs/local-development.md

CLI support:
- cf local create kind
- cf local create k3d
- cf local delete kind
- cf local delete k3d
- cf local kubeconfig kind
- cf local status

Behavior:
- Use local binaries kind or k3d if installed.
- Do not vendor kind/k3d.
- Print clear error if binary is missing.
- Generate local ClusterForge environment:
  live/local/kind
  live/local/k3d

For local Kubernetes:
- Configure Kubernetes provider from current kubeconfig context.
- Allow deployment of workloads/kubernetes/app.
- Allow deployment of platform modules only when appropriate.
- Do not install cloud-specific modules.

Add examples:
- examples/local-kind-app
- examples/local-k3d-app

Tests:
- Unit test command construction.
- Do not require kind/k3d in CI.
- Add skip behavior when binaries are missing.

Docs must explain:
- prerequisites
- create cluster
- generate env
- deploy app
- cleanup
- limitations compared to EKS/AKS/GKE

Rules:
- No cloud credentials.
- No destructive host operations beyond deleting the named local cluster.
- Do not modify global kubeconfig without clear user action.

Run:
- gofmt
- go test ./...
```

---
