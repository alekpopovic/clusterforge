## Prompt 8 — Kubernetes generic provider/bootstrap root primer

```text
Create a live/dev/aws-eks root environment example that composes the AWS network and EKS modules.

Files:
live/dev/aws-eks/
  versions.tf
  backend.tf
  providers.tf
  main.tf
  variables.tf
  outputs.tf
  terraform.tfvars.example
  README.md

Requirements:
- Root module configures providers.
- Use hashicorp/aws, hashicorp/kubernetes, hashicorp/helm.
- Use data sources needed to configure Kubernetes and Helm providers from the EKS cluster.
- Do not commit real credentials.
- backend.tf should include a commented S3 backend example, not hardcoded real bucket.
- terraform.tfvars.example should contain safe example values, not secrets.

main.tf:
- Use modules/core/tags.
- Use modules/cloud/aws/network.
- Use modules/orchestrators/kubernetes/eks.
- Define one default node group.

providers.tf:
- Configure AWS using var.region.
- Configure Kubernetes provider using EKS endpoint, CA data, and token.
- Configure Helm provider using the Kubernetes connection.

README:
- Include commands:
  terraform init
  terraform plan
  terraform apply
- Explain how to copy terraform.tfvars.example to terraform.tfvars.
- Explain production warning: do not apply directly without review.

Run terraform fmt -recursive.
```

---
