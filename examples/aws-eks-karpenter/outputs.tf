output "cluster_name" {
  description = "EKS cluster name."
  value       = module.eks.cluster_name
}

output "karpenter_role_arn" {
  description = "IAM role ARN used by the Karpenter controller."
  value       = module.karpenter_irsa.role_arn
}

output "karpenter_release_name" {
  description = "Karpenter Helm release name."
  value       = module.karpenter.release_name
}
