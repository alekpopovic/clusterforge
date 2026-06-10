locals {
  base_labels = {
    "clusterforge.io/project"      = var.project
    "clusterforge.io/environment"  = var.environment
    "app.kubernetes.io/managed-by" = var.managed_by
  }

  app_labels = {
    "app.kubernetes.io/name"      = var.app
    "app.kubernetes.io/part-of"   = var.part_of
    "app.kubernetes.io/component" = var.component
  }

  raw_labels = merge(
    local.base_labels,
    {
      for key, value in local.app_labels : key => value
      if length(trimspace(value)) > 0
    },
    {
      for key, value in var.extra_labels : key => value
      if length(trimspace(value)) > 0
    }
  )

  normalized_label_values = {
    for key, value in local.raw_labels : key => trim(
      substr(
        replace(
          replace(lower(trimspace(value)), "/[^0-9a-z_.-]+/", "-"),
          "/[-_.]{2,}/",
          "-"
        ),
        0,
        63
      ),
      "-_."
    )
  }

  labels = {
    for key, value in local.normalized_label_values : key => value
    if length(value) > 0
  }
}
