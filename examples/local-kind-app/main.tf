module "app" {
  source = "../../modules/workloads/kubernetes/app"

  name      = "local-kind-app"
  namespace = "clusterforge-local"
  image     = "nginx:1.27"

  ports = [{
    name           = "http"
    container_port = 80
  }]
}
