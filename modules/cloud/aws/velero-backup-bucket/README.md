# cloud/aws/velero-backup-bucket

## Purpose

Creates an S3 bucket for Velero backups with versioning, encryption, public
access blocking, and optional lifecycle expiration. The bucket is intentionally
not force-destroyed by default.

## Status

Implemented.

## Usage

```hcl
module "velero_bucket" {
  source = "../../../modules/cloud/aws/velero-backup-bucket"

  name        = "clusterforge-prod-velero"
  environment = "prod"
  bucket_name = "example-clusterforge-prod-velero-backups"
}
```

Review retention and recovery requirements before enabling lifecycle
expiration. Backup bucket deletion can remove the only practical restore path
for a failed cluster.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
