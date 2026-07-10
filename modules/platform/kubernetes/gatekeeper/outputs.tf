output "namespace" {
  description = "Namespace where Gatekeeper is installed."
  value       = local.namespace
}

output "release_name" {
  description = "Gatekeeper Helm release name."
  value       = helm_release.this.name
}

output "release_status" {
  description = "Gatekeeper Helm release status."
  value       = helm_release.this.status
}
