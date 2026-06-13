output "namespaces" {
  description = "Namespaces managed by this module."
  value       = keys(var.namespaces)
}

output "labels" {
  description = "Pod Security labels applied by namespace."
  value       = local.pod_security_labels
}
