module "tags" {
  source = "../../../modules/core/tags"

  project     = var.project
  environment = var.environment
  component   = "eks"
  owner       = var.owner
  cost_center = var.cost_center
  extra_tags  = var.extra_tags
}

module "network" {
  source = "../../../modules/cloud/aws/network"

  name               = var.name
  environment        = var.environment
  cidr_block         = var.vpc_cidr_block
  availability_zones = var.availability_zones

  public_subnet_cidrs  = var.public_subnet_cidrs
  private_subnet_cidrs = var.private_subnet_cidrs

  enable_nat_gateway = var.enable_nat_gateway
  single_nat_gateway = var.single_nat_gateway

  tags = module.tags.tags

  public_subnet_tags = merge(
    {
      "kubernetes.io/cluster/${var.name}" = "shared"
    },
    var.public_subnet_tags
  )

  private_subnet_tags = merge(
    {
      "kubernetes.io/cluster/${var.name}" = "shared"
    },
    var.private_subnet_tags
  )
}

module "eks" {
  source = "../../../modules/orchestrators/kubernetes/eks"

  name                      = var.name
  environment               = var.environment
  kubernetes_version        = var.kubernetes_version
  vpc_id                    = module.network.vpc_id
  subnet_ids                = module.network.private_subnet_ids
  endpoint_public_access    = var.endpoint_public_access
  endpoint_private_access   = var.endpoint_private_access
  public_access_cidrs       = var.public_access_cidrs
  enabled_cluster_log_types = var.enabled_cluster_log_types
  tags                      = module.tags.tags

  enable_vpc_cni_addon        = var.enable_vpc_cni_addon
  enable_coredns_addon        = var.enable_coredns_addon
  enable_kube_proxy_addon     = var.enable_kube_proxy_addon
  enable_ebs_csi_driver_addon = var.enable_ebs_csi_driver_addon

  node_groups = {
    default = {
      instance_types = var.default_node_instance_types
      capacity_type  = var.default_node_capacity_type
      min_size       = var.default_node_min_size
      desired_size   = var.default_node_desired_size
      max_size       = var.default_node_max_size
      disk_size      = var.default_node_disk_size
      labels = {
        role = "default"
      }
    }
  }
}
