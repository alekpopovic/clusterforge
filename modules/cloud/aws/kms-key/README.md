# cloud/aws/kms-key

## Purpose

Creates a reusable AWS KMS key and alias for encryption use cases such as EKS
secrets encryption, S3 Terraform state encryption, EBS encryption, and
application secrets encryption.

## Status

Implemented.

## Usage

```hcl
module "kms" {
  source = "../../../modules/cloud/aws/kms-key"

  name        = "clusterforge-prod-platform"
  environment = "prod"
  alias_name  = "clusterforge-prod-platform"

  tags = {
    Project = "clusterforge"
  }
}
```

## EKS Encryption Example

```hcl
module "eks_secrets_key" {
  source = "../../../modules/cloud/aws/kms-key"

  name        = "clusterforge-prod-eks-secrets"
  environment = "prod"
  alias_name  = "clusterforge-prod-eks-secrets"
  description = "KMS key for EKS Kubernetes secrets encryption."
}

module "eks" {
  source = "../../../modules/orchestrators/kubernetes/eks"

  enable_cluster_encryption = true
  kms_key_arn               = module.eks_secrets_key.key_arn
}
```

## S3 Backend Encryption Example

```hcl
module "state_key" {
  source = "../../../modules/cloud/aws/kms-key"

  name        = "clusterforge-prod-tfstate"
  environment = "bootstrap"
  alias_name  = "clusterforge-prod-tfstate"
  description = "KMS key for Terraform state bucket encryption."
}

module "tfstate_backend" {
  source = "../../../modules/cloud/aws/tfstate-backend"

  name                = "clusterforge-tfstate"
  environment         = "bootstrap"
  bucket_name         = "example-clusterforge-terraform-state"
  dynamodb_table_name = "example-clusterforge-terraform-locks"
  kms_key_arn         = module.state_key.key_arn
}
```

## Warnings

KMS key deletion has a recovery window between 7 and 30 days. Scheduling key
deletion can make encrypted data unrecoverable after the window expires.

Callers must grant IAM principals permission to use the key for their intended
service. This module does not create broad key policies by default; pass
`policy_json` only after reviewing the exact principals and actions required.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
