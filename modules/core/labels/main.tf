locals {
  standard_labels = {
    "app.kubernetes.io/name"       = var.app
    "app.kubernetes.io/component"  = var.component
    "app.kubernetes.io/part-of"    = var.part_of
    "app.kubernetes.io/managed-by" = var.managed_by
  }

  labels = merge(local.standard_labels, var.extra_labels)
}
