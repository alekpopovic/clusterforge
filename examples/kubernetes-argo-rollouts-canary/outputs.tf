output "controller_namespace" {
  description = "Namespace containing the Argo Rollouts controller."
  value       = module.argo_rollouts.namespace
}

output "rollout_name" {
  description = "Example Rollout name when enabled."
  value       = var.enable_rollout_app ? module.demo[0].name : null
}
