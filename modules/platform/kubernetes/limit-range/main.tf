locals {
  namespace = trimspace(var.namespace)
  name      = "limit-range"
  labels = merge(var.labels, {
    "app.kubernetes.io/managed-by" = "terraform"
    "clusterforge.io/component"    = "limit-range"
  })
}

resource "kubernetes_limit_range_v1" "this" {
  metadata {
    name      = local.name
    namespace = local.namespace
    labels    = local.labels
  }

  spec {
    dynamic "limit" {
      for_each = var.limits

      content {
        type            = limit.value.type
        default         = limit.value.default
        default_request = limit.value.default_request
        max             = limit.value.max
        min             = limit.value.min
      }
    }
  }
}
