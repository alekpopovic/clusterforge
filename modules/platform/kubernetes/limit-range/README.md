# platform/kubernetes/limit-range

Creates one explicit LimitRange in an existing namespace. It does not create
namespaces or apply defaults globally.

## Status

Implemented.

## Usage

```hcl
module "app_limits" {
  source = "../../../modules/platform/kubernetes/limit-range"

  namespace = "apps"
  limits = [{
    type = "Container"
    default_request = {
      cpu    = "100m"
      memory = "128Mi"
    }
    default = {
      cpu    = "500m"
      memory = "512Mi"
    }
  }]
}
```

Defaults are applied by Kubernetes admission and can change workload scheduling
or HPA behavior. Review existing manifests before rollout.

Provider configuration belongs in the calling root module.

## Generated Terraform documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
