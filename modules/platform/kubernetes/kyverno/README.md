# platform/kubernetes/kyverno

Installs Kyverno from its official Helm repository and can apply a small,
opt-in ClusterPolicy baseline plus user-provided policy YAML.

## Status

Implemented. The baseline uses legacy `kyverno.io/v1` ClusterPolicy resources
because that is the scope of this module version. Pin and test a compatible
Kyverno chart; migrate to stable ValidatingPolicy resources before upgrading to
a Kyverno release that removes ClusterPolicy support.

## Usage

```hcl
module "kyverno" {
  source = "../../../modules/platform/kubernetes/kyverno"

  chart_version             = "<reviewed-version>"
  enable_baseline_policies  = true
  baseline_failure_action   = "Audit"
}
```

The baseline can report privileged containers, missing CPU/memory requests and
limits, and images using the `latest` tag. Optional extended controls restrict
host networking, hostPath, root execution and capabilities, apply a registry
allowlist, and require image digests. `Audit` is the default and does not block
workload admission. Change to `Enforce` only after reviewing reports and
remediating affected workloads.

## CRD dependency

`kubernetes_manifest` needs Kyverno CRDs to exist during planning. On a fresh
cluster, first install the Helm release with policies disabled, then enable
policies in a second plan/apply. A Terraform `depends_on` controls apply order but
cannot make an unknown CRD schema available during the initial plan.

User-provided `policies` values are decoded as YAML and applied after the Helm
release. Review them like code. The module does not validate their semantics or
automatically exclude system namespaces.

See `examples/kubernetes-kyverno-production-pack` and
`docs/kubernetes-admission-security.md` for the extended audit-first rollout.

Provider configuration belongs in the calling root module.

## Generated Terraform documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
