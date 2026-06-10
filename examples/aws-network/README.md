# aws-network

Example root module for `modules/cloud/aws/network`.

Provider configuration is intentionally declared in this root module, not in
the reusable network module.

This example defaults to fake AWS credentials so contributors can run a local
no-refresh plan without real AWS credentials:

```bash
terraform plan -refresh=false
```

Do not apply with fake credentials. To use real AWS credentials, set:

```bash
terraform plan -var='use_fake_credentials_for_plan=false'
```

```bash
terraform init
terraform plan
```
