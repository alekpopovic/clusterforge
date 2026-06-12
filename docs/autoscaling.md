---
title: Node Autoscaling
permalink: /autoscaling/
---

# Node Autoscaling

ClusterForge currently supports EKS node autoscaling through Karpenter.
Cluster Autoscaler may be added later as a simpler compatibility option for
teams that prefer Auto Scaling Group based node groups.

## Managed Node Groups And Karpenter

Managed node groups are still supported and should not be removed when enabling
Karpenter. Keep a small managed node group for cluster bootstrap, CoreDNS,
critical platform add-ons, and the Karpenter controller itself.

Karpenter can then provision workload capacity through `NodePool` and
`EC2NodeClass` custom resources. Those resources define instance requirements,
subnet and security group selection, AMI family, disruption behavior,
consolidation, and capacity limits.

## Modules

- `modules/cloud/aws/karpenter-irsa` creates the IAM role used by the Karpenter
  controller service account through EKS IRSA.
- `modules/platform/kubernetes/karpenter` installs the Karpenter Helm chart.
- `modules/platform/kubernetes/bootstrap` can call the Karpenter module when
  `enable_karpenter = true`.

## IAM Requirements

Karpenter needs permission to discover EC2 capacity options, launch and
terminate instances tagged for the cluster, and pass the node role referenced by
`EC2NodeClass`. Some EC2 read permissions are necessarily account-wide. Launch
and termination permissions should be constrained with cluster tags wherever
possible.

Before production use, review the generated IAM policy against the Karpenter
version you plan to run and narrow `iam:PassRole` to approved node role ARNs
where possible.

## Production Caution

- Pin Karpenter chart and CRD versions before production.
- Review Karpenter release notes before upgrades.
- Define `NodePool` and `EC2NodeClass` manifests in a reviewed environment or
  GitOps repository.
- Set conservative capacity limits and disruption policies.
- Consider interruption handling, SQS queues, and spot capacity behavior before
  enabling spot-heavy node pools.
- Test node provisioning, consolidation, and termination behavior in a sandbox
  cluster first.
