module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "prod"
  component   = "vpc-endpoints"
}

module "network" {
  source = "../../modules/cloud/aws/network"

  name               = "clusterforge-private-eks"
  environment        = "prod"
  cidr_block         = "10.70.0.0/16"
  availability_zones = var.availability_zones

  public_subnet_cidrs  = ["10.70.0.0/20", "10.70.16.0/20"]
  private_subnet_cidrs = ["10.70.128.0/20", "10.70.144.0/20"]

  tags = module.tags.tags
}

module "vpc_endpoints" {
  source = "../../modules/cloud/aws/vpc-endpoints"

  name            = "clusterforge-private-eks"
  environment     = "prod"
  vpc_id          = module.network.vpc_id
  subnet_ids      = module.network.private_subnet_ids
  route_table_ids = module.network.private_route_table_ids

  allowed_security_group_ids = []

  gateway_endpoints = ["s3"]
  interface_endpoints = [
    "ecr.api",
    "ecr.dkr",
    "logs",
    "sts",
  ]

  tags = module.tags.tags
}
