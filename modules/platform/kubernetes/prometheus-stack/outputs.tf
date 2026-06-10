output "namespace" {
  description = "Namespace where kube-prometheus-stack is installed."
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
