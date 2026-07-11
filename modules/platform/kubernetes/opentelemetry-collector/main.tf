locals {
  release_name = "opentelemetry-collector"
  base_values = yamlencode({
    mode           = var.mode
    presets        = var.presets == null ? {} : var.presets
    serviceAccount = { annotations = var.service_account_annotations }
  })
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
  repository = "https://open-telemetry.github.io/opentelemetry-helm-charts"
  chart      = "opentelemetry-collector"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = concat([local.base_values], var.values)
  depends_on = [kubernetes_namespace_v1.this]
}
