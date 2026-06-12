module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "eks"
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

module "eks" {
  source = "../../modules/orchestrators/kubernetes/eks"

  name        = "clusterforge-dev"
  environment = "dev"
  vpc_id      = module.network.vpc_id
  subnet_ids  = module.network.private_subnet_ids
  tags        = module.tags.tags

  enable_irsa                 = true
  enable_ebs_csi_driver_addon = true

  node_groups = {
    default = {
      instance_types = ["t3.medium"]
      min_size       = 1
      desired_size   = 2
      max_size       = 4
      disk_size      = 50
      labels = {
        role = "default"
      }
    }
  }
}
