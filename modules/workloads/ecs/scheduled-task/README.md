# workloads/ecs/scheduled-task

## Purpose

This module will manage the ClusterForge workloads/ecs/scheduled-task component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

EventBridge schedule, ECS task definition, and execution roles.

## Usage

```hcl
module "example" {
  source = "path/to/modules/workloads/ecs/scheduled-task"

  name        = "example"
  environment = "dev"
  tags        = {}
}
```
