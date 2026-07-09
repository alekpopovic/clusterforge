## Prompt 99 — AWS ElastiCache Redis module

```text
Implement AWS ElastiCache Redis module.

Path:
- modules/cloud/aws/elasticache-redis

Purpose:
Create Redis-compatible cache for container workloads.

Inputs:
- name
- environment
- vpc_id
- subnet_ids
- allowed_security_group_ids
- node_type
- engine_version
- num_cache_nodes default 1
- automatic_failover_enabled default false
- multi_az_enabled default false
- at_rest_encryption_enabled default true
- transit_encryption_enabled default true
- auth_token_secret_arn default ""
- parameter_group_name default ""
- tags

Resources:
- aws_elasticache_subnet_group
- aws_security_group
- aws_elasticache_replication_group or cluster depending on selected design

Outputs:
- primary_endpoint_address
- reader_endpoint_address if available
- port
- security_group_id

README:
- Basic Redis example.
- EKS/ECS connection example.
- Secret handling.
- Production HA notes.
- Cost notes.

Example:
- examples/aws-elasticache-redis

Rules:
- Do not output auth token.
- Do not make Redis public.
- Encryption enabled by default.
- Keep MVP simple.

Run:
- terraform fmt -recursive
```

---
