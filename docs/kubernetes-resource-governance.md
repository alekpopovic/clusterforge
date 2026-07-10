# Kubernetes resource governance

ClusterForge provides standalone `resource-quota` and `limit-range` modules for
explicit, namespace-by-namespace rollout.

## Requests and limits

A request is the capacity the scheduler reserves. A limit is the maximum runtime
usage Kubernetes allows for a container. CPU above a limit is throttled; memory
above a limit can terminate the container. Choose values from observed usage and
leave headroom for startup and traffic bursts.

LimitRange can supply defaults when a manifest omits values, but hidden defaults
can surprise application owners. Publish the namespace policy and prefer explicit
workload requests and limits over relying indefinitely on admission defaults.

## Namespace quotas

ResourceQuota caps aggregate namespace consumption, including requests, limits,
object counts, and storage classes supported by Kubernetes. A quota can reject a
new deployment even when the cluster has spare capacity, so monitor quota usage
and define an escalation process.

Do not copy the example quantities directly into production. Inventory workloads,
calculate normal and peak use, then introduce quotas with enough headroom.

## HPA interaction

CPU utilization targets compare observed CPU to container CPU requests. Missing
or unrealistic requests make HPA behavior unreliable. LimitRange defaults can
make metrics available, but app-owned requests are clearer. Ensure quota has room
for the maximum replica count, including rollout surge capacity.

## Common pitfalls

- A quota for CPU or memory may require every container to declare the matching
  request or limit.
- Defaults that are too small cause throttling, OOM kills, or unnecessary scaling.
- Defaults that are too large reduce bin packing and exhaust quota early.
- Deployment rollouts temporarily need capacity for old and new replicas.
- Init containers and sidecars also consume quota and must be included.
- Namespace quotas do not enforce node, cluster, or cloud-account budgets.

Roll out to a non-production namespace first, observe admission failures and
utilization, then promote reviewed settings explicitly.
