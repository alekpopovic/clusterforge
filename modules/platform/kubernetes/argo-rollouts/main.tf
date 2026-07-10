locals {
  namespace    = trimspace(var.namespace)
  release_name = "argo-rollouts"
  labels = merge(var.labels, {
    "app.kubernetes.io/managed-by" = "terraform"
    "clusterforge.io/component"    = "argo-rollouts"
  })
  dashboard_values = var.enable_dashboard_ingress ? [yamlencode({
    dashboard = {
      enabled = true
      ingress = {
        enabled = true
        hosts   = [var.dashboard_host]
      }
    }
  })] : []
}

resource "kubernetes_namespace_v1" "this" {
  count = var.create_namespace ? 1 : 0
  metadata {
    name   = local.namespace
    labels = local.labels
  }
}

resource "helm_release" "this" {
  name       = local.release_name
  namespace  = local.namespace
  repository = "https://argoproj.github.io/argo-helm"
  chart      = "argo-rollouts"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = concat(local.dashboard_values, var.values)
  depends_on = [kubernetes_namespace_v1.this]
}
