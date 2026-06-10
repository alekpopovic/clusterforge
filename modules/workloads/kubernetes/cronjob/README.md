# workloads/kubernetes/cronjob

## Purpose

This module will manage the ClusterForge workloads/kubernetes/cronjob component.

## Status

Placeholder. This module currently creates no resources.

## Expected Future Resources

Kubernetes CronJob, schedules, retry settings, and pod defaults.

## Usage

```hcl
module "example" {
  source = "path/to/modules/workloads/kubernetes/cronjob"

  name        = "example"
  environment = "dev"
  labels      = {}
}
```
