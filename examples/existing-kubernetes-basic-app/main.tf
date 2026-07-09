module "app" {
  source = "../../modules/workloads/kubernetes/app"

  name      = "existing-kubernetes-app"
  namespace = "clusterforge-existing"
  image     = "nginx:1.27"

  ports = [{
    name           = "http"
    container_port = 80
  }]
}
