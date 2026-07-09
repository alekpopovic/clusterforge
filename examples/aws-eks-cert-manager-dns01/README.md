# aws-eks-cert-manager-dns01

Example cert-manager Route53 DNS01 setup for EKS using IRSA. Replace OIDC and
hosted zone placeholders before applying.

Run syntax validation:

```bash
terraform init
terraform validate
```

Do not store AWS credentials in Kubernetes Secrets. Use the ACME staging
endpoint before switching production certificate issuance.
