# core/tags

## Purpose

This module will manage the ClusterForge core/tags component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

Tag normalization and merge helpers.

## Usage

```hcl
module "example" {
  source = "path/to/modules/core/tags"

  name        = "example"
  environment = "dev"
  tags        = {}
}
```
