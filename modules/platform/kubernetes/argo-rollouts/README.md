# platform/kubernetes/argo-rollouts

Installs the opt-in Argo Rollouts controller and CRDs from the official Argo Helm
repository. Standard Kubernetes Deployments remain unaffected.

## Status

Implemented.

```hcl
module "argo_rollouts" {
  source = "../../../modules/platform/kubernetes/argo-rollouts"
  chart_version = "<reviewed-version>"
}
```

Dashboard ingress is disabled by default. Prefer the local kubectl plugin
dashboard unless an authenticated ingress design has been reviewed. Provider
configuration belongs in the root module.

## Generated Terraform documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
