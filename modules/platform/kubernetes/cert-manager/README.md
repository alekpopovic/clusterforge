# platform/kubernetes/cert-manager

## Purpose

This module will manage the ClusterForge platform/kubernetes/cert-manager component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

cert-manager installation, issuers, and certificate defaults.

## Usage

```hcl
module "example" {
  source = "path/to/modules/platform/kubernetes/cert-manager"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
