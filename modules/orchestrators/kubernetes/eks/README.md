# orchestrators/kubernetes/eks

## Purpose

This module will manage the ClusterForge orchestrators/kubernetes/eks component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

EKS cluster, node groups, IAM integration, control plane logging, and access entries.

## Usage

```hcl
module "example" {
  source = "path/to/modules/orchestrators/kubernetes/eks"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
