## Prompt 37 — Cert-manager ClusterIssuer module

```text
Add cert-manager issuer support.

Create module:
- modules/platform/kubernetes/cert-manager-issuer

Purpose:
Create cert-manager ClusterIssuer or Issuer resources after cert-manager is installed.

Inputs:
- name: string, default "letsencrypt-prod"
- kind: string, default "ClusterIssuer"
- namespace: string, default ""
- email: string
- server: string, default "https://acme-v02.api.letsencrypt.org/directory"
- private_key_secret_name: string, default "letsencrypt-prod"
- solvers: list(object({
    http01_ingress_class = optional(string)
    dns01_route53 = optional(object({
      region = string
      hosted_zone_id = optional(string)
    }))
  })), default []

Resources:
- kubernetes_manifest for Issuer or ClusterIssuer.

Validation:
- email required.
- kind must be Issuer or ClusterIssuer.
- namespace required when kind=Issuer.
- At least one solver required.

README:
- HTTP01 example with ingress-nginx.
- DNS01 Route53 example.
- Explain IAM requirements for DNS01.
- Explain dependency on cert-manager CRDs.

Update examples:
- examples/kubernetes-with-ingress
  Include cert-manager issuer example.

Rules:
- Do not store ACME account private key manually.
- Do not put DNS provider credentials in Terraform.
- Keep solver schema simple.

Run:
- terraform fmt -recursive
```

---
