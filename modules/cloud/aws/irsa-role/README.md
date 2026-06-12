# cloud/aws/irsa-role

## Purpose

Creates an IAM role for Kubernetes service accounts using IAM Roles for Service
Accounts (IRSA). The module builds a web identity trust policy for one
namespace and service account, then attaches managed and inline IAM policies.

Provider configuration belongs in the root module. This module declares the AWS
provider requirement but does not configure the provider.

## Status

Implemented.

## Usage

```hcl
module "ebs_csi_irsa" {
  source = "../../../modules/cloud/aws/irsa-role"

  name                 = "clusterforge-dev-ebs-csi"
  environment          = "dev"
  oidc_provider_arn    = module.eks.oidc_provider_arn
  oidc_provider_url    = module.eks.oidc_provider_url
  namespace            = "kube-system"
  service_account_name = "ebs-csi-controller-sa"

  policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AmazonEBSCSIDriverPolicy"
  ]
}
```

## Notes

- `oidc_provider_url` may include or omit `https://`.
- Inline policies must be valid JSON policy documents.
- Do not put secret values in inline policies.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
