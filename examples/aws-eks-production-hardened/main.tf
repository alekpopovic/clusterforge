module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "prod"
  component   = "eks"
}

module "network" {
  source = "../../modules/cloud/aws/network"

  name               = "clusterforge-prod"
  environment        = "prod"
  cidr_block         = "10.60.0.0/16"
  availability_zones = var.availability_zones

  public_subnet_cidrs  = ["10.60.0.0/20", "10.60.16.0/20", "10.60.32.0/20"]
  private_subnet_cidrs = ["10.60.128.0/20", "10.60.144.0/20", "10.60.160.0/20"]

  tags = module.tags.tags

  private_subnet_tags = {
    "kubernetes.io/cluster/clusterforge-prod" = "shared"
  }
}

module "eks" {
  source = "../../modules/orchestrators/kubernetes/eks"

  name        = "clusterforge-prod"
  environment = "prod"
  vpc_id      = module.network.vpc_id
  subnet_ids  = module.network.private_subnet_ids
  tags        = module.tags.tags

  endpoint_public_access  = false
  endpoint_private_access = true
  public_access_cidrs     = []

  enabled_cluster_log_types  = ["api", "audit", "authenticator", "controllerManager", "scheduler"]
  cluster_log_retention_days = 90

  enable_cluster_encryption = true
  create_kms_key            = true

  enable_irsa                 = true
  enable_ebs_csi_driver_addon = true

  node_group_ami_type             = "AL2023_x86_64_STANDARD"
  node_group_force_update_version = false
  node_group_update_config = {
    max_unavailable_percentage = 25
  }

  node_groups = {
    system = {
      instance_types = ["m6i.large"]
      min_size       = 2
      desired_size   = 3
      max_size       = 6
      disk_size      = 80
      labels = {
        role = "system"
      }
    }
  }
}
