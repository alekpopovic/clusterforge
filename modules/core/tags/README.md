# core/tags

## Purpose

Produces consistent cloud resource tags for ClusterForge modules and live
environment roots.

## Status

Implemented. This module is provider-free and creates no infrastructure
resources.

## Tags vs Labels

Use this module for cloud tags, such as AWS tags. Use `modules/core/labels`
for Kubernetes labels, which have stricter key and value rules.

`extra_tags` are merged last. This intentionally allows a live environment to
override standard tags when a cloud account, billing policy, or migration needs
an exception.

Empty optional values are omitted so the output map does not contain null or
blank tag values.

## Inputs

| Name | Type | Default | Description |
| --- | --- | --- | --- |
| `project` | `string` | n/a | Project name used for the `Project` tag. |
| `environment` | `string` | n/a | Environment name used for the `Environment` tag. |
| `owner` | `string` | `""` | Optional `Owner` tag value. |
| `cost_center` | `string` | `""` | Optional `CostCenter` tag value. |
| `managed_by` | `string` | `"terraform"` | Tool or workflow responsible for resources. |
| `component` | `string` | `""` | Optional `Component` tag value. |
| `extra_tags` | `map(string)` | `{}` | Additional tags merged last. |

## Outputs

| Name | Description |
| --- | --- |
| `tags` | Merged cloud tags with standard ClusterForge metadata. |

## Usage

```hcl
module "aws_tags" {
  source = "path/to/modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "network"
  owner       = "platform-team"
  cost_center = "cc-1234"

  extra_tags = {
    Compliance = "internal"
  }
}
```
