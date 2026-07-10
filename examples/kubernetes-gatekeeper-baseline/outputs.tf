output "namespace" {
  description = "Namespace where Gatekeeper is installed."
  value       = module.gatekeeper.namespace
}

output "release_name" {
  description = "Gatekeeper Helm release name."
  value       = module.gatekeeper.release_name
}
