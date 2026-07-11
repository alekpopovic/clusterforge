output "server_config" {
  description = "Rendered Nomad server JSON configuration."
  value       = local.server_config
}
output "client_config" {
  description = "Rendered Nomad client JSON configuration."
  value       = local.client_config
}
output "client_cloud_init" {
  description = "Optional client cloud-init user data; review before use."
  value       = local.cloud_init
}
output "install_notes" {
  description = "Manual bootstrap boundary."
  value       = "Install Nomad on operator-managed hosts, distribute TLS/ACL material outside Terraform, then apply the rendered configuration."
}
