# Multi-cluster GitOps

ClusterForge can render Argo CD inventory manifests from local configuration. It
does not register clusters, contact Argo CD or Git, commit files, or handle Git
credentials.

```yaml
gitops:
  provider: argocd
  repo_url: https://github.com/example/gitops
  clusters:
    - name: dev-eks
      environment: dev
    - name: prod-eks
      environment: prod
```

```bash
cf gitops clusters
cf gitops apps
cf gitops render > clusterforge-apps.yaml
cf gitops render --cluster dev-eks > dev-eks-apps.yaml
```

The renderer creates an `AppProject`, one app-of-apps `Application` per cluster,
and cluster-specific application entries using paths such as
`apps/api/overlays/prod`. Destination `name` must already be registered in Argo
CD by an approved external process. Review and commit generated YAML to the
GitOps repository yourself.

Terraform owns cluster infrastructure and the initial GitOps controller
bootstrap. Argo CD owns the application manifests after handoff. Never let both
systems manage the same Kubernetes object or Helm release. Repository credentials
belong in Argo CD's approved secret integration and must not appear in
`clusterforge.yaml` or generated manifests. Flux remains a future provider; the
CLI fails explicitly rather than emitting incompatible resources.
