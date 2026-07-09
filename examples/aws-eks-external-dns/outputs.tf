output "external_dns_role_arn" {
  description = "ExternalDNS IRSA role ARN."
  value       = module.external_dns_irsa.role_arn
}

output "external_dns_namespace" {
  description = "Namespace where ExternalDNS is installed."
  value       = module.external_dns.namespace
}
