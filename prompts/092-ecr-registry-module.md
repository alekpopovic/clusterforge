## Prompt 92 — ECR registry module

```text
Implement AWS ECR registry module.

Path:
- modules/cloud/aws/ecr

Purpose:
Create ECR repositories for application images.

Inputs:
- repositories: map(object({
    image_tag_mutability = optional(string, "IMMUTABLE")
    scan_on_push = optional(bool, true)
    encryption_type = optional(string, "AES256")
    kms_key_arn = optional(string)
    lifecycle_policy_json = optional(string)
  }))
- tags: map(string), default {}

Resources:
- aws_ecr_repository
- aws_ecr_lifecycle_policy optional

Outputs:
- repository_urls
- repository_arns
- repository_names

README:
- Basic app repository example.
- Lifecycle policy example.
- KMS encryption example.
- Explain immutable tags recommendation.
- Explain CI image push flow.

Example:
- examples/aws-ecr-repositories

CLI:
- Optional command design only:
  cf registry add api
Do not implement CLI registry command unless straightforward.

Rules:
- Prefer immutable tags by default.
- Enable scan_on_push by default.
- Do not manage Docker image builds here.

Run:
- terraform fmt -recursive
```

---
