# kubernetes-helm-app

Example root module for `modules/workloads/kubernetes/helm-app`.

This example installs the public `podinfo` Helm chart into an existing
Kubernetes cluster. It assumes the Kubernetes and Helm providers can use a local
kubeconfig.

## Usage

```bash
terraform init
terraform validate
terraform plan
```

The chart version is pinned in the example for repeatability. Do not put secret
values in `values`, `set`, or `set_sensitive`; reference existing Kubernetes
Secrets or an external secret manager where possible.

Use Argo CD instead when app deployments should be driven by GitOps rather than
Terraform plan/apply.
