output "aws_resource_name" {
  description = "Example AWS-style resource name."
  value       = module.aws_resource_name.name
}

output "kubernetes_labels_safe_name" {
  description = "Example Kubernetes-label-safe name."
  value       = module.kubernetes_app_name.labels_safe_name
}

output "platform_dns_safe_name" {
  description = "Example DNS-safe platform component name."
  value       = module.platform_component_name.dns_safe_name
}
