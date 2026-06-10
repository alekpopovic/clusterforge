# platform/kubernetes/ingress-nginx

## Purpose

This module will manage the ClusterForge platform/kubernetes/ingress-nginx component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

ingress-nginx Helm release values and ingress class defaults.

## Usage

```hcl
module "example" {
  source = "path/to/modules/platform/kubernetes/ingress-nginx"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
