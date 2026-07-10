output "namespace" {
  description = "Namespace where Argo Rollouts is installed."
  value       = local.namespace
}

output "release_name" {
  description = "Argo Rollouts Helm release name."
  value       = helm_release.this.name
}

output "release_status" {
  description = "Argo Rollouts Helm release status."
  value       = helm_release.this.status
}
