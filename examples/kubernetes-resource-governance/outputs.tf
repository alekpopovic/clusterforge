output "resource_quota_name" {
  description = "ResourceQuota created by the example."
  value       = module.quota.resource_quota_name
}

output "limit_range_name" {
  description = "LimitRange created by the example."
  value       = module.limits.limit_range_name
}
