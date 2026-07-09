# Wire this module to existing ECS service and ALB target groups after review.
module "blue_green" {
  source = "../../modules/platform/ecs/codedeploy-blue-green"

  name                    = "clusterforge-demo"
  environment             = "dev"
  cluster_name            = "replace-with-cluster"
  service_name            = "replace-with-service"
  alb_listener_arn        = "replace-with-listener-arn"
  blue_target_group_name  = "replace-with-blue"
  blue_target_group_arn   = "replace-with-blue-arn"
  green_target_group_name = "replace-with-green"
  green_target_group_arn  = "replace-with-green-arn"
}
