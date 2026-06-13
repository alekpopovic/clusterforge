# platform/kubernetes/prometheus-stack

## Purpose

Installs kube-prometheus-stack with Helm for Kubernetes metrics, alerting, and
Grafana dashboards.

## Status

Implemented.

## Usage

```hcl
module "prometheus_stack" {
  source = "../../../modules/platform/kubernetes/prometheus-stack"

  namespace        = "monitoring"
  create_namespace = true

  storage_enabled = false
}
```

Enable Grafana ingress only after reviewing authentication, TLS, and network
exposure:

```hcl
module "prometheus_stack" {
  source = "../../../modules/platform/kubernetes/prometheus-stack"

  enable_grafana_ingress = true
  grafana_host           = "grafana.example.com"
}
```

## Secret Handling

Do not hardcode Grafana admin passwords in Terraform. Use chart values that
reference an existing Kubernetes Secret, External Secrets Operator, or another
approved secret path. Terraform state can retain sensitive values passed through
Helm values.

## Notes

This module assumes Kubernetes and Helm providers are configured in the root
module. Review CRDs, retention, storage, alerting, and chart upgrade notes
before production use. Persistent storage is disabled by default.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
