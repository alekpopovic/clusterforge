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

  dynamic "set" {
    for_each = var.set

    content {
      name  = set.key
      value = set.value
    }
  }

  dynamic "set_sensitive" {
    for_each = var.set_sensitive

    content {
      name  = set_sensitive.key
      value = set_sensitive.value
    }
  }

  depends_on = [kubernetes_namespace_v1.this]
}
