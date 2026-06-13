output "bucket_name" {
  description = "S3 bucket name for Terraform state."
  value       = module.tfstate_backend.bucket_name
}

output "bucket_arn" {
  description = "S3 bucket ARN for Terraform state."
  value       = module.tfstate_backend.bucket_arn
}

output "dynamodb_table_name" {
  description = "DynamoDB table name for Terraform state locking."
  value       = module.tfstate_backend.dynamodb_table_name
}

output "dynamodb_table_arn" {
  description = "DynamoDB table ARN for Terraform state locking."
  value       = module.tfstate_backend.dynamodb_table_arn
}

output "backend_config_example" {
  description = "Example ClusterForge backend configuration values."
  value       = merge(module.tfstate_backend.backend_config_example, { region = var.region })
}
