output "cluster_name" {
  description = "Example EKS cluster name."
  value       = module.eks.cluster_name
}

output "cluster_endpoint" {
  description = "Example EKS cluster endpoint."
  value       = module.eks.cluster_endpoint
}

output "node_group_names" {
  description = "Example EKS managed node group names."
  value       = module.eks.node_group_names
}

output "oidc_provider_arn" {
  description = "Example EKS IAM OIDC provider ARN for IRSA."
  value       = module.eks.oidc_provider_arn
}
