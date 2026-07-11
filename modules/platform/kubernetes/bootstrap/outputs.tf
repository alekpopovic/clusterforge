output "enabled_addons" {
  description = "Names of enabled platform add-ons."
  value = compact([
    var.enable_ingress_nginx ? "ingress-nginx" : "",
    var.enable_cert_manager ? "cert-manager" : "",
    var.enable_external_dns ? "external-dns" : "",
    var.enable_external_secrets ? "external-secrets" : "",
    var.enable_pod_security ? "pod-security" : "",
    var.enable_network_policy_baseline ? "network-policy-baseline" : "",
    var.enable_karpenter ? "karpenter" : "",
    var.enable_metrics_server ? "metrics-server" : "",
    var.enable_prometheus_stack ? "prometheus-stack" : "",
    var.enable_loki ? "loki" : "",
    var.enable_log_agent ? "alloy" : "",
    var.enable_argocd ? "argocd" : ""
    , var.enable_opentelemetry_collector ? "opentelemetry-collector" : ""
  ])
}

output "namespaces" {
  description = "Namespaces selected for platform add-ons."
  value       = local.namespaces
}

output "releases" {
  description = "Enabled Helm release names by add-on."
  value = merge(
    var.enable_ingress_nginx ? { ingress_nginx = module.ingress_nginx[0].release_name } : {},
    var.enable_cert_manager ? { cert_manager = module.cert_manager[0].release_name } : {},
    var.enable_external_dns ? { external_dns = module.external_dns[0].release_name } : {},
    var.enable_external_secrets ? { external_secrets = module.external_secrets[0].release_name } : {},
    var.enable_karpenter ? { karpenter = module.karpenter[0].release_name } : {},
    var.enable_metrics_server ? { metrics_server = module.metrics_server[0].release_name } : {},
    var.enable_prometheus_stack ? { prometheus_stack = module.prometheus_stack[0].release_name } : {},
    var.enable_loki ? { loki = module.loki[0].release_name } : {},
    var.enable_log_agent ? { log_agent = module.log_agent[0].release_name } : {},
    var.enable_argocd ? { argocd = module.argocd[0].release_name } : {}
    , var.enable_opentelemetry_collector ? { opentelemetry_collector = module.opentelemetry_collector[0].release_name } : {}
  )
}

output "pod_security_namespaces" {
  description = "Namespaces labeled by the Pod Security module."
  value       = var.enable_pod_security ? module.pod_security[0].namespaces : []
}

output "network_policy_baseline_namespaces" {
  description = "Namespaces where baseline NetworkPolicies are created."
  value       = var.enable_network_policy_baseline ? keys(module.network_policy_baseline) : []
}

output "argocd_app_of_apps_name" {
  description = "Name of the Argo CD app-of-apps Application when enabled."
  value       = var.enable_argocd && var.argocd_enable_app_of_apps ? module.argocd[0].app_of_apps_name : null
}
