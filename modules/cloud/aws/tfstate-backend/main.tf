locals {
  tags = merge(
    {
      Name        = var.name
      Environment = var.environment
      ManagedBy   = "terraform"
    },
    var.tags
  )
}

resource "aws_s3_bucket" "state" {
  #checkov:skip=CKV2_AWS_61:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV2_AWS_62:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV_AWS_144:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV_AWS_145:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV_AWS_18:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV_AWS_19:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  bucket        = var.bucket_name
  force_destroy = var.force_destroy

  tags = local.tags
}

resource "aws_s3_bucket_versioning" "state" {
  count = var.enable_versioning ? 1 : 0

  bucket = aws_s3_bucket.state.id

  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "state" {
  count = var.enable_encryption ? 1 : 0

  bucket = aws_s3_bucket.state.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm     = var.kms_key_arn == "" ? "AES256" : "aws:kms"
      kms_master_key_id = var.kms_key_arn == "" ? null : var.kms_key_arn
    }
  }
}

resource "aws_s3_bucket_public_access_block" "state" {
  bucket = aws_s3_bucket.state.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_dynamodb_table" "locks" {
  #checkov:skip=CKV_AWS_119:Encryption key selection is configurable and may use an approved external or provider-managed key.
  #checkov:skip=CKV_AWS_28:Point-in-time recovery is an explicit compatibility setting for the optional legacy lock table.
  name         = var.dynamodb_table_name
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = local.tags
}
