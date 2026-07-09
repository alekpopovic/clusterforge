# platform/kubernetes/velero

## Purpose

Installs Velero with Helm for Kubernetes backup and restore workflows. The
first implementation supports AWS S3 backup storage. This module does not
create the bucket and does not store cloud credentials in Terraform.

## Status

Implemented.

## Usage

```hcl
module "velero" {
  source = "../../../modules/platform/kubernetes/velero"

  bucket                   = module.velero_bucket.bucket_name
  service_account_role_arn = module.velero_irsa.role_arn
  aws_plugin_image         = "velero/velero-plugin-for-aws:v1.10.1"
}
```

When `service_account_role_arn` is provided, the Velero service account is
annotated for EKS IRSA. Do not pass static cloud credentials through Helm
values.

The Velero provider input is named `velero_provider` because `provider` is a
reserved Terraform module argument name.

## Restore Testing

Installing Velero is not proof that backups are recoverable. Create a test
namespace, back it up, delete it, restore it, and validate the restored
workloads before relying on the configuration.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
