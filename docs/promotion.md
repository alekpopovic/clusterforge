# Environment Promotion

ClusterForge encourages Git-based promotion from dev to staging to prod.

```bash
cf promote plan --from dev --to staging
cf promote diff --from staging --to prod
cf promote diff --from staging --to prod --json
```

Promotion commands compare app manifests, environment config, and generated
Terraform files. They do not apply, copy secrets, or mutate production.

Use immutable image tags, reviewed pull requests, saved plans, and production
approval gates.
