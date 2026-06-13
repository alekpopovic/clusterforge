output "deployment_name" {
  description = "Created worker Deployment name."
  value       = module.worker.deployment_name
}

output "namespace" {
  description = "Namespace for the worker example."
  value       = module.worker.namespace
}
