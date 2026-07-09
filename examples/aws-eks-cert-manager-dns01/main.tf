module "cert_manager_route53_irsa" {
  source = "../../modules/cloud/aws/cert-manager-route53-irsa"

  name              = "clusterforge-prod-cert-manager"
  environment       = "prod"
  oidc_provider_arn = var.oidc_provider_arn
  oidc_provider_url = var.oidc_provider_url
  hosted_zone_ids   = var.hosted_zone_ids
}

module "cert_manager" {
  source = "../../modules/platform/kubernetes/cert-manager"

  service_account_annotations = {
    "eks.amazonaws.com/role-arn" = module.cert_manager_route53_irsa.role_arn
  }
}

module "issuer" {
  source = "../../modules/platform/kubernetes/cert-manager-issuer"

  name  = "letsencrypt-prod"
  kind  = "ClusterIssuer"
  email = var.acme_email

  solvers = [
    {
      dns01_route53 = {
        region         = var.aws_region
        hosted_zone_id = var.hosted_zone_ids[0]
      }
    }
  ]

  depends_on = [module.cert_manager]
}
