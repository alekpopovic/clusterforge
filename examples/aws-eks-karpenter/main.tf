module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "karpenter"
}

module "network" {
  source = "../../modules/cloud/aws/network"

  name               = "clusterforge-karpenter-dev"
  environment        = "dev"
  cidr_block         = "10.50.0.0/16"
  availability_zones = var.availability_zones

  public_subnet_cidrs  = ["10.50.0.0/20", "10.50.16.0/20"]
  private_subnet_cidrs = ["10.50.128.0/20", "10.50.144.0/20"]

  tags = module.tags.tags

  private_subnet_tags = {
    "kubernetes.io/cluster/clusterforge-karpenter-dev" = "shared"
  }
}

module "eks" {
  source = "../../modules/orchestrators/kubernetes/eks"

  name        = "clusterforge-karpenter-dev"
  environment = "dev"
  vpc_id      = module.network.vpc_id
  subnet_ids  = module.network.private_subnet_ids
  tags        = module.tags.tags

  enable_irsa = true

  # Keep a small managed node group so system add-ons and Karpenter itself have
  # bootstrap capacity. Karpenter-created nodes should be governed separately
  # through reviewed NodePool and EC2NodeClass manifests.
  node_groups = {
    system = {
      instance_types = ["t3.medium"]
      min_size       = 1
      desired_size   = 2
      max_size       = 3
      labels = {
        role = "system"
      }
    }
  }
}

module "karpenter_irsa" {
  source = "../../modules/cloud/aws/karpenter-irsa"

  cluster_name      = module.eks.cluster_name
  oidc_provider_arn = module.eks.oidc_provider_arn
  oidc_provider_url = module.eks.oidc_provider_url
  tags              = module.tags.tags
}

data "aws_eks_cluster_auth" "this" {
  name = module.eks.cluster_name
}

module "karpenter" {
  source = "../../modules/platform/kubernetes/karpenter"

  chart_version            = var.karpenter_chart_version
  cluster_name             = module.eks.cluster_name
  cluster_endpoint         = module.eks.cluster_endpoint
  service_account_role_arn = module.karpenter_irsa.role_arn

  labels = {
    "clusterforge.io/environment" = "dev"
    "clusterforge.io/managed-by"  = "terraform"
  }

  depends_on = [module.eks]
}
