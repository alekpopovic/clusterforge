# orchestrators/ecs/cluster

## Purpose

This module will manage the ClusterForge orchestrators/ecs/cluster component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

ECS clusters, capacity providers, service discovery hooks, and execution defaults.

## Usage

```hcl
module "example" {
  source = "path/to/modules/orchestrators/ecs/cluster"

  name        = "example"
  environment = "dev"
  tags        = {}
}
```
