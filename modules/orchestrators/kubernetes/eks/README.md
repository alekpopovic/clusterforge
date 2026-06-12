# orchestrators/kubernetes/eks

## Purpose

Creates AWS EKS cluster infrastructure with managed node groups. This module
manages the EKS control plane, IAM roles and policy attachments, managed node
groups, and selected EKS add-ons.

Provider configuration belongs in the root module. This module declares the
AWS provider requirement but does not configure the provider.

This module only creates cluster infrastructure. It does not deploy Kubernetes
applications, Helm charts, ingress controllers, GitOps tooling, or workloads.
Those belong in platform and workload modules after the root config wires
Kubernetes and Helm providers from the EKS outputs.

## Add-ons

The module can manage these EKS add-ons:

- `vpc-cni`
- `coredns`
- `kube-proxy`
- `aws-ebs-csi-driver`

## IRSA And OIDC Provider

IAM Roles for Service Accounts (IRSA) lets Kubernetes workloads assume AWS IAM
roles through projected service account tokens instead of node instance
permissions. When `enable_irsa = true`, this module creates an IAM OIDC
provider from the EKS cluster issuer URL.

The OIDC provider outputs can be used by other modules that create workload or
platform service account roles:

- `oidc_provider_arn`
- `oidc_provider_url`
- `oidc_issuer_hostpath`

## EBS CSI Driver

When `enable_ebs_csi_driver_addon = true`, the module manages the
`aws-ebs-csi-driver` EKS add-on.

By default, `create_ebs_csi_irsa_role = true` creates an IRSA role for the
`kube-system/ebs-csi-controller-sa` service account and attaches the AWS
managed `AmazonEBSCSIDriverPolicy`. The role ARN is passed to the EKS add-on
with `service_account_role_arn`.

Set `create_ebs_csi_irsa_role = false` only when the service account role is
managed outside this module.

## Usage With Network Module

```hcl
provider "aws" {
  region = "us-east-1"
}

module "network" {
  source = "../../../modules/cloud/aws/network"

  name               = "clusterforge-dev"
  environment        = "dev"
  cidr_block         = "10.40.0.0/16"
  availability_zones = ["us-east-1a", "us-east-1b"]

  public_subnet_cidrs  = ["10.40.0.0/20", "10.40.16.0/20"]
  private_subnet_cidrs = ["10.40.128.0/20", "10.40.144.0/20"]

  private_subnet_tags = {
    "kubernetes.io/cluster/clusterforge-dev" = "shared"
  }
}

module "eks" {
  source = "../../../modules/orchestrators/kubernetes/eks"

  name        = "clusterforge-dev"
  environment = "dev"
  vpc_id      = module.network.vpc_id
  subnet_ids  = module.network.private_subnet_ids

  node_groups = {
    default = {
      instance_types = ["t3.medium"]
      min_size       = 1
      desired_size   = 2
      max_size       = 4
    }
  }
}
```

## EBS CSI IRSA Example

```hcl
module "eks" {
  source = "../../../modules/orchestrators/kubernetes/eks"

  name        = "clusterforge-dev"
  environment = "dev"
  vpc_id      = module.network.vpc_id
  subnet_ids  = module.network.private_subnet_ids

  enable_irsa                 = true
  enable_ebs_csi_driver_addon = true
  create_ebs_csi_irsa_role    = true
}
```

## Configuring Kubernetes And Helm Providers In Roots

Configure Kubernetes and Helm providers in the root module after creating the
cluster. A typical root uses the EKS outputs plus an AWS auth token data source:

```hcl
data "aws_eks_cluster_auth" "this" {
  name = module.eks.cluster_name
}

provider "kubernetes" {
  host                   = module.eks.cluster_endpoint
  cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
  token                  = data.aws_eks_cluster_auth.this.token
}

provider "helm" {
  kubernetes = {
    host                   = module.eks.cluster_endpoint
    cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
    token                  = data.aws_eks_cluster_auth.this.token
  }
}
```

## Inputs

| Name | Type | Default | Description |
| --- | --- | --- | --- |
| `name` | `string` | n/a | EKS cluster name. |
| `environment` | `string` | n/a | Environment name for tagging and resource names. |
| `kubernetes_version` | `string` | `"1.30"` | Kubernetes version for the EKS control plane. |
| `vpc_id` | `string` | n/a | VPC ID where EKS will run. |
| `subnet_ids` | `list(string)` | n/a | Subnet IDs for the EKS control plane and default node groups. |
| `endpoint_public_access` | `bool` | `true` | Whether the EKS API endpoint is publicly reachable. |
| `endpoint_private_access` | `bool` | `true` | Whether the EKS API endpoint is reachable from inside the VPC. |
| `public_access_cidrs` | `list(string)` | `["0.0.0.0/0"]` | CIDR blocks allowed to reach the public EKS API endpoint. |
| `enabled_cluster_log_types` | `list(string)` | `["api", "audit", "authenticator"]` | EKS control plane log types to enable. |
| `tags` | `map(string)` | `{}` | Tags applied to supported AWS resources. |
| `node_groups` | `map(object)` | `{ default = {} }` | Managed node group definitions. |
| `enable_vpc_cni_addon` | `bool` | `true` | Whether to manage the VPC CNI add-on. |
| `enable_coredns_addon` | `bool` | `true` | Whether to manage the CoreDNS add-on. |
| `enable_kube_proxy_addon` | `bool` | `true` | Whether to manage the kube-proxy add-on. |
| `enable_irsa` | `bool` | `true` | Whether to create an IAM OIDC provider for IRSA. |
| `enable_ebs_csi_driver_addon` | `bool` | `false` | Whether to manage the EBS CSI Driver add-on. |
| `create_ebs_csi_irsa_role` | `bool` | `true` | Whether to create an IRSA role for the EBS CSI controller service account. |

## Outputs

| Name | Description |
| --- | --- |
| `cluster_name` | EKS cluster name. |
| `cluster_arn` | EKS cluster ARN. |
| `cluster_endpoint` | EKS Kubernetes API endpoint. |
| `cluster_certificate_authority_data` | Base64-encoded EKS cluster certificate authority data. |
| `cluster_oidc_issuer_url` | EKS cluster OIDC issuer URL. |
| `oidc_provider_arn` | IAM OIDC provider ARN when IRSA is enabled. |
| `oidc_provider_url` | EKS OIDC provider issuer URL. |
| `oidc_issuer_hostpath` | EKS OIDC issuer URL without the https:// prefix. |
| `node_group_names` | Managed node group names. |
| `node_group_arns` | Managed node group ARNs. |
| `cluster_security_group_id` | EKS cluster security group ID created by EKS. |

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
