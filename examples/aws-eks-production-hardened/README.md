# aws-eks-production-hardened

Production-oriented EKS example with private API endpoint access, control plane
logs, secrets encryption, IRSA, and managed node group rollout controls.

The example defaults to fake AWS credentials for local validation only:

```bash
terraform init
terraform validate
terraform plan -refresh=false
```

Do not apply with fake credentials. For real AWS use, set:

```bash
terraform plan -var='use_fake_credentials_for_plan=false'
```

Private endpoint-only access means Terraform and Kubernetes operations must run
from a network path that can reach the VPC, such as VPN, Direct Connect, a
bastion, or a CI runner inside the VPC.
