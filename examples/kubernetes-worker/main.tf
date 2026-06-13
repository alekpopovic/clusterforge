module "worker" {
  source = "../../modules/workloads/kubernetes/worker"

  name      = "email-worker"
  namespace = var.namespace
  image     = "busybox:1.36"
  replicas  = 2

  command = ["/bin/sh", "-c"]
  args    = ["while true; do echo processing queue; sleep 30; done"]

  env = {
    LOG_LEVEL = "info"
    QUEUE     = "emails"
  }

  secret_env = {
    DATABASE_URL = {
      secret_name = "worker-secrets"
      secret_key  = "database-url"
    }
  }

  resources = {
    cpu_request    = "50m"
    memory_request = "64Mi"
    cpu_limit      = "250m"
    memory_limit   = "256Mi"
  }

  labels = {
    "clusterforge.io/example" = "kubernetes-worker"
  }
}
