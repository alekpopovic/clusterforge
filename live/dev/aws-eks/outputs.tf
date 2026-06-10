output "vpc_id" {
  description = "Development VPC ID."
  value       = module.network.vpc_id
}

output "public_subnet_ids" {
  description = "Development public subnet IDs."
  value       = module.network.public_subnet_ids
}

output "private_subnet_ids" {
  description = "Development private subnet IDs."
  value       = module.network.private_subnet_ids
}

output "cluster_name" {
  description = "Development EKS cluster name."
  value       = module.eks.cluster_name
}

output "cluster_endpoint" {
  description = "Development EKS cluster endpoint."
  value       = module.eks.cluster_endpoint
}

output "cluster_oidc_issuer_url" {
  description = "Development EKS cluster OIDC issuer URL."
  value       = module.eks.cluster_oidc_issuer_url
}

output "node_group_names" {
  description = "Development EKS managed node group names."
  value       = module.eks.node_group_names
}
