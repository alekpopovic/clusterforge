module "app" {
  source = "../../modules/workloads/kubernetes/app"

  name      = "advanced-app"
  namespace = "clusterforge-advanced"
  image     = "nginx:1.27"

  service_account = {
    create          = true
    automount_token = false
    annotations     = {}
  }

  rbac = {
    create = true
    rules = [{
      api_groups = [""]
      resources  = ["configmaps"]
      verbs      = ["get", "list"]
    }]
  }
}
