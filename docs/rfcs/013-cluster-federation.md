# RFC 013: Cluster federation and multi-cluster placement

Status: proposed

## Context

ClusterForge can describe multiple clusters, run read-only fleet health checks,
and render Argo CD applications from inventory. Those capabilities do not make
the clusters a federation: they do not provide global scheduling, shared service
discovery, data replication, traffic failover, or a common security boundary.

This RFC evaluates multi-cluster application placement while keeping
infrastructure, GitOps, traffic, and data ownership visible.

## Use cases

- **Active/active applications:** independently deploy replicas in multiple
  regions for latency and regional application resilience.
- **Region failover:** maintain a warm or cold target and move traffic through a
  reviewed DNS/load-balancer procedure after health and data checks.
- **Tenant isolation:** place tenants or sensitivity tiers in separate clusters
  while retaining consistent policy and deployment metadata.
- **Edge clusters:** distribute selected workloads to constrained or
  intermittently connected sites.
- **Disaster recovery:** activate workloads elsewhere from tested configuration
  and backups within documented RTO/RPO targets.

These cases have different consistency, networking, availability, cost, and
operating models. A single `federated = true` abstraction would hide critical
decisions and is not proposed.

## Options considered

| Option | What it provides | Fit and limitations |
|---|---|---|
| GitOps multi-cluster rendering | Deterministic desired-state placement per inventory cluster | Recommended initial control plane; cluster registration remains separate and runtime failover is not coordinated |
| Argo CD ApplicationSet | Generators, placement templates, per-cluster labels, Git-driven rollout | Likely Argo CD evolution; generator credentials and cluster registration stay outside ClusterForge |
| Flux multi-cluster | Native Kustomization/HelmRelease reconciliation per cluster | Valuable second provider requiring a native renderer, not an artificial common API |
| Kubernetes Cluster API | Declarative cluster and machine lifecycle | Complementary provisioning, not application federation or traffic/data management; adds management-cluster risk |
| Service mesh multi-cluster | East-west identity, discovery and traffic policy | Deferred due to connectivity, trust, certificate, gateway, latency and failure-mode ownership |
| DNS-based failover | Coarse public traffic steering with health and weighted/failover records | Useful for many stateless services; caching, false health and data readiness prevent hidden/instant failover |

ApplicationSet should follow only after the static renderer proves useful. Flux
support should emit native Flux resources. Cluster API and service mesh may
integrate later, but neither is required for a safe GitOps foundation.

## Non-goals

- An automatic global scheduler in early ClusterForge versions.
- Replacing cloud load balancers, DNS services, CDNs, or traffic managers.
- Hidden failover, automatic write promotion, or zero-downtime claims.
- Replicating databases, queues, volumes, encryption keys, or secrets.
- Automatically registering clusters or storing Git/cloud/cluster credentials.
- Treating multiple clusters as one failure or security domain.

## Proposed initial support

1. Keep cluster/environment identity and placement inputs in inventory, using
   stable names and explicit regions/providers.
2. Render reviewed Argo CD app-of-apps and per-cluster overlays. Add
   ApplicationSet output only after dedicated diff and exclusion tests.
3. Maintain DNS failover guidance exposing health criteria, TTL, approval, data
   readiness, rollback, and provider-native behavior.
4. Expand fleet health as a read-only signal. It may inform operators but must
   not mutate placement or traffic.
5. Keep scheduling explicit in Git. Perform no automatic cross-cluster
   scheduling, traffic switching, or data promotion.

The initial model is coordinated multi-cluster management, not a federation
product. Terraform owns infrastructure/bootstrap; GitOps owns declared workloads
after handoff.

## Placement model

A future declaration may select explicit clusters or reviewed labels such as
region, environment, tenant, capability, or compliance boundary. Rendering must
resolve selectors to a visible destination list and fail when nothing matches.
Production placement must not be inferred only from naming conventions.

Placement changes require a diff, policy evaluation, approval, ordered rollout,
health observation, and rollback plan. Stateful applications additionally need
application-owned data topology and promotion procedures.

## Risks and required controls

- **Complexity:** controller, repository, version and failure states multiply.
  Keep one owner per resource and publish cluster-specific status.
- **Network connectivity:** private APIs, east-west traffic, DNS, firewalls and
  asymmetric routes fail independently. Document every required path.
- **Data consistency:** active/active writes can diverge without explicit
  application/database support. Placement does not imply replication.
- **Operational burden:** upgrades, policy, capacity, cost and on-call duties
  multiply. Fleet summaries must retain per-cluster detail.
- **Secrets replication:** copying Secrets expands blast radius and creates stale
  credentials. Each cluster should resolve external references using workload
  identity; ClusterForge must never copy values.
- **Correlated failure:** shared Git, identity, DNS, management clusters, CAs or
  CI can defeat redundancy. Architecture reviews must map shared dependencies.

## Acceptance criteria for further implementation

- Two-cluster output is deterministic, credential-free, previewable for one
  cluster, and identifies every destination.
- A disposable exercise demonstrates deploy, partial failure, rollback and
  removal without affecting an unintended cluster.
- DNS exercises record detection, approval, propagation, data readiness,
  recovery time and failback.
- Fleet health distinguishes unknown/unreachable from unhealthy and never
  triggers mutation.
- Secret references, data ownership, capacity and recovery are defined per
  target cluster.

Only after operating evidence exists should ClusterForge consider ApplicationSet
or Flux generation. Automatic scheduling remains a separate future RFC.

## Decision

Adopt **multi-cluster inventory and GitOps placement** as initial support, while
describing cluster federation as evaluated but not implemented. Do not build a
global scheduler or automatic failover system in the early roadmap.
