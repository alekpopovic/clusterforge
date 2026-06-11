# ecs-cluster-minimal

Minimal example for `modules/orchestrators/ecs/cluster`.

This example defaults to fake AWS credentials so contributors can run a local
no-refresh plan without real AWS credentials:

```bash
terraform init
terraform plan -refresh=false
```

Do not apply with fake credentials. To use real AWS credentials, set:

```bash
terraform plan -var='use_fake_credentials_for_plan=false'
```
