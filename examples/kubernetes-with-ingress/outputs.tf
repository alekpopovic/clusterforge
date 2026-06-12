output "issuer_name" {
  description = "Name of the example cert-manager ClusterIssuer."
  value       = module.cert_manager_issuer.name
}

output "issuer_kind" {
  description = "Kind of the example cert-manager issuer resource."
  value       = module.cert_manager_issuer.kind
}
