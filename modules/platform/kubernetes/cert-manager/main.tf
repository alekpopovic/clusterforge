locals {
  release_name = "cert-manager"

  default_values = yamlencode({
    serviceAccount = {
      annotations = var.service_account_annotations
    }
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
  repository = "https://charts.jetstack.io"
  chart      = "cert-manager"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = concat([local.default_values], var.values)

  depends_on = [kubernetes_namespace_v1.this]
}
