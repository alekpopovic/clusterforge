module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "postgres"
}

module "network" {
  source = "../../modules/cloud/aws/network"

  name               = "clusterforge-postgres"
  environment        = "dev"
  cidr_block         = "10.80.0.0/16"
  availability_zones = var.availability_zones

  public_subnet_cidrs  = ["10.80.0.0/20", "10.80.16.0/20"]
  private_subnet_cidrs = ["10.80.128.0/20", "10.80.144.0/20"]

  tags = module.tags.tags
}

resource "aws_security_group" "app" {
  name        = "clusterforge-postgres-app"
  description = "Example app security group allowed to reach Postgres."
  vpc_id      = module.network.vpc_id
  tags        = module.tags.tags
}

module "postgres" {
  source = "../../modules/cloud/aws/rds-postgres"

  name                       = "clusterforge-postgres"
  environment                = "dev"
  vpc_id                     = module.network.vpc_id
  subnet_ids                 = module.network.private_subnet_ids
  allowed_security_group_ids = [aws_security_group.app.id]
  database_name              = "app"
  tags                       = module.tags.tags
}
