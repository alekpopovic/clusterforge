# aws-network

Runnable example for `modules/cloud/aws/network`.

This root composes:

- `modules/core/tags`
- `modules/cloud/aws/network`

Provider configuration is declared in this example root, not in the reusable
network module.

## Safe Local Validation

The example defaults to fake AWS credentials so contributors can run local
formatting, validation, and no-refresh plans without real cloud access:

```bash
terraform init
terraform validate
terraform plan -refresh=false
```

Do not apply with fake credentials.

## Real AWS Use

To plan against real AWS credentials:

```bash
terraform plan -var='use_fake_credentials_for_plan=false'
```

Review the generated plan carefully before applying. This example creates VPC,
subnet, route table, internet gateway, and NAT gateway resources.
