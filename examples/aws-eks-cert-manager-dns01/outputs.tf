output "cert_manager_role_arn" {
  description = "cert-manager Route53 DNS01 IRSA role ARN."
  value       = module.cert_manager_route53_irsa.role_arn
}

output "issuer_name" {
  description = "Created issuer name."
  value       = module.issuer.name
}
