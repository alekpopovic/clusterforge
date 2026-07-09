# platform/kubernetes/cert-manager

Installs cert-manager with Helm.

This module assumes Kubernetes and Helm providers are configured in the root
module. Review cert-manager CRD lifecycle before upgrades.

## Usage

```hcl
module "cert_manager" {
  source = "../../../modules/platform/kubernetes/cert-manager"

  namespace     = "cert-manager"
  chart_version = "v1.14.5"
}
```

## AWS Route53 DNS01 IRSA

```hcl
module "cert_manager_route53_irsa" {
  source = "../../../modules/cloud/aws/cert-manager-route53-irsa"

  name              = "clusterforge-prod-cert-manager"
  environment       = "prod"
  oidc_provider_arn = module.eks.oidc_provider_arn
  oidc_provider_url = module.eks.oidc_provider_url
  hosted_zone_ids   = ["ZREPLACEEXAMPLE"]
}

module "cert_manager" {
  source = "../../../modules/platform/kubernetes/cert-manager"

  service_account_annotations = {
    "eks.amazonaws.com/role-arn" = module.cert_manager_route53_irsa.role_arn
  }
}
```

Do not store AWS access keys in Kubernetes Secrets for DNS01. Prefer IRSA.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
