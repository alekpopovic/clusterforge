output "name" {
  description = "Normalized name truncated to max_length."
  value       = local.name
}

output "full_name" {
  description = "Normalized full name before max_length truncation."
  value       = local.full_name
}

output "parts" {
  description = "Non-empty name parts used to build the generated names."
  value       = local.parts
}

output "labels_safe_name" {
  description = "Kubernetes-label-friendly lowercase name, limited to 63 characters."
  value       = local.labels_safe_name
}

output "dns_safe_name" {
  description = "DNS-friendly lowercase name without underscores, limited to 63 characters."
  value       = local.dns_safe_name
}
