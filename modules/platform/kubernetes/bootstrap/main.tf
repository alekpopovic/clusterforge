locals {
  namespace_prefix = trimspace(var.namespace_prefix)

  namespaces = {
    ingress_nginx    = local.namespace_prefix == "" ? "ingress-nginx" : "${local.namespace_prefix}-ingress-nginx"
    cert_manager     = local.namespace_prefix == "" ? "cert-manager" : "${local.namespace_prefix}-cert-manager"
    external_dns     = local.namespace_prefix == "" ? "external-dns" : "${local.namespace_prefix}-external-dns"
    external_secrets = local.namespace_prefix == "" ? "external-secrets" : "${local.namespace_prefix}-external-secrets"
    karpenter        = local.namespace_prefix == "" ? "karpenter" : "${local.namespace_prefix}-karpenter"
    metrics_server   = local.namespace_prefix == "" ? "metrics-server" : "${local.namespace_prefix}-metrics-server"
    prometheus_stack = local.namespace_prefix == "" ? "monitoring" : "${local.namespace_prefix}-monitoring"
    loki             = local.namespace_prefix == "" ? "logging" : "${local.namespace_prefix}-logging"
    log_agent        = local.namespace_prefix == "" ? "logging" : "${local.namespace_prefix}-logging"
    argocd           = local.namespace_prefix == "" ? "argocd" : "${local.namespace_prefix}-argocd"
    opentelemetry    = local.namespace_prefix == "" ? "observability" : "${local.namespace_prefix}-observability"
  }
}

module "ingress_nginx" {
  count = var.enable_ingress_nginx ? 1 : 0

  source = "../ingress-nginx"

  namespace = local.namespaces.ingress_nginx
  labels    = var.common_labels
}

module "cert_manager" {
  count = var.enable_cert_manager ? 1 : 0

  source = "../cert-manager"

  namespace = local.namespaces.cert_manager
  labels    = var.common_labels
}

module "external_dns" {
  count = var.enable_external_dns ? 1 : 0

  source = "../external-dns"

  namespace = local.namespaces.external_dns
  labels    = var.common_labels
}

module "external_secrets" {
  count = var.enable_external_secrets ? 1 : 0

  source = "../external-secrets"

  namespace = local.namespaces.external_secrets
  labels    = var.common_labels
}

module "pod_security" {
  count = var.enable_pod_security ? 1 : 0

  source = "../pod-security"

  namespaces = var.pod_security_namespaces
  labels     = var.common_labels
}

module "network_policy_baseline" {
  for_each = var.enable_network_policy_baseline ? toset(var.network_policy_namespaces) : toset([])

  source = "../network-policy-baseline"

  namespace            = each.value
  default_deny_ingress = var.network_policy_default_deny_ingress
  default_deny_egress  = var.network_policy_default_deny_egress
  allow_dns_egress     = var.network_policy_allow_dns_egress
  labels               = var.common_labels
}

module "karpenter" {
  count = var.enable_karpenter ? 1 : 0

  source = "../karpenter"

  namespace                = local.namespaces.karpenter
  chart_version            = var.karpenter_chart_version
  cluster_name             = var.karpenter_cluster_name
  cluster_endpoint         = var.karpenter_cluster_endpoint
  service_account_role_arn = var.karpenter_service_account_role_arn
  values                   = var.karpenter_values
  labels                   = var.common_labels
}

module "metrics_server" {
  count = var.enable_metrics_server ? 1 : 0

  source = "../metrics-server"

  namespace = local.namespaces.metrics_server
  labels    = var.common_labels
}

module "prometheus_stack" {
  count = var.enable_prometheus_stack ? 1 : 0

  source = "../prometheus-stack"

  namespace = local.namespaces.prometheus_stack
  labels    = var.common_labels
}

module "loki" {
  count = var.enable_loki ? 1 : 0

  source = "../loki"

  namespace = local.namespaces.loki
  labels    = var.common_labels
}

module "log_agent" {
  count = var.enable_log_agent ? 1 : 0

  source = "../alloy"

  namespace        = local.namespaces.log_agent
  chart_version    = var.log_agent_chart_version
  values           = var.log_agent_values
  labels           = var.common_labels
  create_namespace = !var.enable_loki
}

module "argocd" {
  count = var.enable_argocd ? 1 : 0

  source = "../argocd"

  namespace                         = local.namespaces.argocd
  labels                            = var.common_labels
  enable_app_of_apps                = var.argocd_enable_app_of_apps
  app_of_apps_name                  = var.argocd_app_of_apps_name
  app_of_apps_repo_url              = var.argocd_app_of_apps_repo_url
  app_of_apps_path                  = var.argocd_app_of_apps_path
  app_of_apps_revision              = var.argocd_app_of_apps_revision
  app_of_apps_destination_namespace = var.argocd_app_of_apps_destination_namespace
  app_of_apps_project               = var.argocd_app_of_apps_project
}

module "opentelemetry_collector" {
  count                       = var.enable_opentelemetry_collector ? 1 : 0
  source                      = "../opentelemetry-collector"
  namespace                   = local.namespaces.opentelemetry
  chart_version               = var.opentelemetry_collector_chart_version
  mode                        = var.opentelemetry_collector_mode
  values                      = var.opentelemetry_collector_values
  service_account_annotations = var.opentelemetry_collector_service_account_annotations
  labels                      = var.common_labels
}
