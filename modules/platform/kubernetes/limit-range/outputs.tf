output "namespace" {
  description = "Namespace containing the LimitRange."
  value       = local.namespace
}

output "limit_range_name" {
  description = "Name of the LimitRange."
  value       = kubernetes_limit_range_v1.this.metadata[0].name
}
