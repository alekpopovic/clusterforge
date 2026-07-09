# cloud/aws/external-dns-irsa

## Purpose

Creates an IAM role for ExternalDNS on EKS using IRSA and an inline Route53
policy. The policy is restricted to the supplied hosted zone IDs by default.

## Status

Implemented.

## Usage

```hcl
module "external_dns_irsa" {
  source = "../../../modules/cloud/aws/external-dns-irsa"

  name              = "clusterforge-prod-external-dns"
  environment       = "prod"
  oidc_provider_arn = module.eks.oidc_provider_arn
  oidc_provider_url = module.eks.oidc_provider_url
  hosted_zone_ids   = ["ZREPLACEEXAMPLE"]
  policy_mode       = "upsert-only"
}
```

Use `policy_mode = "sync"` only when ExternalDNS is intentionally allowed to
delete records. `upsert-only` avoids deletes but can leave stale records.

Set `allow_all_hosted_zones = true` only for explicitly reviewed non-production
or shared DNS automation cases. The default requires `hosted_zone_ids`.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
