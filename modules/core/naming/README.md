# core/naming

## Purpose

This module will manage the ClusterForge core/naming component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

Name normalization, prefixes, and resource-safe identifiers.

## Usage

```hcl
module "example" {
  source = "path/to/modules/core/naming"

  name        = "example"
  environment = "dev"
}
```
