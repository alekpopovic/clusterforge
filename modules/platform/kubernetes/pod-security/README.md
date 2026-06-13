# platform/kubernetes/pod-security

## Purpose

Applies Kubernetes Pod Security Admission labels to existing namespaces.

## Status

Implemented.

## Usage

```hcl
module "pod_security" {
  source = "../../../modules/platform/kubernetes/pod-security"

  namespaces = {
    apps = {
      enforce = "baseline"
      audit   = "restricted"
      warn    = "restricted"
    }
  }
}
```

## Pod Security Levels

- `privileged`: least restrictive; allows known privilege escalations.
- `baseline`: blocks common privilege escalations while preserving broad
  compatibility.
- `restricted`: most restrictive; can break workloads that need elevated Linux
  capabilities, host access, or non-default pod security settings.

## Approach And Limitations

This module uses `kubernetes_labels` against Namespace objects. That manages
only labels and avoids taking ownership of existing namespaces with
`kubernetes_namespace_v1`.

The namespace must already exist. If another controller also manages the same
labels, Terraform and that controller may drift or fight over values. Roll out
Pod Security labels gradually, starting with `audit` and `warn` before moving
important workloads to stricter `enforce` settings.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
