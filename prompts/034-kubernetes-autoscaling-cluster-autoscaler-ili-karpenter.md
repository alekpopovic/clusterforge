## Prompt 34 — Kubernetes autoscaling: Cluster Autoscaler ili Karpenter

```text
Add Kubernetes node autoscaling support for EKS.

Goal:
Provide optional platform support for node autoscaling on EKS.

Evaluate and implement one of:
1. Karpenter
2. Cluster Autoscaler

Preferred first implementation:
- Karpenter for EKS, if module complexity stays manageable.

Create modules:
- modules/platform/kubernetes/karpenter
- modules/cloud/aws/karpenter-irsa

Requirements:
1. cloud/aws/karpenter-irsa:
   - Create IAM role and policies needed by Karpenter.
   - Use EKS OIDC provider.
   - Inputs:
     - cluster_name
     - oidc_provider_arn
     - oidc_provider_url
     - namespace default "karpenter"
     - service_account_name default "karpenter"
     - tags
   - Outputs:
     - role_arn
     - role_name

2. platform/kubernetes/karpenter:
   - Install Karpenter via Helm.
   - Inputs:
     - namespace default "karpenter"
     - chart_version
     - cluster_name
     - cluster_endpoint
     - service_account_role_arn
     - values
   - Outputs:
     - namespace
     - release_name

3. Optional provisioner/nodepool manifests:
   - If current Karpenter API version is known in codebase docs, implement a simple example manifest.
   - Otherwise add TODO and docs explaining user must provide NodePool/EC2NodeClass.

4. Update bootstrap module:
   - enable_karpenter: bool default false

5. Add example:
   - examples/aws-eks-karpenter

Docs:
- Explain relationship between managed node groups and Karpenter.
- Explain production caution.
- Explain IAM requirements.
- Explain that chart and CRD versions should be pinned.

Rules:
- Do not remove existing managed node group support.
- Do not make Karpenter default.
- Do not create overly broad IAM permissions without comments.
- Keep module focused.

Run:
- terraform fmt -recursive
- validation where possible

Final response:
- State whether Karpenter or Cluster Autoscaler was implemented.
- List resources/modules added.
- Mention production TODOs.
```

---
