# workloads/kubernetes/rollout-app

Deploys an opt-in Argo Rollout with stable/active and canary/preview Services.
It does not replace `workloads/kubernetes/app` or create the controller.

## Status

Implemented.

```hcl
module "api_rollout" {
  source    = "../../../modules/workloads/kubernetes/rollout-app"
  name      = "api"
  namespace = "apps"
  image     = "ghcr.io/example/api:1.2.3"
  strategy = {
    type = "canary"
    canary_steps = [
      { set_weight = 20 },
      { pause = true },
      { set_weight = 100 }
    ]
  }
}
```

The Argo Rollouts CRD must exist before Terraform plans this module. Basic
canary weights without a traffic router approximate traffic by pod counts. For
precise traffic shifting, configure a supported ingress or service mesh and
extend the strategy deliberately. Provider configuration belongs in the root.

## Generated Terraform documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
