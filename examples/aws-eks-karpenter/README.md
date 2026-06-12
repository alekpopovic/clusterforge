# aws-eks-karpenter

## Purpose

Example EKS environment that composes the ClusterForge AWS network, EKS,
Karpenter IRSA, and Karpenter Helm modules.

## Usage

```bash
terraform init
terraform plan -refresh=false
```

The example includes fake AWS credentials for local syntax planning with
`-refresh=false`. Use real AWS credentials before applying.

## Production Notes

- Karpenter is not enabled by default in ClusterForge bootstrap.
- Keep a small managed node group for bootstrap and critical system add-ons.
- Pin `karpenter_chart_version` and review chart and CRD release notes before
  production use.
- This example installs the Karpenter controller only. Add reviewed
  `NodePool` and `EC2NodeClass` manifests for actual node provisioning.
- Do not store Git, cloud, or Kubernetes credentials in Terraform variables.

## NodePool And EC2NodeClass

Karpenter scheduling behavior is controlled by Kubernetes custom resources.
ClusterForge leaves those manifests to the environment or GitOps repository so
teams can review AMI family, subnet and security group selectors, disruption
budgets, consolidation behavior, capacity limits, and node IAM role choices.
