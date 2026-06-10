locals {
  module_path     = "workloads/kubernetes/helm-app"
  normalized_name = var.name == null ? null : lower(var.name)
  environment     = var.environment == null ? "unknown" : lower(var.environment)
  common_labels = merge(var.labels, {
    "clusterforge.io/module"      = local.module_path
    "clusterforge.io/environment" = local.environment
  })
}

# TODO: Implement the workloads/kubernetes/helm-app module without adding provider configuration
# to this reusable child module.
