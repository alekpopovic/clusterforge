output "redis_primary_endpoint" {
  description = "Primary Redis endpoint address."
  value       = module.redis.primary_endpoint_address
}

output "redis_security_group_id" {
  description = "Redis security group ID."
  value       = module.redis.security_group_id
}
