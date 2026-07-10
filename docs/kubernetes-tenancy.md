# Kubernetes tenancy

ClusterForge tenants provide namespace-level organizational and safety controls.
They are soft multi-tenancy, not a security boundary equivalent to separate
clusters.

## Namespace model

Namespace-per-team reduces administrative overhead and fits teams with a shared
release and access model. Namespace-per-app provides clearer ownership, quotas,
and network boundaries but creates more objects to operate. A practical model is
namespace-per-team per environment, with dedicated app namespaces for workloads
that need separate permissions or lifecycle.

The tenant module can own several namespaces and labels all of them with
`clusterforge.io/tenant`. Avoid assigning the same namespace to multiple tenant
module instances.

## Quotas and limits

ResourceQuota caps aggregate namespace consumption. LimitRange supplies or
constrains per-container defaults. Both are opt-in because introducing them can
reject existing workloads. Inventory current requests and limits before rollout,
then start with measured headroom.

Quotas only work predictably when workloads declare requests and limits. An HPA
using CPU utilization also needs CPU requests; otherwise metrics-based scaling
may be unavailable or misleading.

## Pod Security

Namespaces receive Pod Security Admission labels. The default enforces
`baseline` while auditing and warning on `restricted`. Move enforcement to
`restricted` only after warnings have been resolved. Privileged system workloads
should use separate, tightly controlled namespaces.

## Network policies

Default deny ingress and egress are disabled by default. Enable them only with a
NetworkPolicy-capable CNI and after defining the traffic that must remain allowed.
The optional DNS egress policy opens TCP and UDP port 53 toward `kube-system`;
clusters with NodeLocal DNSCache or custom DNS need adjusted policies.

## RBAC limitations

Tenant RBAC creates a Role and RoleBinding in each namespace. It never creates a
ClusterRole by default. The built-in rule is read-only for pods, services, and
ConfigMaps; supply explicit rules when a team needs deployment permissions.
Kubernetes RBAC grants access but does not configure identity-provider groups,
admission policy, network isolation, or secret-store authorization.

Review subject names carefully. ServiceAccount subjects default to the tenant
namespace unless an explicit namespace is provided.
