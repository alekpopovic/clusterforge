# cloud/aws/ecr

## Purpose

Creates AWS ECR repositories for application container images. This module
manages registry repositories only; it does not build, scan, sign, or push
Docker images.

## Status

Implemented.

## Basic App Repository

```hcl
module "ecr" {
  source = "../../../modules/cloud/aws/ecr"

  repositories = {
    api = {}
  }

  tags = {
    Project = "clusterforge"
  }
}
```

Repositories default to immutable tags and scan-on-push enabled. Immutable tags
are recommended so a deployed tag cannot be silently replaced after rollout.

## Lifecycle Policy Example

```hcl
module "ecr" {
  source = "../../../modules/cloud/aws/ecr"

  repositories = {
    api = {
      lifecycle_policy_json = jsonencode({
        rules = [
          {
            rulePriority = 1
            description  = "Keep the last 30 images."
            selection = {
              tagStatus   = "any"
              countType   = "imageCountMoreThan"
              countNumber = 30
            }
            action = {
              type = "expire"
            }
          }
        ]
      })
    }
  }
}
```

## KMS Encryption Example

```hcl
module "ecr_key" {
  source = "../../../modules/cloud/aws/kms-key"

  name        = "clusterforge-prod-ecr"
  environment = "prod"
  alias_name  = "clusterforge-prod-ecr"
}

module "ecr" {
  source = "../../../modules/cloud/aws/ecr"

  repositories = {
    api = {
      encryption_type = "KMS"
      kms_key_arn     = module.ecr_key.key_arn
    }
  }
}
```

## CI Image Push Flow

A typical CI pipeline builds an image, scans it, signs it when that control is
enabled, logs in to ECR, and pushes a versioned tag or digest. Avoid `latest`
for production workloads. ClusterForge workload manifests should reference a
specific version tag or digest.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
