output "cluster_name" {
  description = "EKS cluster name."
  value       = aws_eks_cluster.this.name
}

output "cluster_arn" {
  description = "EKS cluster ARN."
  value       = aws_eks_cluster.this.arn
}

output "cluster_endpoint" {
  description = "EKS Kubernetes API endpoint."
  value       = aws_eks_cluster.this.endpoint
}

output "cluster_certificate_authority_data" {
  description = "Base64-encoded EKS cluster certificate authority data."
  value       = aws_eks_cluster.this.certificate_authority[0].data
  sensitive   = true
}

output "cluster_oidc_issuer_url" {
  description = "EKS cluster OIDC issuer URL."
  value       = aws_eks_cluster.this.identity[0].oidc[0].issuer
}

output "oidc_provider_arn" {
  description = "IAM OIDC provider ARN when IRSA is enabled."
  value       = var.enable_irsa ? aws_iam_openid_connect_provider.this[0].arn : null
}

output "oidc_provider_url" {
  description = "EKS OIDC provider issuer URL when IRSA is enabled."
  value       = local.oidc_provider_url
}

output "oidc_issuer_hostpath" {
  description = "EKS OIDC issuer URL without the https:// prefix, suitable for IAM condition keys."
  value       = local.oidc_issuer_hostpath
}

output "node_group_names" {
  description = "Managed node group names."
  value       = [for key in sort(keys(aws_eks_node_group.this)) : aws_eks_node_group.this[key].node_group_name]
}

output "node_group_arns" {
  description = "Managed node group ARNs."
  value       = { for key, node_group in aws_eks_node_group.this : key => node_group.arn }
}

output "cluster_security_group_id" {
  description = "EKS cluster security group ID created by EKS."
  value       = aws_eks_cluster.this.vpc_config[0].cluster_security_group_id
}

output "cluster_encryption_key_arn" {
  description = "KMS key ARN used for EKS secrets encryption, when enabled."
  value       = local.cluster_encryption_key_arn
}

output "control_plane_log_group_name" {
  description = "CloudWatch log group name for EKS control plane logs, when enabled."
  value       = local.create_control_plane_log_group ? aws_cloudwatch_log_group.control_plane[0].name : null
}
