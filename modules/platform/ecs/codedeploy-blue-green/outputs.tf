output "codedeploy_app_name" {
  description = "CodeDeploy application name."
  value       = aws_codedeploy_app.this.name
}

output "deployment_group_name" {
  description = "CodeDeploy deployment group name."
  value       = aws_codedeploy_deployment_group.this.deployment_group_name
}

output "codedeploy_role_arn" {
  description = "CodeDeploy service role ARN."
  value       = aws_iam_role.codedeploy.arn
}
