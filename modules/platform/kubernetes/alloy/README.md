# platform/kubernetes/alloy

## Purpose

Installs Grafana Alloy with Helm. Alloy is the preferred log/telemetry agent
for new ClusterForge observability examples.

## Status

Implemented.

## Usage

```hcl
module "alloy" {
  source = "../../../modules/platform/kubernetes/alloy"

  namespace        = "logging"
  create_namespace = true

  values = [
    yamlencode({
      alloy = {
        configMap = {
          content = "/* provide Alloy config here */"
        }
      }
    })
  ]
}
```

## Notes

This module only installs the Helm release. Production log collection needs an
explicit Alloy configuration that selects sources, processors, labels, and Loki
or other destinations. Do not put credentials in Terraform values; reference
Kubernetes Secrets or an external secret manager instead.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
