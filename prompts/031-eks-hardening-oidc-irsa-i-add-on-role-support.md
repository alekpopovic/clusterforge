## Prompt 31 — EKS hardening: OIDC, IRSA i add-on role support

```text
Harden the EKS Terraform module for production-style IAM and add-on integration.

Target module:
- modules/orchestrators/kubernetes/eks

Goal:
Add OIDC provider and IRSA support foundations.

Tasks:
1. Add optional creation of IAM OIDC provider:
   Input:
   - enable_irsa: bool, default true

2. Output:
   - oidc_provider_arn
   - oidc_provider_url
   - oidc_issuer_hostpath

3. Add reusable IRSA helper module:
   modules/cloud/aws/irsa-role

IRSA module inputs:
- name: string
- environment: string
- oidc_provider_arn: string
- oidc_provider_url: string
- namespace: string
- service_account_name: string
- policy_arns: list(string), default []
- inline_policies: map(string), default {}
- tags: map(string), default {}

IRSA module resources:
- aws_iam_role
- aws_iam_role_policy_attachment
- aws_iam_role_policy for inline policies

IRSA module outputs:
- role_arn
- role_name

4. EBS CSI driver support:
   Add optional support for EBS CSI driver IAM role.
   Inputs:
   - enable_ebs_csi_driver_addon: bool, default false
   - create_ebs_csi_irsa_role: bool, default true

   When enabled:
   - Create IRSA role for kube-system/ebs-csi-controller-sa.
   - Attach AWS managed policy AmazonEBSCSIDriverPolicy.
   - Configure aws_eks_addon service_account_role_arn.

5. Update README:
   - Explain IRSA.
   - Explain OIDC provider.
   - Explain EBS CSI add-on.
   - Show example.

6. Update examples/aws-eks-minimal.

Rules:
- Do not configure providers in child modules.
- Do not hardcode AWS account ID.
- Keep backward compatibility where practical.
- Avoid breaking existing inputs.
- Use explicit depends_on only where needed.

Run:
- terraform fmt -recursive
- terraform validate where possible

Final response:
- Summarize EKS hardening changes.
- List new inputs/outputs.
- Mention validation result.
```

---
