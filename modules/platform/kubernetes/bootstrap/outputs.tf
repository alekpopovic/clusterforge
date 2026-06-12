output "enabled_addons" {
  description = "Names of enabled platform add-ons."
  value = compact([
    var.enable_ingress_nginx ? "ingress-nginx" : "",
    var.enable_cert_manager ? "cert-manager" : "",
    var.enable_external_dns ? "external-dns" : "",
    var.enable_external_secrets ? "external-secrets" : "",
    var.enable_metrics_server ? "metrics-server" : "",
    var.enable_prometheus_stack ? "prometheus-stack" : "",
    var.enable_loki ? "loki" : "",
    var.enable_argocd ? "argocd" : ""
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
    var.enable_metrics_server ? { metrics_server = module.metrics_server[0].release_name } : {},
    var.enable_prometheus_stack ? { prometheus_stack = module.prometheus_stack[0].release_name } : {},
    var.enable_loki ? { loki = module.loki[0].release_name } : {},
    var.enable_argocd ? { argocd = module.argocd[0].release_name } : {}
  )
}
