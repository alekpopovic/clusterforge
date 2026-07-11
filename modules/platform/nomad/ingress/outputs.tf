output "service_tags" {
  description = "Service registration tags consumed by an operator-installed ingress controller."
  value       = local.service_tags
}
