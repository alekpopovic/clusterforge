output "tenant_namespaces" {
  description = "Namespaces created for the example tenant."
  value       = module.payments_tenant.namespaces
}
