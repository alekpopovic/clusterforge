# platform/kubernetes/external-dns

Installs external-dns with Helm.

This module assumes Kubernetes and Helm providers are configured in the root
module. Configure DNS provider credentials outside this module.

## Usage

```hcl
module "external_dns" {
  source = "../../../modules/platform/kubernetes/external-dns"

  namespace     = "external-dns"
  chart_version = "1.14.5"
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
