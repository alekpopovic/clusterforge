# workloads/ecs/service

## Purpose

This module will manage the ClusterForge workloads/ecs/service component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

ECS task definition, service, networking, load balancer attachment, and autoscaling hooks.

## Usage

```hcl
module "example" {
  source = "path/to/modules/workloads/ecs/service"

  name        = "example"
  environment = "dev"
  tags        = {}
}
```
