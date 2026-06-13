# platform/kubernetes/loki

## Purpose

Installs Loki with Helm for Kubernetes log storage and querying.

## Status

Implemented.

## Usage

```hcl
module "loki" {
  source = "../../../modules/platform/kubernetes/loki"

  namespace        = "logging"
  create_namespace = true

  storage_enabled = false
}
```

Enable persistent storage only after selecting a StorageClass and retention
strategy:

```hcl
module "loki" {
  source = "../../../modules/platform/kubernetes/loki"

  storage_enabled    = true
  storage_class_name = "gp3"
}
```

## Notes

This module assumes Kubernetes and Helm providers are configured in the root
module. Persistent storage is disabled by default. Production Loki deployments
usually need tuned retention, object storage, resource limits, and chart values
specific to the cluster.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
