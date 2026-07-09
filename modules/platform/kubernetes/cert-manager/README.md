# platform/kubernetes/cert-manager

Installs cert-manager with Helm.

This module assumes Kubernetes and Helm providers are configured in the root
module. Review cert-manager CRD lifecycle before upgrades.

## Usage

```hcl
module "cert_manager" {
  source = "../../../modules/platform/kubernetes/cert-manager"

  namespace     = "cert-manager"
  chart_version = "v1.14.5"
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
