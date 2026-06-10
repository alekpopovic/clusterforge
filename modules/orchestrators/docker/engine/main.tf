locals {
  module_path     = "orchestrators/docker/engine"
  normalized_name = var.name == null ? null : lower(var.name)
  environment     = var.environment == null ? "unknown" : lower(var.environment)
  common_labels = merge(var.labels, {
    "clusterforge.io/module"      = local.module_path
    "clusterforge.io/environment" = local.environment
  })
}

# TODO: Implement the orchestrators/docker/engine module without adding provider configuration
# to this reusable child module.
