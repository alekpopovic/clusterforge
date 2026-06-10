output "name" {
  description = "Generated resource name."
  value       = local.short_name
}

output "dns_label" {
  description = "DNS-compatible generated label."
  value       = local.dns_label
}

output "path_prefix" {
  description = "Slash-separated prefix useful for parameter or secret paths."
  value       = local.path_prefix
}

output "parts" {
  description = "Name parts used to build the generated name."
  value       = local.parts
}
