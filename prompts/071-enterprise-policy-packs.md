## Prompt 71 — Enterprise policy packs

```text
Add enterprise policy pack support.

Goal:
Allow organizations to enforce ClusterForge policies consistently.

Create:
- policies/packs/
  baseline/
  production/
  kubernetes-security/
  aws-security/

Docs:
- docs/policies.md

Policy types:
1. CLI built-in policies
2. OPA/Conftest policies
3. Checkov/Trivy rules
4. Terraform plan JSON risk checks

Baseline pack:
- no plaintext secrets in tfvars
- no auto-approve in prod
- plan file required for prod apply
- destroy blocked in prod

Production pack:
- remote backend required
- version pinning required
- provider constraints required
- module sources must use tags, not main
- public ingress requires explicit annotation/approval

Kubernetes security pack:
- namespace labels for pod security recommended
- network policy recommended
- privileged containers flagged
- LoadBalancer services flagged unless allowed

AWS security pack:
- public S3 blocked
- unrestricted security groups flagged
- unencrypted state bucket blocked
- IAM wildcard policy warning

CLI:
- cf policy list
- cf policy check <env> --pack baseline
- cf policy check <env> --pack production

Rules:
- Start with documentation and simple checks.
- Avoid false confidence.
- Clearly label advisory vs blocking policies.

Run:
- gofmt
- go test ./...
```

---
