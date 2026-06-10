module "aws_tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "network"
  owner       = "platform-team"
  cost_center = "cc-1234"

  extra_tags = {
    Compliance = "internal"
  }
}

module "kubernetes_labels" {
  source = "../../modules/core/labels"

  project     = "clusterforge"
  environment = "dev"
  app         = "api"
  component   = "web"
  part_of     = "customer-portal"

  extra_labels = {
    "team.example.com/name" = "Platform Team"
  }
}
