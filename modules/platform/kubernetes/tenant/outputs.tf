output "name" {
  description = "Logical tenant name."
  value       = local.name
}

output "namespaces" {
  description = "Tenant namespaces managed by this module."
  value       = sort(tolist(local.namespaces))
}

output "namespace_labels" {
  description = "Labels, including Pod Security Admission levels, applied to tenant namespaces."
  value       = local.namespace_labels
}

output "role_names" {
  description = "Namespace Role names keyed by namespace when RBAC is enabled."
  value       = { for namespace, role in kubernetes_role_v1.this : namespace => role.metadata[0].name }
}
