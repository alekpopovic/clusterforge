# ecs-fargate-app-with-alb

## Purpose

Example ECS/Fargate service behind an AWS Application Load Balancer.

This root composes:

- `modules/core/tags`
- `modules/cloud/aws/network`
- `modules/orchestrators/ecs/cluster`
- `modules/platform/ecs/alb`
- `modules/workloads/ecs/service`

## Usage

```bash
terraform init
terraform plan -refresh=false
```

The example includes fake AWS credentials for local syntax planning with
`-refresh=false`. Use real AWS credentials before applying.

## HTTPS

Leave `certificate_arn` empty for HTTP-only usage. To enable HTTPS, pass an ACM
certificate ARN from the root environment:

```bash
terraform plan -var='certificate_arn=arn:aws:acm:REGION:ACCOUNT:certificate/ID'
```

Do not commit real account IDs, certificate ARNs, or credentials into example
files.

## DNS

This example does not create Route53 records. Point DNS at the `alb_dns_name`
output using your DNS process or a future ClusterForge DNS composition.
