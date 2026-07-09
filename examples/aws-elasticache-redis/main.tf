module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "redis"
}

module "network" {
  source = "../../modules/cloud/aws/network"

  name               = "clusterforge-redis"
  environment        = "dev"
  cidr_block         = "10.90.0.0/16"
  availability_zones = var.availability_zones

  public_subnet_cidrs  = ["10.90.0.0/20", "10.90.16.0/20"]
  private_subnet_cidrs = ["10.90.128.0/20", "10.90.144.0/20"]

  tags = module.tags.tags
}

resource "aws_security_group" "app" {
  name        = "clusterforge-redis-app"
  description = "Example app security group allowed to reach Redis."
  vpc_id      = module.network.vpc_id
  tags        = module.tags.tags
}

module "redis" {
  source = "../../modules/cloud/aws/elasticache-redis"

  name                       = "clusterforge-redis"
  environment                = "dev"
  vpc_id                     = module.network.vpc_id
  subnet_ids                 = module.network.private_subnet_ids
  allowed_security_group_ids = [aws_security_group.app.id]
  tags                       = module.tags.tags
}
