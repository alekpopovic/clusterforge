# aws-eks-external-dns

Example ExternalDNS installation for EKS with Route53 permissions through
IRSA. Replace the OIDC and hosted zone placeholders before applying.

Run syntax validation:

```bash
terraform init
terraform validate
```

Use `policy_mode = "sync"` only when DNS deletion is reviewed and intended.
Do not store AWS credentials in Helm values.
