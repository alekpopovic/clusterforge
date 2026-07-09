output "role_arn" {
  description = "IAM role ARN for cert-manager service account."
  value       = module.irsa.role_arn
}

output "role_name" {
  description = "IAM role name for cert-manager service account."
  value       = module.irsa.role_name
}

output "policy_json" {
  description = "Rendered Route53 DNS01 IAM policy JSON."
  value       = data.aws_iam_policy_document.route53_dns01.json
}
