output "release_name" {
  description = "Helm release name."
  value       = module.podinfo.release_name
}

output "namespace" {
  description = "Namespace where the Helm app is installed."
  value       = module.podinfo.namespace
}

output "status" {
  description = "Helm release status."
  value       = module.podinfo.status
}
