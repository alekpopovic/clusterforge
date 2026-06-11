module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "ecs"
}

module "ecs_cluster" {
  source = "../../modules/orchestrators/ecs/cluster"

  name        = "clusterforge-dev-ecs"
  environment = "dev"
  tags        = module.tags.tags
}
