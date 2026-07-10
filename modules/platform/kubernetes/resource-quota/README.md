# platform/kubernetes/resource-quota

Creates one explicit ResourceQuota in an existing namespace. It does not create
namespaces or apply quotas globally.

## Status

Implemented.

## Usage

```hcl
module "app_quota" {
  source = "../../../modules/platform/kubernetes/resource-quota"

  namespace = "apps"
  hard = {
    "requests.cpu"    = "4"
    "requests.memory" = "8Gi"
    "limits.cpu"      = "8"
    "limits.memory"   = "16Gi"
    pods              = "40"
  }
}
```

Introducing a quota can reject workloads that omit requests or exceed remaining
namespace capacity. Measure current usage and review rollout order first.

Provider configuration belongs in the calling root module.

## Generated Terraform documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
