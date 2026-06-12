output "role_arn" {
  description = "ARN of the Karpenter controller IAM role."
  value       = module.role.role_arn
}

output "role_name" {
  description = "Name of the Karpenter controller IAM role."
  value       = module.role.role_name
}
