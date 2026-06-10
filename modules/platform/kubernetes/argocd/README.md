# platform/kubernetes/argocd

## Purpose

This module will manage the ClusterForge platform/kubernetes/argocd component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

Argo CD installation, projects, repositories, and application bootstrap.

## Usage

```hcl
module "example" {
  source = "path/to/modules/platform/kubernetes/argocd"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
