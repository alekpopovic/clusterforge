## Prompt 74 — ECS blue/green deployment design

```text
Create design and partial implementation for ECS blue/green deployments.

Create RFC:
- docs/rfcs/005-ecs-blue-green.md

Goal:
Support safer ECS/Fargate production deployments using CodeDeploy and ALB target groups.

Cover:
- ECS service deployment controller CODE_DEPLOY
- two target groups
- listener rules
- CodeDeploy application
- CodeDeploy deployment group
- rollback settings
- health checks
- IAM roles

Implement if feasible:
- modules/platform/ecs/codedeploy-blue-green
- Update workloads/ecs/service with deployment_controller option

Inputs:
- name
- environment
- cluster_name
- service_name
- alb_listener_arn
- blue_target_group_name/arn
- green_target_group_name/arn
- termination_wait_time
- auto_rollback_enabled
- tags

Rules:
- Do not break existing ECS service module.
- Keep default deployment simple rolling update.
- Blue/green must be opt-in.
- Document limitations.

Create example:
- examples/ecs-blue-green

Run:
- terraform fmt -recursive
```

---
