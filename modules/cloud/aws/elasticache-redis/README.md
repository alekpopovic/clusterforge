# cloud/aws/elasticache-redis

## Purpose

Creates a private Redis-compatible AWS ElastiCache replication group for
container workloads. Encryption in transit and at rest are enabled by default.

## Status

Implemented.

## Basic Example

```hcl
module "redis" {
  source = "../../../modules/cloud/aws/elasticache-redis"

  name                       = "clusterforge-prod-cache"
  environment                = "prod"
  vpc_id                     = module.network.vpc_id
  subnet_ids                 = module.network.private_subnet_ids
  allowed_security_group_ids = [module.api.security_group_id]
}
```

## EKS And ECS Connection Pattern

Allow the workload security group to connect to this module security group on
port `6379`. Store any auth token in AWS Secrets Manager or another approved
secret manager, then sync it into Kubernetes or ECS task secrets. This module
does not output auth tokens.

## Production HA Notes

Set `num_cache_nodes >= 2`, `automatic_failover_enabled = true`, and
`multi_az_enabled = true` for higher availability. That increases cost.

## Cost Notes

ElastiCache charges for node hours, backup storage where enabled outside this
module, and data transfer. Larger node types and Multi-AZ increase cost.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
