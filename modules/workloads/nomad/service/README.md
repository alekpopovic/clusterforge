# workloads/nomad/service

## Purpose

This module will manage the ClusterForge workloads/nomad/service component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

Nomad service job specification and service discovery metadata.

## Usage

```hcl
module "example" {
  source = "path/to/modules/workloads/nomad/service"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
