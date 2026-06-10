locals {
  standard_tags = {
    Project     = var.project
    Environment = var.environment
    ManagedBy   = var.managed_by
    Framework   = "ClusterForge"
  }

  tags = merge(local.standard_tags, var.extra_tags)
}
