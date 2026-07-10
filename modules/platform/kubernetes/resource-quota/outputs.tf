output "namespace" {
  description = "Namespace containing the ResourceQuota."
  value       = local.namespace
}

output "resource_quota_name" {
  description = "Name of the ResourceQuota."
  value       = kubernetes_resource_quota_v1.this.metadata[0].name
}
