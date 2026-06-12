output "name" {
  description = "Name of the SecretStore or ClusterSecretStore."
  value       = var.name
}

output "kind" {
  description = "Kind of External Secrets store created."
  value       = var.kind
}

output "manifest" {
  description = "Rendered Kubernetes manifest for the SecretStore or ClusterSecretStore."
  value       = kubernetes_manifest.this.manifest
}
