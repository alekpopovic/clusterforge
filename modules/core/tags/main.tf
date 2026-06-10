locals {
  standard_tags = {
    Project     = trimspace(var.project)
    Environment = trimspace(var.environment)
    ManagedBy   = trimspace(var.managed_by)
  }

  optional_tags = {
    Component  = trimspace(var.component)
    Owner      = trimspace(var.owner)
    CostCenter = trimspace(var.cost_center)
  }

  non_empty_optional_tags = {
    for key, value in local.optional_tags : key => value
    if length(value) > 0
  }

  non_empty_extra_tags = {
    for key, value in var.extra_tags : key => value
    if length(trimspace(value)) > 0
  }

  tags = merge(local.standard_tags, local.non_empty_optional_tags, local.non_empty_extra_tags)
}
