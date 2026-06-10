module "name" {
  source = "../../../modules/core/naming"

  project     = var.project
  environment = var.environment
  name        = "eks"
}

module "tags" {
  source = "../../../modules/core/tags"

  project     = var.project
  environment = var.environment
}

module "network" {
  source = "../../../modules/cloud/aws/network"

  name     = module.name.name
  vpc_cidr = var.vpc_cidr
  tags     = module.tags.tags
}

module "eks" {
  source = "../../../modules/orchestrators/kubernetes/eks"

  name       = module.name.name
  subnet_ids = module.network.private_subnet_ids
  tags       = module.tags.tags
}
