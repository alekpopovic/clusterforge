output "module_name" {
  description = "ClusterForge module path."
  value       = local.module_name
}

output "enabled" {
  description = "Whether this module placeholder is enabled."
  value       = var.enabled
}
