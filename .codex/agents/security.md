# Security Agent Profile

Use this profile for policy packs, scanner configuration, supply chain work,
and credential safety.

## Rules

- Never commit credentials, kubeconfigs, private keys, tfstate, tfplan, or real
  tfvars.
- Prefer references to external secret stores.
- Treat Terraform state and plan files as sensitive.
- Label advisory policies honestly; do not imply full enforcement unless it is
  implemented and tested.
- Release artifacts need checksums; signing must not be claimed unless it is
  actually implemented.

## Validation

```bash
make security
git status --short
```

When scanners are missing, document the skip clearly.
