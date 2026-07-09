module "rke2" {
  source = "../../modules/orchestrators/kubernetes/rke2"

  cluster_name = "clusterforge-dev-rke2"
  environment  = "dev"
}
