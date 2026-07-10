# Kubernetes policy with Gatekeeper

Gatekeeper is ClusterForge's optional OPA-based alternative to Kyverno. The
module installs the official Helm chart but installs no policy by default.

## Gatekeeper or Kyverno

Choose Gatekeeper when the organization already operates OPA/Rego, needs the
ConstraintTemplate/Constraint model, or wants to reuse a reviewed Gatekeeper
policy library. Choose Kyverno when Kubernetes-native YAML authoring and its
broader mutate/generate/image policy workflow better match the platform team.

Neither engine is universally simpler. Avoid installing both by default: two
admission systems increase latency, failure modes, exclusions, and the risk of
contradictory decisions.

## Audit and enforcement

Gatekeeper constraints use `spec.enforcementAction`. Start with `dryrun` so
violations appear in audit without rejecting admission. `warn` can surface
admission warnings, while `deny` blocks violations. Confirm the exact actions
supported by the pinned Gatekeeper release.

Recommended rollout:

1. Pin and install Gatekeeper in its dedicated namespace.
2. Verify webhook and audit controller health.
3. Apply structural `templates.gatekeeper.sh/v1` ConstraintTemplates.
4. Add narrowly scoped `dryrun` constraints.
5. Review violations and test allowed/disallowed fixtures with `gator`.
6. Remediate workloads, then promote one constraint at a time to `deny`.

The CRDs must exist before Terraform can plan template and constraint manifests,
so fresh clusters require a two-phase apply. Review Rego/CEL as executable policy
code and keep system namespace exclusions explicit.

Official references:

- [Gatekeeper Helm installation](https://open-policy-agent.github.io/gatekeeper/website/docs/v3.14.x/install/)
- [ConstraintTemplates](https://open-policy-agent.github.io/gatekeeper/website/docs/constrainttemplates/)
- [Gatekeeper enforcement points](https://open-policy-agent.github.io/gatekeeper/website/docs/v3.17.x/enforcement-points/)
