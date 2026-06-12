# platform/kubernetes/karpenter

## Purpose

Installs Karpenter on EKS with Helm. Karpenter can launch and terminate nodes
based on pending Kubernetes workloads and AWS capacity availability.

This module installs the controller only. It does not create `NodePool` or
`EC2NodeClass` manifests because those APIs and production choices should be
version-pinned and reviewed per environment.

Provider configuration belongs in the root module. This module declares Helm
and Kubernetes provider requirements but does not configure providers.

## Status

Implemented.

## Usage

```hcl
module "karpenter" {
  source = "../../../modules/platform/kubernetes/karpenter"

  cluster_name             = module.eks.cluster_name
  cluster_endpoint         = module.eks.cluster_endpoint
  service_account_role_arn = module.karpenter_irsa.role_arn
  chart_version            = "1.0.0"
}
```

## NodePool And EC2NodeClass

Karpenter requires Karpenter CRDs plus `NodePool` and `EC2NodeClass` resources
to actually provision nodes. This module intentionally does not manage those
manifests yet. Add them through GitOps or a focused environment module after
pinning the Karpenter API version you intend to run.

## Managed Node Groups And Karpenter

Keep at least one small managed node group or another bootstrap capacity path
for system add-ons and Karpenter itself. Karpenter can then manage additional
workload capacity.

## Notes

- Karpenter is not enabled by default in ClusterForge.
- Pin chart and CRD versions before production.
- Review node IAM role, subnet, security group, and interruption handling
  design before production use.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
