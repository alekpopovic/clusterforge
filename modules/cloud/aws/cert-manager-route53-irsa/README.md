# cloud/aws/cert-manager-route53-irsa

## Purpose

Creates an IAM role for cert-manager on EKS using IRSA and a Route53 DNS01
challenge policy. The policy is restricted to the supplied hosted zone IDs by
default and to `_acme-challenge.*` record changes.

## Status

Implemented.

## Usage

```hcl
module "cert_manager_route53_irsa" {
  source = "../../../modules/cloud/aws/cert-manager-route53-irsa"

  name              = "clusterforge-prod-cert-manager"
  environment       = "prod"
  oidc_provider_arn = module.eks.oidc_provider_arn
  oidc_provider_url = module.eks.oidc_provider_url
  hosted_zone_ids   = ["ZREPLACEEXAMPLE"]
}
```

Keep `allow_all_hosted_zones = false` for production. Prefer one role per
cluster and an explicit hosted zone list.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
