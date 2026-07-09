output "postgres_endpoint" {
  description = "RDS PostgreSQL endpoint."
  value       = module.postgres.endpoint
}

output "postgres_secret_arn" {
  description = "AWS-managed master password secret ARN."
  value       = module.postgres.secret_arn
}
