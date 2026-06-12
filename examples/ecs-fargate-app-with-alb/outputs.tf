output "cluster_name" {
  description = "Example ECS cluster name."
  value       = module.ecs_cluster.cluster_name
}

output "service_name" {
  description = "Example ECS service name."
  value       = module.service.service_name
}

output "alb_dns_name" {
  description = "Application Load Balancer DNS name."
  value       = module.alb.alb_dns_name
}

output "target_group_arns" {
  description = "Target group ARNs created by the ALB module."
  value       = module.alb.target_group_arns
}
