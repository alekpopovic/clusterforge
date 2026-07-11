# OpenTelemetry Collector

Installs the upstream collector Helm chart. Disabled unless explicitly composed.

```hcl
module "otel" {
  source = "../../modules/platform/kubernetes/opentelemetry-collector"
  chart_version = "<pin-reviewed-version>"
  mode = "deployment"
}
```

Provider configuration remains in the root. API keys must come from an external
secret mechanism and must not be embedded in `values` or Terraform state.
