locals {
  name      = trimspace(var.name)
  namespace = trimspace(var.namespace)

  labels = merge(
    {
      "app.kubernetes.io/name"       = local.name
      "app.kubernetes.io/instance"   = local.name
      "app.kubernetes.io/component"  = "helm-app"
      "app.kubernetes.io/managed-by" = "terraform"
    },
    var.labels
  )
}

resource "kubernetes_namespace_v1" "this" {
  count = var.create_namespace ? 1 : 0

  metadata {
    name   = local.namespace
    labels = local.labels
  }
}

resource "helm_release" "this" {
  name             = local.name
  namespace        = local.namespace
  repository       = var.repository
  chart            = var.chart
  version          = var.chart_version == "" ? null : var.chart_version
  values           = var.values
  timeout          = var.timeout
  atomic           = var.atomic
  cleanup_on_fail  = var.cleanup_on_fail
  wait             = var.wait
  create_namespace = false
  set = [
    for name, value in var.set : {
      name  = name
      value = value
    }
  ]
  set_sensitive = [
    for name, value in var.set_sensitive : {
      name  = name
      value = value
    }
  ]

  depends_on = [kubernetes_namespace_v1.this]
}
