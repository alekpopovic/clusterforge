# Policies

ClusterForge policy packs are a mix of built-in CLI checks, OPA/Conftest,
Checkov/Trivy guidance, and Terraform plan JSON risk checks.

```bash
cf policy list
cf policy check dev --pack baseline
cf policy check prod --pack production --plan-file .cf/plans/prod.tfplan
```

Packs:

- baseline: no plaintext secrets in tfvars, no auto-approve in prod, plan file
  required for prod apply, destroy blocked in prod
- production: remote backend, version pinning, provider constraints, tagged
  module sources, public ingress approval
- kubernetes-security: pod security labels, network policy, privileged
  container review, LoadBalancer review
- aws-security: public S3, unrestricted security groups, unencrypted state, IAM
  wildcard warnings

Some checks are advisory until dedicated OPA/Checkov rules are implemented.
