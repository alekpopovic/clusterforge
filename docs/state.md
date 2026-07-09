# State Helpers

Terraform state can contain sensitive values. ClusterForge exposes only
read-only helpers:

```bash
cf state list dev
cf state show dev module.network.aws_vpc.this
cf state pull dev --output /secure/path/state.json
```

`state pull` requires `--output` and refuses repository paths unless
`--allow-repo-output` is provided. Do not commit state JSON.

ClusterForge intentionally does not expose `state rm`, `state mv`, or direct
state editing. Use native Terraform/OpenTofu with reviewed runbooks for those
rare operations.
