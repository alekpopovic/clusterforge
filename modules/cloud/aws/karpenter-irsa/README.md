# cloud/aws/karpenter-irsa

## Purpose

Creates the IAM role and controller policy used by Karpenter on EKS through
IRSA. The trust relationship is scoped to the configured namespace and service
account.

Provider configuration belongs in the root module. This module declares the AWS
provider requirement but does not configure the provider.

## Status

Implemented.

## Usage

```hcl
module "karpenter_irsa" {
  source = "../../../modules/cloud/aws/karpenter-irsa"

  cluster_name      = module.eks.cluster_name
  oidc_provider_arn = module.eks.oidc_provider_arn
  oidc_provider_url = module.eks.oidc_provider_url

  tags = module.tags.tags
}
```

## Notes

- This module creates controller permissions only.
- Karpenter nodes still need a node IAM role and instance profile referenced by
  an `EC2NodeClass`.
- The controller policy uses broad EC2 read permissions and constrained launch
  and termination permissions. Review and tighten for your organization before
  production.
- Chart, CRD, NodePool, and EC2NodeClass versions should be pinned and tested
  before production rollouts.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
