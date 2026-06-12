# platform/kubernetes/cert-manager-issuer

## Purpose

Creates cert-manager `Issuer` or `ClusterIssuer` resources after cert-manager
and its CRDs are installed.

Provider configuration belongs in the root module. This module declares the
Kubernetes provider requirement but does not configure the provider.

## Status

Implemented.

## HTTP01 With ingress-nginx

```hcl
module "letsencrypt" {
  source = "../../../modules/platform/kubernetes/cert-manager-issuer"

  name  = "letsencrypt-prod"
  kind  = "ClusterIssuer"
  email = "platform@example.com"

  solvers = [
    {
      http01_ingress_class = "nginx"
    }
  ]
}
```

## DNS01 With Route53

```hcl
module "letsencrypt_dns" {
  source = "../../../modules/platform/kubernetes/cert-manager-issuer"

  name  = "letsencrypt-prod"
  kind  = "ClusterIssuer"
  email = "platform@example.com"

  solvers = [
    {
      dns01_route53 = {
        region         = "us-east-1"
        hosted_zone_id = "Z000000000000EXAMPLE"
      }
    }
  ]
}
```

## IAM Requirements For DNS01

Route53 DNS01 challenges require cert-manager to have IAM permissions to read
and change challenge records in the hosted zone. Prefer IRSA on EKS and keep
cloud credentials out of Terraform variables and Kubernetes manifests.

Do not put DNS provider access keys in Terraform. Manage workload identity or
external secret references separately.

## cert-manager CRD Dependency

This module uses `kubernetes_manifest` for cert-manager custom resources. Run
it only after cert-manager and its CRDs are installed. In root compositions,
use `depends_on = [module.cert_manager]` when Terraform installs cert-manager.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
