# core/naming

Builds consistent names from `project`, `environment`, and `name`.

This module creates no infrastructure and has no provider configuration.

## Example

```hcl
module "name" {
  source = "../../../modules/core/naming"

  project     = "clusterforge"
  environment = "dev"
  name        = "eks"
}
```
