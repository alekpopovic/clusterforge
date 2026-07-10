# platform/kubernetes/gatekeeper

Optional OPA Gatekeeper installation with user-supplied ConstraintTemplates and
Constraints. No policies are installed or enforced by default.

## Status

Implemented.

## Usage

```hcl
module "gatekeeper" {
  source = "../../../modules/platform/kubernetes/gatekeeper"

  chart_version       = "<reviewed-version>"
  constraint_templates = {}
  constraints          = {}
}
```

On a fresh cluster, install Gatekeeper first and add templates/constraints in a
second plan because `kubernetes_manifest` needs CRD schemas during planning.
Start constraints with `enforcementAction: dryrun`, review audit results, and
promote individually. Do not run Gatekeeper and Kyverno together without an
explicit ownership and conflict analysis.

Provider configuration belongs in the calling root module.

## Generated Terraform documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
