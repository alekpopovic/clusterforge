## Prompt 96 — External DNS production hardening

```text
Harden ExternalDNS support.

Target:
- modules/platform/kubernetes/external-dns
- modules/cloud/aws/irsa-role or new external-dns-irsa module

Add AWS-specific helper:
- modules/cloud/aws/external-dns-irsa

Inputs:
- name
- environment
- oidc_provider_arn
- oidc_provider_url
- namespace default "external-dns"
- service_account_name default "external-dns"
- hosted_zone_ids list(string)
- policy_mode string default "sync"
- tags

Behavior:
- Create least-privilege IAM policy for Route53 changes.
- Restrict to provided hosted zone IDs when possible.
- Output role ARN.

ExternalDNS module:
- Support service account annotations.
- Support provider-specific values through values input.
- Provide AWS example.

Docs:
- Explain hosted zone restriction.
- Explain sync vs upsert-only policy.
- Explain danger of accidental DNS deletion.
- Explain TXT ownership ID.

Example:
- examples/aws-eks-external-dns

Rules:
- Do not allow wildcard all zones by default.
- If hosted_zone_ids is empty, fail or require explicit allow_all_zones=true.
- Do not store AWS credentials in Helm values.

Run:
- terraform fmt -recursive
```

---
