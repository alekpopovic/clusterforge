output "cluster_name" {
  description = "Example ECS cluster name."
  value       = module.ecs_cluster.cluster_name
}

output "cluster_arn" {
  description = "Example ECS cluster ARN."
  value       = module.ecs_cluster.cluster_arn
}

output "service_name" {
  description = "Example ECS service name."
  value       = module.service.service_name
}

output "private_subnet_ids" {
  description = "Private subnet IDs used by the service."
  value       = module.network.private_subnet_ids
}
