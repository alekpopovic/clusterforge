# platform/ecs/codedeploy-blue-green

Opt-in ECS blue/green deployment support using CodeDeploy.

Status: experimental.

Use this with an ECS service whose `deployment_controller` is `CODE_DEPLOY`
and two ALB target groups. The existing ECS service module still defaults to
rolling `ECS` deployments.

## Usage

```hcl
module "blue_green" {
  source = "../../../modules/platform/ecs/codedeploy-blue-green"

  name                    = "api-prod"
  environment             = "prod"
  cluster_name            = "clusterforge-prod"
  service_name            = "api"
  alb_listener_arn        = "arn:aws:elasticloadbalancing:region:account:listener/app/example"
  blue_target_group_name  = "api-blue"
  blue_target_group_arn   = "arn:aws:elasticloadbalancing:region:account:targetgroup/api-blue"
  green_target_group_name = "api-green"
  green_target_group_arn  = "arn:aws:elasticloadbalancing:region:account:targetgroup/api-green"
}
```
