output "bucket_name" {
  description = "S3 bucket name for Terraform state."
  value       = aws_s3_bucket.state.bucket
}

output "bucket_arn" {
  description = "S3 bucket ARN for Terraform state."
  value       = aws_s3_bucket.state.arn
}

output "dynamodb_table_name" {
  description = "DynamoDB table name for Terraform state locking."
  value       = aws_dynamodb_table.locks.name
}

output "dynamodb_table_arn" {
  description = "DynamoDB table ARN for Terraform state locking."
  value       = aws_dynamodb_table.locks.arn
}

output "backend_config_example" {
  description = "Example ClusterForge backend configuration values for this backend."
  value = {
    type           = "s3"
    bucket         = aws_s3_bucket.state.bucket
    dynamodb_table = aws_dynamodb_table.locks.name
  }
}

output "kms_key_arn" {
  description = "KMS key ARN used for S3 state bucket encryption, when configured."
  value       = var.kms_key_arn == "" ? null : var.kms_key_arn
}
