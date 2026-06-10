output "vpc_id" {
  description = "Example VPC ID."
  value       = module.network.vpc_id
}

output "public_subnet_ids" {
  description = "Example public subnet IDs."
  value       = module.network.public_subnet_ids
}

output "private_subnet_ids" {
  description = "Example private subnet IDs."
  value       = module.network.private_subnet_ids
}
