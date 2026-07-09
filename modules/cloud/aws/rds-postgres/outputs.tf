output "endpoint" {
  description = "RDS PostgreSQL endpoint."
  value       = aws_db_instance.this.address
}

output "port" {
  description = "RDS PostgreSQL port."
  value       = aws_db_instance.this.port
}

output "db_name" {
  description = "Database name."
  value       = aws_db_instance.this.db_name
}

output "security_group_id" {
  description = "Security group ID attached to the database."
  value       = aws_security_group.this.id
}

output "secret_arn" {
  description = "AWS Secrets Manager secret ARN for the master user password when AWS-managed password is enabled."
  value       = var.manage_master_user_password ? aws_db_instance.this.master_user_secret[0].secret_arn : null
}
