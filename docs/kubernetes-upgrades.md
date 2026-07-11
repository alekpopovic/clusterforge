# Kubernetes upgrade planning

`cf k8s upgrade plan <cluster-or-env> --target-version 1.31` performs a local,
read-only assessment. `check` returns a blocking exit code for invalid jumps;
`versions` prints configured control-plane/node versions.

The planner permits one minor version at a time, compares the target with the
local `VERSION_MATRIX.md`, checks node alignment, scans YAML/Terraform for known
deprecated API versions, and reminds operators to review Helm/platform add-ons.
It does not query the internet or perform an upgrade. When kubectl/live access
is unavailable, this limitation is explicit.

Review provider support, backups, disruption budgets, deprecated APIs, CRDs,
and workload health before upgrading. Control-plane downgrade is generally not
a rollback strategy; preserve tested backups and use provider-specific
forward-fix or restore procedures.
