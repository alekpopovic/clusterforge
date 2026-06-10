locals {
  module_path     = "core/naming"
  normalized_name = var.name == null ? null : lower(var.name)
  environment     = var.environment == null ? "unknown" : lower(var.environment)
  name_parts      = compact([local.environment, local.normalized_name])
}

# TODO: Implement the core/naming module without adding provider configuration
# to this reusable child module.
