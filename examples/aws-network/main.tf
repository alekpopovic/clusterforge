module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "network"
}

module "network" {
  source = "../../modules/cloud/aws/network"

  name               = "clusterforge-dev"
  environment        = "dev"
  cidr_block         = "10.40.0.0/16"
  availability_zones = var.availability_zones

  public_subnet_cidrs  = ["10.40.0.0/20", "10.40.16.0/20"]
  private_subnet_cidrs = ["10.40.128.0/20", "10.40.144.0/20"]

  tags = module.tags.tags

  private_subnet_tags = {
    "kubernetes.io/cluster/clusterforge-dev" = "shared"
  }
}
