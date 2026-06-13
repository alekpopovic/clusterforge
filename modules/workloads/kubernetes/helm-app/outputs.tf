output "release_name" {
  description = "Helm release name."
  value       = helm_release.this.name
}

output "namespace" {
  description = "Namespace where the Helm release is installed."
  value       = local.namespace
}

output "chart" {
  description = "Helm chart name."
  value       = helm_release.this.chart
}

output "version" {
  description = "Helm chart version selected for the release."
  value       = helm_release.this.version
}

output "status" {
  description = "Helm release status."
  value       = helm_release.this.status
}
