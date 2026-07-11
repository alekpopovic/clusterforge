output "client_config" {
  description = "Rendered Nomad client configuration."
  value       = module.cluster_config.client_config
}
