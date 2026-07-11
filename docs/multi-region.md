# Multi-region strategy

Region aliases are metadata; each generated environment still has its own
region and Terraform state. Use active/passive for controlled promotion, or
active/active only when applications and data stores support concurrency.

DNS health checks can support failover, but ClusterForge performs no automatic
failover or cross-region data replication. Define and test workload-specific
RTO/RPO, restore, DNS, and rollback procedures. Replication lag, split brain,
backups, egress, duplicate capacity, and observability add cost.

Keep state isolated by environment and region. A successful infrastructure
plan does not prove data is replicated or recoverable. Use `cf region list`,
`cf region show primary`, and `cf fleet status --region eu-central-1`.
