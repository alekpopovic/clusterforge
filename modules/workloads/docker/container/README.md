# workloads/docker/container

## Purpose

This module will manage the ClusterForge workloads/docker/container component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

Docker container configuration and host-level runtime options.

## Usage

```hcl
module "example" {
  source = "path/to/modules/workloads/docker/container"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
