# TLS With cert-manager

ClusterForge supports installing cert-manager and creating Issuer or
ClusterIssuer resources. Certificate automation is powerful, but production DNS
permissions must be narrow and reviewed.

## HTTP01 vs DNS01

HTTP01 validates a domain by serving a challenge through ingress. It is simple
for public HTTP services, but it requires working ingress and public reachability.

DNS01 validates by creating `_acme-challenge` DNS records. It works for private
services and wildcard certificates, but it requires DNS provider permissions.

## Route53 DNS01 On EKS

Prefer IRSA instead of static AWS credentials:

```hcl
module "cert_manager_route53_irsa" {
  source = "../modules/cloud/aws/cert-manager-route53-irsa"

  name              = "clusterforge-prod-cert-manager"
  environment       = "prod"
  oidc_provider_arn = module.eks.oidc_provider_arn
  oidc_provider_url = module.eks.oidc_provider_url
  hosted_zone_ids   = ["ZREPLACEEXAMPLE"]
}

module "cert_manager" {
  source = "../modules/platform/kubernetes/cert-manager"

  service_account_annotations = {
    "eks.amazonaws.com/role-arn" = module.cert_manager_route53_irsa.role_arn
  }
}
```

## Production Warnings

- Restrict Route53 permissions to known hosted zones.
- Do not store AWS access keys in Kubernetes Secrets for DNS01.
- Use staging ACME endpoints before production issuance changes.
- Review wildcard certificate scope before issuing.
- Test renewal before relying on automation.
