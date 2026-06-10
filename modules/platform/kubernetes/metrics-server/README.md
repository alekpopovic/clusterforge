# platform/kubernetes/metrics-server

## Purpose

This module will manage the ClusterForge platform/kubernetes/metrics-server component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

metrics-server installation and cluster metrics defaults.

## Usage

```hcl
module "example" {
  source = "path/to/modules/platform/kubernetes/metrics-server"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
