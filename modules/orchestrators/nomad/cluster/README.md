# orchestrators/nomad/cluster

## Purpose

This module will manage the ClusterForge orchestrators/nomad/cluster component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

Nomad server/client cluster primitives and scheduling metadata.

## Usage

```hcl
module "example" {
  source = "path/to/modules/orchestrators/nomad/cluster"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
