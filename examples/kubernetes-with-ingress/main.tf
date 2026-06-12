locals {
  issuer_name = "letsencrypt-prod"
}

module "cert_manager_issuer" {
  source = "../../modules/platform/kubernetes/cert-manager-issuer"

  name  = local.issuer_name
  kind  = "ClusterIssuer"
  email = var.acme_email

  solvers = [
    {
      http01_ingress_class = "nginx"
    }
  ]
}
