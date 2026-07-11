## Prompt 152 — Cross-cluster GitOps support

```text
Design and implement cross-cluster GitOps support.

Goal:
Allow Argo CD or Flux to manage apps across multiple clusters using ClusterForge inventory.

Create docs:
- docs/gitops-multi-cluster.md

Config:
clusterforge.yaml:
  gitops:
    provider: argocd
    repo_url: https://github.com/example/gitops
    clusters:
      - name: dev-eks
        environment: dev
      - name: prod-eks
        environment: prod

CLI:
- cf gitops render
- cf gitops render --cluster dev-eks
- cf gitops clusters
- cf gitops apps

Generated output:
- Argo CD Application manifests
- AppProject manifests
- app-of-apps structure
- cluster-specific overlays if practical

Rules:
- Do not store Git credentials.
- Do not register clusters automatically unless explicitly implemented later.
- Generated manifests should be committed to GitOps repo by user.
- Keep Terraform boundary clear.

Tests:
- Render app-of-apps for two clusters.
- Render cluster-specific app.
- Missing repo_url fails.
- No secrets in output.

Run:
- gofmt
- go test ./...
```


---
