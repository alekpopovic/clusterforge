output "deployment_name" {
  description = "Created deployment name."
  value       = module.app.deployment_name
}

output "service_name" {
  description = "Created service name."
  value       = module.app.service_name
}

output "namespace" {
  description = "Namespace for the example app."
  value       = module.app.namespace
}
