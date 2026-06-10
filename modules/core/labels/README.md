# core/labels

## Purpose

Produces consistent Kubernetes-compatible labels for ClusterForge platform and
workload modules.

## Status

Implemented. This module is provider-free and creates no infrastructure
resources.

## Tags vs Labels

Use this module for Kubernetes labels. Use `modules/core/tags` for cloud tags,
such as AWS tags.

Kubernetes label values are normalized to lowercase strings, unsupported
characters are replaced with dashes, repeated separators are collapsed, and
values are truncated to 63 characters. Empty optional labels are omitted.

`extra_labels` are merged last. This gives callers explicit override ability
for edge cases while still keeping ClusterForge defaults visible.

## Inputs

| Name | Type | Default | Description |
| --- | --- | --- | --- |
| `project` | `string` | n/a | Project name used for `clusterforge.io/project`. |
| `environment` | `string` | n/a | Environment name used for `clusterforge.io/environment`. |
| `app` | `string` | `""` | Optional `app.kubernetes.io/name` label value. |
| `component` | `string` | `""` | Optional `app.kubernetes.io/component` label value. |
| `part_of` | `string` | `""` | Optional `app.kubernetes.io/part-of` label value. |
| `managed_by` | `string` | `"terraform"` | `app.kubernetes.io/managed-by` label value. |
| `extra_labels` | `map(string)` | `{}` | Additional labels sanitized and merged last. |

## Outputs

| Name | Description |
| --- | --- |
| `labels` | Merged Kubernetes-compatible labels with standard ClusterForge metadata. |

## Usage

```hcl
module "app_labels" {
  source = "path/to/modules/core/labels"

  project     = "clusterforge"
  environment = "dev"
  app         = "api"
  component   = "web"
  part_of     = "customer-portal"

  extra_labels = {
    "team.example.com/name" = "platform"
  }
}
```
