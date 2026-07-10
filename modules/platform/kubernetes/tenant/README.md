# platform/kubernetes/tenant

Creates one or more namespaces for a team or application tenant, applies Pod
Security Admission labels, and optionally creates quotas, default limits,
network isolation, and namespace-scoped RBAC.

## Status

Implemented.

## Usage

```hcl
module "payments_tenant" {
  source = "../../../modules/platform/kubernetes/tenant"

  name       = "payments"
  namespaces = ["payments-dev", "payments-prod"]

  pod_security = {
    enforce = "baseline"
    audit   = "restricted"
    warn    = "restricted"
  }

  resource_quota = {
    enabled = true
    hard = {
      "requests.cpu"    = "4"
      "requests.memory" = "8Gi"
      pods              = "30"
    }
  }

  network_policy = {
    default_deny_ingress = true
    default_deny_egress  = false
  }

  rbac = {
    create = true
    subjects = [{
      kind = "Group"
      name = "payments-developers"
    }]
  }
}
```

## Safe defaults

Quotas, limit ranges, network default-deny policies, and RBAC are disabled by
default. Pod Security defaults to baseline enforcement with restricted audit and
warning. The default RBAC rule is read-only and namespace-scoped. No ClusterRole
or ClusterRoleBinding is created.

Default-deny policies require a NetworkPolicy-capable CNI and explicit allow
policies for application traffic. DNS egress is only added when default-deny
egress is enabled. Review cluster DNS topology before production rollout.

Provider configuration belongs in the calling root module.

## Generated Terraform documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
