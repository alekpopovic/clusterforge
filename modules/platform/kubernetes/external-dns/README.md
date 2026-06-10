# platform/kubernetes/external-dns

## Purpose

This module will manage the ClusterForge platform/kubernetes/external-dns component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

external-dns installation and DNS provider integration inputs.

## Usage

```hcl
module "example" {
  source = "path/to/modules/platform/kubernetes/external-dns"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
