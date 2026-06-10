locals {
  namespace_prefix = trimspace(var.namespace_prefix)

  namespaces = {
    ingress_nginx    = local.namespace_prefix == "" ? "ingress-nginx" : "${local.namespace_prefix}-ingress-nginx"
    cert_manager     = local.namespace_prefix == "" ? "cert-manager" : "${local.namespace_prefix}-cert-manager"
    external_dns     = local.namespace_prefix == "" ? "external-dns" : "${local.namespace_prefix}-external-dns"
    metrics_server   = local.namespace_prefix == "" ? "metrics-server" : "${local.namespace_prefix}-metrics-server"
    prometheus_stack = local.namespace_prefix == "" ? "monitoring" : "${local.namespace_prefix}-monitoring"
    loki             = local.namespace_prefix == "" ? "loki" : "${local.namespace_prefix}-loki"
    argocd           = local.namespace_prefix == "" ? "argocd" : "${local.namespace_prefix}-argocd"
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

module "argocd" {
  count = var.enable_argocd ? 1 : 0

  source = "../argocd"

  namespace = local.namespaces.argocd
  labels    = var.common_labels
}
