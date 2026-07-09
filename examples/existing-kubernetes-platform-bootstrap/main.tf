module "bootstrap" {
  source = "../../modules/platform/kubernetes/bootstrap"

  enable_ingress_nginx    = false
  enable_cert_manager     = false
  enable_metrics_server   = true
  enable_external_dns     = false
  enable_prometheus_stack = false
  enable_loki             = false
  enable_argocd           = false
}
