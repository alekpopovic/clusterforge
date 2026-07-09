output "cluster_name" {
  description = "Example EKS cluster name."
  value       = module.eks.cluster_name
}

output "cluster_endpoint" {
  description = "Example EKS cluster endpoint."
  value       = module.eks.cluster_endpoint
}

output "cluster_encryption_key_arn" {
  description = "KMS key ARN used for EKS secrets encryption."
  value       = module.eks.cluster_encryption_key_arn
}

output "control_plane_log_group_name" {
  description = "CloudWatch log group name for EKS control plane logs."
  value       = module.eks.control_plane_log_group_name
}
