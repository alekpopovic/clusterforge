locals {
  namespace = trimspace(var.namespace)
  name      = "resource-quota"
  labels = merge(var.labels, {
    "app.kubernetes.io/managed-by" = "terraform"
    "clusterforge.io/component"    = "resource-quota"
  })
}

resource "kubernetes_resource_quota_v1" "this" {
  metadata {
    name      = local.name
    namespace = local.namespace
    labels    = local.labels
  }

  spec {
    hard = var.hard
  }
}
