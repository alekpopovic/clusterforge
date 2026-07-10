output "name" {
  description = "Worker logical name."
  value       = local.name
}

output "namespace" {
  description = "Worker namespace."
  value       = local.namespace
}

output "deployment_name" {
  description = "Kubernetes Deployment name."
  value       = kubernetes_deployment_v1.this.metadata[0].name
}

output "labels" {
  description = "Labels applied to worker resources."
  value       = local.labels
}

output "service_account_name" {
  description = "Service account used by the workload, or null when the namespace default is used."
  value       = var.service_account_name == "" ? null : var.service_account_name
}
