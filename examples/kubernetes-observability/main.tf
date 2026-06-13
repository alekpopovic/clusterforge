module "prometheus_stack" {
  source = "../../modules/platform/kubernetes/prometheus-stack"

  namespace        = "monitoring"
  create_namespace = true

  enable_grafana_ingress = var.grafana_host != ""
  grafana_host           = var.grafana_host
  storage_enabled        = var.enable_persistent_storage
  storage_class_name     = var.storage_class_name

  labels = {
    "clusterforge.io/example" = "kubernetes-observability"
  }
}

module "loki" {
  source = "../../modules/platform/kubernetes/loki"

  namespace        = "logging"
  create_namespace = true

  storage_enabled    = var.enable_persistent_storage
  storage_class_name = var.storage_class_name

  labels = {
    "clusterforge.io/example" = "kubernetes-observability"
  }
}

module "alloy" {
  source = "../../modules/platform/kubernetes/alloy"

  namespace        = "logging"
  create_namespace = false
  values           = var.alloy_values

  labels = {
    "clusterforge.io/example" = "kubernetes-observability"
  }

  depends_on = [module.loki]
}
