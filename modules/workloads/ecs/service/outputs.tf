output "service_name" {
  description = "ECS service name."
  value       = aws_ecs_service.this.name
}

output "service_id" {
  description = "ECS service ID."
  value       = aws_ecs_service.this.id
}

output "task_definition_arn" {
  description = "ECS task definition ARN."
  value       = aws_ecs_task_definition.this.arn
}

output "log_group_name" {
  description = "CloudWatch log group name used by the task."
  value       = local.log_group
}
