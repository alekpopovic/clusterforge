output "cluster_id" {
  description = "Example ECS cluster ID."
  value       = module.ecs_cluster.cluster_id
}

output "cluster_arn" {
  description = "Example ECS cluster ARN."
  value       = module.ecs_cluster.cluster_arn
}

output "cluster_name" {
  description = "Example ECS cluster name."
  value       = module.ecs_cluster.cluster_name
}
