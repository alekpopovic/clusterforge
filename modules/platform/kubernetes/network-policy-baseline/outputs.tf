output "namespace" {
  description = "Namespace where baseline NetworkPolicies are created."
  value       = var.namespace
}

output "network_policy_names" {
  description = "NetworkPolicy names created by this module."
  value = compact([
    var.default_deny_ingress ? kubernetes_network_policy_v1.default_deny_ingress[0].metadata[0].name : "",
    var.default_deny_egress ? kubernetes_network_policy_v1.default_deny_egress[0].metadata[0].name : "",
    var.default_deny_egress && var.allow_dns_egress ? kubernetes_network_policy_v1.allow_dns_egress[0].metadata[0].name : ""
  ])
}
