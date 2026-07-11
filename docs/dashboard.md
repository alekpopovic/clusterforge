# Dashboard proposal

The proposed ClusterForge dashboard is a read-only operational index, not a
control plane. Its first version reads a local JSON snapshot generated with:

```bash
cf dashboard export --output dashboard-data.json
```

The proposed document contains `schema_version`, `generated_at`, source commit,
and redacted inventory, health/drift, policy, cost, audit, runbook and release
sections. Missing or stale data must remain visible rather than being rendered
as healthy.

Operators should generate snapshots on trusted workstations or isolated CI,
review them before sharing and delete them according to retention policy. Do
not include state/plan values, credentials, kubeconfigs, secret environment
variables, private audit payloads or unredacted customer/infrastructure data.

The MVP deliberately has no apply/destroy button. Infrastructure mutation
continues through Git, reviewed plan files, policy gates and explicit production
approval. See [RFC 012](rfcs/012-web-dashboard.md).
