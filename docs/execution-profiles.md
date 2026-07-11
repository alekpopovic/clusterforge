# Execution profiles

Execution profiles define reviewable defaults for Terraform/OpenTofu commands:

```yaml
execution_profiles:
  local:
    engine: terraform
    parallelism: 10
    refresh: true
    lock_timeout: 5m
  prod:
    engine: terraform
    parallelism: 3
    refresh: true
    lock_timeout: 20m
    input: false
    require_plan_file: true
```

Use `cf profile list`, `cf profile show prod`, `cf plan dev --profile local`,
and `cf apply prod --profile prod --plan-file reviewed.tfplan`.

Profiles can set engine, parallelism, refresh behavior for plans, lock timeout,
and interactive input. A production profile can require a saved plan even when
an environment name is unconventional. Profiles never add `-auto-approve` and
do not weaken the existing production confirmation or destroy protections.
