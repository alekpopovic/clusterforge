module "aws_resource_name" {
  source = "../../modules/core/naming"

  project     = "clusterforge"
  environment = "dev"
  component   = "network"
  name        = "main-vpc"
}

module "kubernetes_app_name" {
  source = "../../modules/core/naming"

  project     = "clusterforge"
  environment = "staging"
  component   = "app"
  name        = "api_server"
  extra_parts = ["blue"]
}

module "platform_component_name" {
  source = "../../modules/core/naming"

  project     = "clusterforge"
  environment = "prod"
  component   = "platform"
  name        = "ingress-nginx"
  suffix      = "controller"
  max_length  = 48
}
