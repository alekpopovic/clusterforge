# platform/ecs/codedeploy-blue-green

Opt-in ECS blue/green deployment support using CodeDeploy.

Status: experimental.

Use this with an ECS service whose `deployment_controller` is `CODE_DEPLOY`
and two ALB target groups. The existing ECS service module still defaults to
rolling `ECS` deployments.
