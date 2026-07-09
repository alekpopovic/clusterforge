# platform/kubernetes/metrics-server

Installs metrics-server with Helm.

This module assumes Kubernetes and Helm providers are configured in the root
module.

## Usage

```hcl
module "metrics_server" {
  source = "../../../modules/platform/kubernetes/metrics-server"

  namespace     = "metrics-server"
  chart_version = "3.12.1"
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
