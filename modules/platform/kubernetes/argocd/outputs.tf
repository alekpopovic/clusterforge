output "namespace" {
  description = "Namespace where argo-cd is installed."
  value       = var.namespace
}

output "release_name" {
  description = "Helm release name."
  value       = helm_release.this.name
}

output "release_status" {
  description = "Helm release status."
  value       = helm_release.this.status
}

output "app_of_apps_name" {
  description = "Name of the app-of-apps Application when enabled."
  value       = var.enable_app_of_apps ? var.app_of_apps_name : null
}
