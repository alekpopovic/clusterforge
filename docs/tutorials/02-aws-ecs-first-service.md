# Tutorial 02: AWS ECS First Service

## Goal

Generate an AWS ECS environment and understand the first service path.

## Prerequisites

- AWS CLI configured outside the repository
- `terraform` or `tofu`
- ClusterForge `cf`

This tutorial can create billable AWS resources if applied.

## Commands

```bash
cf project init ecs-demo
cf env create dev-ecs --cloud aws --orchestrator ecs --region us-east-2
cf generate dev-ecs
cp live/dev-ecs/aws-ecs/terraform.tfvars.example live/dev-ecs/aws-ecs/terraform.tfvars
cf init dev-ecs
cf plan dev-ecs --out .cf/plans/dev-ecs.tfplan --risk-summary
```

Apply only after reviewing the plan:

```bash
cf apply dev-ecs --plan-file .cf/plans/dev-ecs.tfplan
```

## Generated Files

- `clusterforge.yaml`
- `live/dev-ecs/aws-ecs/*.tf`
- `.cf/plans/dev-ecs.tfplan`

## What Terraform Creates

- VPC and subnets
- optional NAT Gateway
- ECS cluster
- CloudWatch settings

## Validate

```bash
aws ecs describe-clusters --clusters <cluster-name> --region us-east-2
```

## Cleanup

```bash
cf destroy dev-ecs --allow-destroy
```

## Troubleshooting

Check AWS identity, region, quotas, and stuck load balancers or ENIs. Do not
commit `terraform.tfvars`, state, or plan files.
