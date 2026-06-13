locals {
  release_name = "loki"

  generated_values = var.storage_enabled ? {
    singleBinary = {
      persistence = {
        enabled      = true
        storageClass = var.storage_class_name == "" ? null : var.storage_class_name
        size         = var.storage_size
      }
    }
  } : {}

  generated_values_yaml = length(keys(local.generated_values)) > 0 ? [yamlencode(local.generated_values)] : []
}

resource "kubernetes_namespace_v1" "this" {
  count = var.create_namespace ? 1 : 0

  metadata {
    name   = var.namespace
    labels = var.labels
  }
}

resource "helm_release" "this" {
  name       = local.release_name
  namespace  = var.namespace
  repository = "https://grafana.github.io/helm-charts"
  chart      = "loki"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = concat(local.generated_values_yaml, var.values)

  depends_on = [kubernetes_namespace_v1.this]
}
