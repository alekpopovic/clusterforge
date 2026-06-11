# core/naming

## Purpose

Generates deterministic names for cloud resources, Kubernetes resources,
platform components, and workloads.

## Status

Implemented. This module is provider-free and creates no infrastructure
resources.

## Behavior

The module builds a base name from `project`, `environment`, `component`,
`name`, `extra_parts`, and `suffix`. Empty parts are removed, the remaining
parts are joined with `separator`, unsupported characters are replaced,
duplicated separators are collapsed, leading and trailing separators are
trimmed, and the final `name` output is truncated to `max_length`.

`labels_safe_name` and `dns_safe_name` are always lowercase, dash-separated,
and limited to 63 characters.

## Inputs

| Name | Type | Default | Description |
| --- | --- | --- | --- |
| `project` | `string` | n/a | Project identifier used as the first name segment. |
| `environment` | `string` | n/a | Environment identifier used as the second name segment. |
| `component` | `string` | n/a | Component or layer identifier. |
| `name` | `string` | n/a | Specific resource, platform component, or workload name. |
| `separator` | `string` | `"-"` | Separator used when joining name parts. Must be `"-"`, `"_"`, or `""`. |
| `max_length` | `number` | `63` | Maximum length for the generated `name`. |
| `lowercase` | `bool` | `true` | Whether to lowercase the generated `name` and `full_name`. |
| `extra_parts` | `list(string)` | `[]` | Additional parts appended after `name` and before `suffix`. |
| `suffix` | `string` | `""` | Optional suffix appended after `extra_parts`. |

## Outputs

| Name | Description |
| --- | --- |
| `name` | Normalized name truncated to `max_length`. |
| `full_name` | Normalized full name before `max_length` truncation. |
| `parts` | Non-empty name parts used to build generated names. |
| `labels_safe_name` | Kubernetes-label-friendly lowercase name, limited to 63 characters. |
| `dns_safe_name` | DNS-friendly lowercase name without underscores, limited to 63 characters. |

## Usage

### AWS resource name

```hcl
module "vpc_name" {
  source = "../../modules/core/naming"

  project     = "clusterforge"
  environment = "dev"
  component   = "network"
  name        = "vpc"
}
```

### Kubernetes app name

```hcl
module "app_name" {
  source = "../../modules/core/naming"

  project     = "clusterforge"
  environment = "staging"
  component   = "app"
  name        = "api"
  extra_parts = ["blue"]
}
```

### Platform component name

```hcl
module "platform_name" {
  source = "../../modules/core/naming"

  project     = "clusterforge"
  environment = "prod"
  component   = "platform"
  name        = "ingress-nginx"
  suffix      = "controller"
  max_length  = 48
}
```

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
