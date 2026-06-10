module "app" {
  source = "../../modules/workloads/kubernetes/app"

  name      = "hello"
  namespace = var.namespace
  image     = "nginx:1.27"
  replicas  = 2

  labels = {
    "clusterforge.io/example" = "kubernetes-basic-app"
  }

  ports = [{
    name           = "http"
    container_port = 80
  }]

  service = {
    enabled          = true
    type             = "ClusterIP"
    port             = 80
    target_port_name = "http"
  }

  readiness_probe = {
    path                  = "/"
    port                  = "http"
    initial_delay_seconds = 5
    period_seconds        = 10
    timeout_seconds       = 2
    failure_threshold     = 3
  }
}
