output "deployment_name" {
  description = "Deployment name."
  value       = kubernetes_deployment_v1.this.metadata[0].name
}

output "service_name" {
  description = "Service name."
  value       = kubernetes_service_v1.this.metadata[0].name
}

output "namespace" {
  description = "Namespace where the application was deployed."
  value       = var.namespace
}
