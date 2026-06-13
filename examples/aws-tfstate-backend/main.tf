module "tfstate_backend" {
  source = "../../modules/cloud/aws/tfstate-backend"

  name                = var.name
  environment         = var.environment
  bucket_name         = var.bucket_name
  dynamodb_table_name = var.dynamodb_table_name
  force_destroy       = var.force_destroy
  enable_versioning   = var.enable_versioning
  enable_encryption   = var.enable_encryption
  tags                = var.tags
}
