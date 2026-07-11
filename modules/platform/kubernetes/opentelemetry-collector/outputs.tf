output "namespace" {
  description = "Collector namespace."
  value       = var.namespace
}
output "release_name" {
  description = "Helm release name."
  value       = helm_release.this.name
}
