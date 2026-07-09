module "k3s" {
  source = "../../modules/orchestrators/kubernetes/k3s"

  cluster_name = "clusterforge-dev-k3s"
  environment  = "dev"
}
