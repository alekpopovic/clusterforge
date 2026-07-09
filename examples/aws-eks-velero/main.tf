module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "prod"
  component   = "velero"
}

module "velero_bucket" {
  source = "../../modules/cloud/aws/velero-backup-bucket"

  name        = "clusterforge-prod-velero"
  environment = "prod"
  bucket_name = "example-clusterforge-prod-velero-backups"
  tags        = module.tags.tags
}

module "velero" {
  source = "../../modules/platform/kubernetes/velero"

  bucket                   = module.velero_bucket.bucket_name
  service_account_role_arn = var.velero_service_account_role_arn

  labels = {
    "clusterforge.io/managed-by" = "terraform"
  }
}
