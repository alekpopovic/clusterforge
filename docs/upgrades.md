# Upgrades

ClusterForge can migrate project configuration safely:

```bash
cf upgrade check
cf upgrade plan
cf upgrade apply --yes
```

The first migration adds missing `clusterforge_version`, default policies,
backend skeletons, and normalized environment paths. `apply` creates backups in
`.cf/backups/<timestamp>/` before writing.

Upgrade commands never modify Terraform state and never run `terraform apply`.
