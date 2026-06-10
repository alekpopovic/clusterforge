locals {
  module_path     = "cloud/aws/storage"
  normalized_name = var.name == null ? null : lower(var.name)
  environment     = var.environment == null ? "unknown" : lower(var.environment)
  common_tags = merge(var.tags, {
    Module      = local.module_path
    Environment = local.environment
  })
}

# TODO: Implement the cloud/aws/storage module without adding provider configuration
# to this reusable child module.
