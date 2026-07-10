output "name" {
  description = "Rollout application name."
  value       = local.name
}

output "namespace" {
  description = "Rollout application namespace."
  value       = local.namespace
}

output "stable_service_name" {
  description = "Stable or active Service name."
  value       = kubernetes_service_v1.stable.metadata[0].name
}

output "preview_service_name" {
  description = "Canary or preview Service name."
  value       = kubernetes_service_v1.preview.metadata[0].name
}

output "rollout_manifest" {
  description = "Rendered Argo Rollout manifest."
  value       = kubernetes_manifest.rollout.manifest
}
