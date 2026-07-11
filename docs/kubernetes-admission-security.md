# Kubernetes admission security

ClusterForge supports either Kyverno or Gatekeeper as an opt-in admission layer.
Both module examples default to non-blocking audit/dry-run behavior; installing a
controller does not silently enforce the extended pack.

The extended controls cover image `latest` tags, registry allowlists and optional
production digest pinning; privileged containers, hostPath, host networking,
non-root execution and Linux capabilities; requests and limits; NetworkPolicy
coverage; secret-like literal environment values; and approval annotations for
public ingress. NetworkPolicy coverage and secret-pattern detection are advisory
because namespace intent and false positives require human context.

## Rollout

1. Choose one engine, pin a version compatible with the cluster, and install it
   without policies. Avoid overlapping Kyverno and Gatekeeper ownership.
2. Enable policies in `Audit` (Kyverno) or `dryrun` (Gatekeeper). Collect reports
   across normal deploy, autoscaling, jobs, system namespaces, and upgrades.
3. Assign finding owners, fix workloads, and document narrowly scoped exceptions
   with expiry dates. Never exempt an entire production namespace by default.
4. Promote one rule and a small non-production scope at a time. Test failure
   messages, controller outage behavior, latency, rollback, and break-glass access.
5. Opt into `Enforce` or `deny` only after violation counts are understood and
   rollback is rehearsed. Monitor rejected admissions and controller health.

Use `examples/kubernetes-kyverno-production-pack` for the built-in Kyverno
controls. The Gatekeeper example shows the ConstraintTemplate/Constraint pattern
and forces the selected enforcement action onto supplied constraints. The rule
catalog in `policies/packs/kubernetes-security/policies.yaml` provides stable
control identifiers; it does not itself install admission resources.
