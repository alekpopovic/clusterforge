locals {
  module_path     = "orchestrators/ecs/cluster"
  normalized_name = var.name == null ? null : lower(var.name)
  environment     = var.environment == null ? "unknown" : lower(var.environment)
  common_tags = merge(var.tags, {
    Module      = local.module_path
    Environment = local.environment
  })
}

# TODO: Implement the orchestrators/ecs/cluster module without adding provider configuration
# to this reusable child module.
