## Prompt 7 — EKS orchestrator modul

```text
Implement modules/orchestrators/kubernetes/eks.

Purpose:
Provision an AWS EKS cluster with managed node groups.

Provider:
- Use hashicorp/aws.
- Do not configure provider inside this module.

Inputs:
- name: string
- environment: string
- kubernetes_version: string, default "1.30"
- vpc_id: string
- subnet_ids: list(string)
- endpoint_public_access: bool, default true
- endpoint_private_access: bool, default true
- public_access_cidrs: list(string), default ["0.0.0.0/0"]
- enabled_cluster_log_types: list(string), default ["api", "audit", "authenticator"]
- tags: map(string), default {}

Node groups:
Create variable node_groups as map(object({
  subnet_ids: optional(list(string))
  instance_types: optional(list(string), ["t3.medium"])
  capacity_type: optional(string, "ON_DEMAND")
  min_size: optional(number, 1)
  desired_size: optional(number, 2)
  max_size: optional(number, 4)
  disk_size: optional(number, 50)
  labels: optional(map(string), {})
  taints: optional(list(object({
    key = string
    value = optional(string)
    effect = string
  })), [])
}))

Validation:
- name/environment/vpc_id non-empty.
- subnet_ids length > 0.
- node group min <= desired <= max.
- capacity_type must be ON_DEMAND or SPOT.

Resources:
- IAM role for EKS cluster
- IAM role policy attachments for cluster
- aws_eks_cluster
- IAM role for node groups
- IAM role policy attachments for nodes
- aws_eks_node_group for each node group

Add-ons:
Optional inputs:
- enable_vpc_cni_addon: bool default true
- enable_coredns_addon: bool default true
- enable_kube_proxy_addon: bool default true
- enable_ebs_csi_driver_addon: bool default false

Resources:
- aws_eks_addon for enabled add-ons
- Include service account role for EBS CSI only if needed, or leave TODO if too much for first implementation.

Outputs:
- cluster_name
- cluster_arn
- cluster_endpoint
- cluster_certificate_authority_data
- cluster_oidc_issuer_url
- node_group_names
- node_group_arns
- cluster_security_group_id

README:
- Explain the module.
- Show usage with network module.
- Explain how root module should configure kubernetes/helm providers using EKS outputs.
- Explain that this module only creates cluster infrastructure, not app deployments.

Create examples/aws-eks-minimal that composes:
- modules/cloud/aws/network
- modules/orchestrators/kubernetes/eks

Run terraform fmt -recursive.
```

---
