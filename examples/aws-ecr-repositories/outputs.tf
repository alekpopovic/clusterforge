output "repository_urls" {
  description = "ECR repository URLs keyed by repository name."
  value       = module.ecr.repository_urls
}

output "repository_names" {
  description = "ECR repository names."
  value       = module.ecr.repository_names
}
