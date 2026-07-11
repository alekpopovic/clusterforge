output "nomad_consul_config" {
  description = "Rendered Nomad Consul integration JSON."
  value       = local.consul_config
}
