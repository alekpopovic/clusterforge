# aws-ecr-repositories

Example ECR repositories for ClusterForge application images. Repositories use
immutable tags and scan-on-push by default.

Run local validation:

```bash
terraform init
terraform validate
terraform plan -refresh=false
```

This example does not build or push Docker images.
