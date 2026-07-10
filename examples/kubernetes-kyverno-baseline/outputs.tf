output "namespace" {
  description = "Namespace where Kyverno is installed."
  value       = module.kyverno.namespace
}

output "release_name" {
  description = "Kyverno Helm release name."
  value       = module.kyverno.release_name
}

output "baseline_policy_names" {
  description = "Audit baseline policy names when enabled."
  value       = module.kyverno.baseline_policy_names
}
