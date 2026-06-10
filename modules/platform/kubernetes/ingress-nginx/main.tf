locals {
  release_name = "ingress-nginx"
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
  repository = "https://kubernetes.github.io/ingress-nginx"
  chart      = "ingress-nginx"
  version    = var.chart_version == "" ? null : var.chart_version
  values     = var.values

  depends_on = [kubernetes_namespace_v1.this]
}
