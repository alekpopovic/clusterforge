output "aws_tags" {
  description = "Example AWS/cloud resource tags."
  value       = module.aws_tags.tags
}

output "kubernetes_labels" {
  description = "Example Kubernetes-compatible labels."
  value       = module.kubernetes_labels.labels
}
