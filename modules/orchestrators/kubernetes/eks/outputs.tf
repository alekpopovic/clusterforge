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
