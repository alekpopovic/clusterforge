output "endpoint_ids" {
  description = "VPC endpoint IDs keyed by service suffix."
  value       = module.vpc_endpoints.endpoint_ids
}

output "security_group_id" {
  description = "Created interface endpoint security group ID."
  value       = module.vpc_endpoints.security_group_id
}
