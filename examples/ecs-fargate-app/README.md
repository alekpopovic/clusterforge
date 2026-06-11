# ecs-fargate-app

Example AWS ECS/Fargate app root.

This root composes:

- `modules/core/tags`
- `modules/cloud/aws/network`
- `modules/orchestrators/ecs/cluster`
- `modules/workloads/ecs/service`

It also creates a minimal service security group in the example root because
security group rules are environment-specific.

## Safe Local Validation

The example defaults to fake AWS credentials so contributors can run local
validation and no-refresh plans:

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

Real apply creates networking, an ECS cluster, IAM roles from the service
module, a CloudWatch log group, and an ECS service. Review security group
rules, NAT gateway cost, and task execution permissions before applying.

For production, you would usually add an ALB target group and pass it through
the ECS service module `load_balancer` input.
