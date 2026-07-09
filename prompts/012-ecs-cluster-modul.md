## Prompt 12 — ECS cluster modul

```text
Implement modules/orchestrators/ecs/cluster.

Purpose:
Create an AWS ECS cluster suitable for Fargate services.

Provider:
- hashicorp/aws
- Provider configured in root.

Inputs:
- name: string
- environment: string
- tags: map(string), default {}
- enable_container_insights: bool, default true
- capacity_providers: list(string), default ["FARGATE", "FARGATE_SPOT"]
- default_capacity_provider_strategy: list(object({
    capacity_provider = string
    weight = optional(number)
    base = optional(number)
  })), default []

Resources:
- aws_ecs_cluster
- aws_ecs_cluster_capacity_providers

Behavior:
- Enable containerInsights setting when requested.
- Attach capacity providers.
- Use FARGATE default strategy if none provided.

Outputs:
- cluster_id
- cluster_arn
- cluster_name

README:
- Explain Fargate/ECS use case.
- Show root usage.
- Mention that services are created by workloads/ecs/service.

Create examples/ecs-cluster-minimal.

Run terraform fmt -recursive.
```

---
