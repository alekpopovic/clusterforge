# platform/kubernetes/external-dns

Installs external-dns with Helm.

This module assumes Kubernetes and Helm providers are configured in the root
module. Configure DNS provider credentials outside this module. For AWS on EKS,
prefer IRSA with `modules/cloud/aws/external-dns-irsa`.

## Usage

```hcl
module "external_dns" {
  source = "../../../modules/platform/kubernetes/external-dns"

  namespace     = "external-dns"
  chart_version = "1.14.5"
}
```

## AWS Route53 Example

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

module "external_dns" {
  source = "../../../modules/platform/kubernetes/external-dns"

  service_account_annotations = {
    "eks.amazonaws.com/role-arn" = module.external_dns_irsa.role_arn
  }

  values = [
    yamlencode({
      provider = "aws"
      policy   = "upsert-only"
      txtOwnerId = "clusterforge-prod"
    })
  ]
}
```

Restrict ExternalDNS to specific hosted zone IDs whenever possible. `sync`
mode can delete DNS records; `upsert-only` is safer but can leave stale records.
Set a stable TXT ownership ID so multiple ExternalDNS installations do not
fight over the same records. Do not store AWS credentials in Helm values.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
