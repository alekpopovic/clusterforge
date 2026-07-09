output "role_arn" {
  description = "IAM role ARN for ExternalDNS service account."
  value       = module.irsa.role_arn
}

output "role_name" {
  description = "IAM role name for ExternalDNS service account."
  value       = module.irsa.role_name
}

output "policy_json" {
  description = "Rendered ExternalDNS Route53 IAM policy JSON."
  value       = data.aws_iam_policy_document.external_dns.json
}
