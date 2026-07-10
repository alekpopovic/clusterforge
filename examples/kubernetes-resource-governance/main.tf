module "quota" {
  source = "../../modules/platform/kubernetes/resource-quota"

  namespace = var.namespace
  hard = {
    "requests.cpu"    = "4"
    "requests.memory" = "8Gi"
    "limits.cpu"      = "8"
    "limits.memory"   = "16Gi"
    pods              = "40"
  }
}

module "limits" {
  source = "../../modules/platform/kubernetes/limit-range"

  namespace = var.namespace
  limits = [{
    type = "Container"
    default_request = {
      cpu    = "100m"
      memory = "128Mi"
    }
    default = {
      cpu    = "500m"
      memory = "512Mi"
    }
    max = {
      cpu    = "2"
      memory = "2Gi"
    }
  }]
}
