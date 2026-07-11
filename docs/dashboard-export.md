# Dashboard data export

ClusterForge can produce a versioned JSON snapshot for a future dashboard. The
export is read-only, works offline, and only reads the project configuration and
repository files.

```bash
cf dashboard export
cf dashboard export --env prod
cf dashboard export --fleet
cf dashboard export --output dashboard-data.json
```

The default output is `dashboard-data.json`. It includes project, organization,
workspace, environment, cluster, stack, app, service catalog, runbook, and module
catalog metadata. When `.cf/dashboard/policy.json`, `drift.json`, or `cost.json`
exist, the export records their availability and paths without copying their
contents.

The schema is identified by `schema_version`. Consumers should reject unsupported
major versions and tolerate additional fields in compatible versions.

## Security

App environment variables, secret references, kubeconfig paths, credentials, and
cloud API data are excluded. The file is created with owner-only permissions
(`0600`) because the remaining operational metadata can still be sensitive. Review
the result before sharing it outside the project.
