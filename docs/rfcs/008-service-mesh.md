# RFC 008: Optional service mesh support

Status: proposed

## Context

ClusterForge now has opt-in progressive delivery and namespace security building
blocks. A service mesh could add workload identity, mTLS, traffic control, and
uniform telemetry, but it also inserts critical infrastructure into every
service-to-service request path. It must remain an explicit platform choice.

## Goals

- Mutual TLS for service-to-service traffic with verifiable workload identity.
- Weighted traffic shifting for canary and blue/green delivery.
- Consistent request metrics, traces, and access telemetry.
- Service authorization policy independent of application code.
- Stronger east-west service security without static credentials.

## Non-goals for the first implementation

- Multi-cluster mesh or cross-region failover.
- Automatic zero-trust policy generation or remediation.
- Enterprise or hosted control planes.
- Enabling a mesh in bootstrap or existing environments by default.
- Hiding provider-native APIs behind an incomplete universal abstraction.

## Options

| Option | Strengths | Costs and gaps | Best fit |
| --- | --- | --- | --- |
| Istio | Rich mTLS, authorization, ingress/egress, telemetry, and weighted traffic APIs; strong Argo Rollouts integrations; sidecar and ambient data-plane choices | Largest API and operational surface; Envoy/waypoint resource overhead; upgrades and debugging require dedicated ownership | Teams needing complete L7 traffic and policy capability on Kubernetes |
| Linkerd | Small operational model, automatic mTLS, useful golden metrics, and lightweight Rust proxy | Legacy SMI TrafficSplit/linkerd-smi path is deprecated; advanced routing integration must use the current dynamic routing model; Kubernetes-focused | Teams prioritizing simple mTLS and observability over a broad traffic-policy API |
| Consul service mesh | Envoy-based mTLS, intentions, L7 traffic management, and alignment across Kubernetes, Nomad, VMs, and other runtimes | Adds Consul servers/catalog and another service-discovery control plane; Kubernetes-only users pay extra operational cost; enterprise boundaries require review | Organizations already operating Consul or requiring cross-runtime service networking |

Istio supports percentage-based routing and detailed traffic policy through its
traffic APIs. Linkerd automatically applies mTLS to meshed pods, but its older
SMI TrafficSplit extension is deprecated. Consul is the strongest cross-runtime
fit and supports Kubernetes and Nomad, but introduces the broadest additional
control-plane commitment for a Kubernetes-first user.

## Recommendation

Implement **Istio sidecar mode first**, as an optional and narrowly scoped
Kubernetes module. It best matches the current Argo Rollouts traffic-shifting
work while covering all stated security, policy, and telemetry goals. Sidecar
mode is selected for the first iteration because workload injection and L7
behavior are explicit and mature. Evaluate ambient mode separately after its
feature and upgrade constraints are tested against ClusterForge examples.

The first module must install only the control plane. It must not enable global
injection, strict mTLS, an ingress gateway, or default-deny authorization. Those
are separate opt-in changes with staged examples and rollback guidance.

Linkerd remains the preferred second implementation for teams that want a
smaller mTLS/observability footprint. Consul service mesh should follow only when
the existing Nomad/Consul work proves a cross-runtime requirement.

## Proposed modules

- `modules/platform/kubernetes/istio`
  - install base CRDs and control plane with pinned charts;
  - expose revision and profile controls;
  - optionally create a separately managed ingress gateway;
  - never label workload namespaces automatically by default.
- `modules/platform/kubernetes/linkerd`
  - install CRDs and control plane;
  - require externally managed trust roots/issuer strategy for production;
  - keep injection opt-in.
- `modules/platform/kubernetes/consul-service-mesh`
  - install or connect to an explicit Consul control plane;
  - expose transparent proxy and injection settings;
  - keep Kubernetes service registration behavior visible.

## Workload impact

Mesh participation changes pod startup, shutdown, probes, resources, networking,
and debugging. Workload and tenant modules need optional labels/annotations for:

- namespace or pod-level sidecar injection;
- revisioned control-plane selection;
- application/proxy port exclusions and protocol hints;
- telemetry tags and scraping;
- ingress-gateway routing and service exposure.

Namespace injection must be applied separately from installing the control
plane. Existing workloads need readiness/probe validation, sidecar resource
budgets, Pod Security compatibility, and an explicit egress inventory.

## Ingress, traffic, and telemetry

Ingress gateways are optional public-facing workloads and must not be bundled
into a safe control-plane default. ClusterForge should initially integrate Argo
Rollouts with version-specific Services and Istio routing resources in a dedicated
example. Metrics-based promotion needs a separately configured provider and
reviewed success thresholds.

Mesh telemetry complements but does not replace application metrics and traces.
Sampling, cardinality, retention, and sensitive header handling remain explicit
operator responsibilities.

## Operational risks

- **Complexity:** new CRDs, proxies, certificate authorities, gateways, and policy
  APIs expand the incident surface.
- **Upgrades:** control plane, CRDs, data plane, gateways, and workload revisions
  may require ordered or revision-based migration.
- **Debugging:** application, DNS, CNI, proxy, policy, and certificate failures can
  look alike without mesh-specific tools and runbooks.
- **Resource overhead:** sidecars consume pod CPU/memory and increase scheduling,
  startup, and connection costs.
- **Compatibility:** custom protocols, existing TLS, init containers, probes,
  NetworkPolicy, host networking, and batch jobs need explicit tests.
- **Availability:** webhook or control-plane failure settings can either block
  deployments or weaken policy; both outcomes require documented recovery.

## Rollout and acceptance criteria

1. Install a pinned Istio version on a disposable cluster with no injected apps.
2. Opt one test namespace into injection and verify readiness, DNS, and egress.
3. Demonstrate mTLS identity and plaintext/non-mesh boundaries honestly.
4. Demonstrate Argo Rollouts canary routing with observable weights.
5. Exercise control-plane outage, certificate rotation, upgrade, rollback, and
   full cleanup.
6. Publish resource overhead and compatibility evidence before production status.

No mesh module should be considered stable until non-mesh workloads remain
unaffected and the opt-out/cleanup path is verified.

## References

- [Istio traffic management](https://istio.io/latest/docs/concepts/traffic-management/)
- [Istio sidecar and ambient modes](https://istio.io/latest/docs/overview/dataplane-modes/)
- [Linkerd automatic mTLS](https://linkerd.io/docs/features/automatic-mtls/)
- [Linkerd traffic split deprecation notice](https://linkerd.io/docs/features/traffic-split/)
- [Consul service mesh](https://developer.hashicorp.com/consul/docs/connect)
- [Consul L7 traffic management](https://developer.hashicorp.com/consul/docs/manage-traffic)
