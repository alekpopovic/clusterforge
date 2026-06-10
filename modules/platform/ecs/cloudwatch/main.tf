locals {
  module_path     = "platform/ecs/cloudwatch"
  normalized_name = var.name == null ? null : lower(var.name)
  environment     = var.environment == null ? "unknown" : lower(var.environment)
  common_tags = merge(var.tags, {
    Module      = local.module_path
    Environment = local.environment
  })
}

# TODO: Implement the platform/ecs/cloudwatch module without adding provider configuration
# to this reusable child module.
