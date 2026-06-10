output "name" {
  description = "Workload name."
  value       = local.name
}

output "namespace" {
  description = "Workload namespace."
  value       = local.namespace
}

output "deployment_name" {
  description = "Deployment name."
  value       = kubernetes_deployment_v1.this.metadata[0].name
}

output "service_name" {
  description = "Service name, or null when service is disabled."
  value       = var.service.enabled ? kubernetes_service_v1.this[0].metadata[0].name : null
}

output "ingress_name" {
  description = "Ingress name, or null when ingress is disabled."
  value       = var.ingress.enabled ? kubernetes_ingress_v1.this[0].metadata[0].name : null
}

output "labels" {
  description = "Labels applied to workload resources."
  value       = local.labels
}
