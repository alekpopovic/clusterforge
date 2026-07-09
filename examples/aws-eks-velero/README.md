# aws-eks-velero

Example Velero backup storage and Helm installation for an existing EKS
cluster. This example does not create the EKS cluster or the Velero IRSA role.

Run syntax validation:

```bash
terraform init
terraform validate
```

Before applying, create and review an IAM role for the Velero service account,
then pass its ARN with `velero_service_account_role_arn`. Do not store static
cloud credentials in Terraform values.
