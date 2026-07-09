## Prompt 97 — Cert-manager Route53 DNS01 IAM module

```text
Add AWS IAM support for cert-manager Route53 DNS01.

Create module:
- modules/cloud/aws/cert-manager-route53-irsa

Inputs:
- name
- environment
- oidc_provider_arn
- oidc_provider_url
- namespace default "cert-manager"
- service_account_name default "cert-manager"
- hosted_zone_ids list(string)
- tags map(string)

Resources:
- IAM role for service account
- least-privilege IAM policy for Route53 DNS01 challenges

Outputs:
- role_arn
- role_name
- policy_arn or policy_json

Update:
- modules/platform/kubernetes/cert-manager
  - support service account annotations
- modules/platform/kubernetes/cert-manager-issuer
  - ensure DNS01 Route53 example is documented

Docs:
- docs/tls-cert-manager.md
- HTTP01 vs DNS01
- Route53 hosted zone permissions
- IRSA annotation example
- production warnings

Example:
- examples/aws-eks-cert-manager-dns01

Rules:
- Do not grant access to all hosted zones unless explicitly allowed.
- Do not store AWS credentials in Kubernetes secrets.
- Prefer IRSA.

Run:
- terraform fmt -recursive
```

---
