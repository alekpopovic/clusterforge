## Prompt 13 — ECS service workload modul

```text
Implement modules/workloads/ecs/service.

Purpose:
Deploy a generic ECS Fargate service.

Provider:
- hashicorp/aws.
- Provider configured in root.

Inputs:
- name: string
- environment: string
- cluster_arn: string
- subnet_ids: list(string)
- security_group_ids: list(string)
- assign_public_ip: bool, default false
- desired_count: number, default 1
- image: string
- cpu: number, default 256
- memory: number, default 512
- container_port: number
- protocol: string, default "tcp"
- environment_variables: map(string), default {}
- secrets: list(object({
    name = string
    value_from = string
  })), default []
- execution_role_arn: string, default ""
- task_role_arn: string, default ""
- create_log_group: bool, default true
- log_retention_in_days: number, default 30
- tags: map(string), default {}

Load balancer:
- load_balancer: object({
    enabled = bool
    target_group_arn = optional(string)
    container_name = optional(string)
    container_port = optional(number)
  }), default disabled

Autoscaling:
- autoscaling: object({
    enabled = bool
    min_capacity = optional(number, 1)
    max_capacity = optional(number, 3)
    cpu_target_value = optional(number, 70)
  }), default disabled

Resources:
- aws_cloudwatch_log_group optional
- IAM execution role optional only if execution_role_arn not provided
- IAM task role optional only if task_role_arn not provided
- aws_ecs_task_definition
- aws_ecs_service
- aws_appautoscaling_target optional
- aws_appautoscaling_policy optional

Behavior:
- Avoid storing secret values; use value_from ARNs for secrets.
- Use awslogs log configuration.
- Support Fargate network mode awsvpc.
- Use load balancer block only when enabled.

Validation:
- name, cluster_arn, image non-empty.
- subnet_ids length > 0.
- container_port > 0.
- cpu/memory should be valid common Fargate combinations if practical; otherwise document.

Outputs:
- service_name
- service_id
- task_definition_arn
- log_group_name

README:
- Basic service example.
- Service with ALB target group example.
- Secrets example using SSM/Secrets Manager ARN references.
- Explain difference between ECS workload and Kubernetes workload.

Run terraform fmt -recursive.
```

---
