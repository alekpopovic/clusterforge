# Drift Detection

Drift means the real infrastructure no longer matches Terraform/OpenTofu state
and configuration. ClusterForge detects drift with:

```bash
cf drift check dev
cf drift check dev --stack network
cf drift check dev --json
```

The command runs `plan -detailed-exitcode` and never applies changes.

- exit `0`: no drift
- exit `2`: drift or changes detected
- exit `1`: error

Limitations: provider refresh behavior, ignored lifecycle changes, missing
credentials, and out-of-band resources can affect results. Production drift
must be reviewed by humans before remediation.

Scheduled CI drift checks should be opt-in and configured with environment
credentials outside the repository.
