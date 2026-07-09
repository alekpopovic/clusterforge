module "external_dns_irsa" {
  source = "../../modules/cloud/aws/external-dns-irsa"

  name              = "clusterforge-prod-external-dns"
  environment       = "prod"
  oidc_provider_arn = var.oidc_provider_arn
  oidc_provider_url = var.oidc_provider_url
  hosted_zone_ids   = var.hosted_zone_ids
  policy_mode       = "upsert-only"
}

module "external_dns" {
  source = "../../modules/platform/kubernetes/external-dns"

  service_account_annotations = {
    "eks.amazonaws.com/role-arn" = module.external_dns_irsa.role_arn
  }

  values = [
    yamlencode({
      provider   = "aws"
      policy     = "upsert-only"
      txtOwnerId = "clusterforge-prod"
    })
  ]
}
