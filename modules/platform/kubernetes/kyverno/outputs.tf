output "namespace" {
  description = "Namespace where Kyverno is installed."
  value       = local.namespace
}

output "release_name" {
  description = "Kyverno Helm release name."
  value       = helm_release.this.name
}

output "release_status" {
  description = "Kyverno Helm release status."
  value       = helm_release.this.status
}

output "baseline_policy_names" {
  description = "Baseline ClusterPolicy names keyed by logical policy name."
  value       = { for key, policy in local.baseline_policies : key => policy.metadata.name }
}
