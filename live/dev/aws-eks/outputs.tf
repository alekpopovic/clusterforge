output "cluster_name" {
  description = "Development EKS cluster name."
  value       = module.eks.cluster_name
}

output "vpc_id" {
  description = "Development VPC ID."
  value       = module.network.vpc_id
}
