# platform/kubernetes/loki

## Purpose

This module will manage the ClusterForge platform/kubernetes/loki component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

Loki logging stack installation and retention defaults.

## Usage

```hcl
module "example" {
  source = "path/to/modules/platform/kubernetes/loki"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
