# aws-elasticache-redis

Example private ElastiCache Redis replication group.

Run local validation:

```bash
terraform init
terraform validate
terraform plan -refresh=false
```

Do not apply with fake credentials. Review Redis node cost, HA settings, and
secret handling before using this in a real account.
