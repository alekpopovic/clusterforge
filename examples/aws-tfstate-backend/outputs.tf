output "bucket_name" {
  description = "S3 bucket name for Terraform state."
  value       = aws_s3_bucket.state.bucket
}

output "dynamodb_table_name" {
  description = "DynamoDB table name for Terraform state locking."
  value       = aws_dynamodb_table.locks.name
}

output "backend_config_example" {
  description = "Example ClusterForge backend configuration values."
  value = {
    type           = "s3"
    bucket         = aws_s3_bucket.state.bucket
    region         = var.region
    dynamodb_table = aws_dynamodb_table.locks.name
  }
}
