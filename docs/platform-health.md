# Platform health checks

`cf health check <env>` and `cf health report <env>` perform read-only checks
for config/path availability, local state presence, workload manifests and,
when explicitly configured, kubectl node/namespace and Helm release reads.
Use `--json` for automation.

Remote state is not contacted automatically, and live cluster checks are
skipped cleanly without a configured kubeconfig or required binary. A passing
local report is not proof of application availability, ingress reachability,
data correctness, SLO attainment or disaster recovery. Production validation
must include external monitoring and manually reviewed evidence.
