output "namespace" {
  description = "Namespace where Argo CD is installed."
  value       = module.argocd.namespace
}

output "release_name" {
  description = "Argo CD Helm release name."
  value       = module.argocd.release_name
}

output "app_of_apps_name" {
  description = "App-of-apps Application name."
  value       = module.argocd.app_of_apps_name
}
