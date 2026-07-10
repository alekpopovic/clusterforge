# Kubernetes policy with Kyverno

Kyverno is an admission and background policy engine. ClusterForge installs it
with the official Helm chart and keeps baseline policies disabled unless a root
module explicitly enables them.

## Audit and enforce

`Audit` allows admission and records policy violations. It is the module default
and the recommended starting point for existing clusters. `Enforce` rejects new
violations and can disrupt deployments, controllers, or recovery procedures when
rolled out without evidence.

Use this progression:

1. Pin a Kyverno chart version compatible with the cluster version.
2. Install Kyverno in its dedicated namespace with policies disabled.
3. Confirm controller readiness and webhook connectivity.
4. Enable baseline policies in `Audit` and review PolicyReports and warnings.
5. Add exclusions only when justified and narrowly scoped.
6. Remediate workloads, test failure and recovery paths, then promote individual
   policies to `Enforce` through a reviewed change.

High availability, resource sizing, failure policy, and namespace exclusions are
production decisions passed through `values`. Do not assume the chart defaults
fit every cluster.

## Baseline scope

The optional baseline reports or blocks:

- privileged containers;
- containers missing CPU and memory requests or limits;
- container images explicitly tagged `latest`.

Resource and image checks can be disabled independently. Custom YAML policies can
be supplied through the `policies` map. Test custom policies against representative
resources before cluster-wide rollout.

## CRDs and first installation

The Helm chart installs Kyverno CRDs. Terraform's `kubernetes_manifest` resource
needs those schemas during planning, so a new cluster requires two phases: install
Kyverno first, then enable baseline/custom policies. `depends_on` alone cannot
solve schema discovery during the first plan.

## ClusterPolicy lifecycle

This MVP follows the prompt's ClusterPolicy requirement. Current Kyverno docs
mark ClusterPolicy as legacy/deprecated and introduce stable ValidatingPolicy
types. Keep the chart pinned and schedule migration before upgrading beyond a
version that supports `kyverno.io/v1` ClusterPolicy.

Official references:

- [Kyverno installation](https://kyverno.io/docs/installation/installation/)
- [Policy type overview](https://kyverno.io/docs/policy-types/overview/)
- [Validation failure actions](https://kyverno.io/docs/policy-types/cluster-policy/validate/)
