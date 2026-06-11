# aws-eks-minimal

Minimal AWS EKS example.

This root composes:

- `modules/core/tags`
- `modules/cloud/aws/network`
- `modules/orchestrators/kubernetes/eks`

No backend is configured. Terraform state is local unless you add a backend in
this root.

## Safe Local Validation

The example defaults to fake AWS credentials so contributors can run local
validation and no-refresh plans:

```bash
terraform init
terraform validate
terraform plan -refresh=false
```

Do not apply with fake credentials.

## Real AWS Use

To plan against real AWS credentials:

```bash
terraform plan -var='use_fake_credentials_for_plan=false'
```

Real EKS creation requires AWS credentials with permissions for VPC, IAM, EKS,
EC2, CloudWatch Logs, and related managed node group resources. Review IAM and
cost impact before applying.
