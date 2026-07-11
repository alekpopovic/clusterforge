# RFC 014: Edge and constrained deployments

Status: proposed

## Context

ClusterForge has experimental K3s and RKE2 cloud-init generators, but it does not
yet provide an edge lifecycle: hardware provisioning, offline artifact delivery,
device identity, local storage, upgrades, observability, or recovery remain
operator responsibilities. Edge sites amplify the cost of every dependency and
make cloud-style connectivity assumptions unsafe.

This RFC defines an evaluation boundary. It does not make existing lightweight
targets production-ready and proposes no code in this phase.

## Use cases

- Retail locations running local transaction, inventory, or signage workloads.
- IoT gateways aggregating devices and filtering data near its source.
- On-premises mini clusters serving a factory, office, vessel, or branch.
- Disconnected or intermittently connected environments with delayed sync.
- Disposable and persistent lab environments for development and validation.

Critical safety, medical, industrial-control, or life-support workloads require
domain-specific certification and controls outside ClusterForge's scope.

## Target assessment

| Target | Initial position | Important caveats |
|---|---|---|
| K3s | Preferred first lightweight Kubernetes target; experimental | Bundled component choices, SQLite vs external/embedded etcd, token handling, upgrade channels and architecture compatibility need testing |
| RKE2 | Preferred hardened/on-prem alternative; experimental | Higher resource footprint; configuration, CIS profiles, embedded registry and server/agent lifecycle need tested modules |
| MicroK8s | Evaluate, do not implement initially | Snap-based lifecycle and host assumptions can conflict with immutable/offline appliance models; clustering and add-on behavior need separate research |
| Single-node Kubernetes | Supported architecture pattern only after failure semantics are explicit | No node HA; maintenance is an outage; local storage loss can be total cluster/data loss |
| Remote Docker host | Experimental non-Kubernetes option for simple workloads | No Kubernetes API/policy ecosystem, weak multi-service orchestration boundary, and remote daemon access is highly privileged |

The first supported profile should be one Linux architecture and OS matrix using
K3s, followed by RKE2. Multi-architecture bundles and MicroK8s should not be
claimed until independently exercised.

## Environmental constraints

- **Limited CPU and memory:** controllers, admission webhooks, sidecars, log
  agents, scanners, and monitoring stacks need measured budgets and opt-in
  profiles. Control-plane starvation can look like application failure.
- **Unreliable connectivity:** management must tolerate long outages, avoid
  command-and-control dependence, expose last-sync time, and queue only bounded
  data. Site workloads need autonomous degraded behavior.
- **Offline upgrades:** every binary, image, chart, provider and signature must be
  resolved into a versioned bundle before transfer. Upgrade and downgrade limits
  require rehearsal on matching hardware.
- **Local storage:** storage class, disk health, encryption, capacity, backup and
  node replacement behavior must be explicit. A single local disk is not HA.
- **Simplified observability:** retain small local health/log windows, bounded
  cardinality and backpressure; forward when connected without exhausting disk.
- **Local registry:** mirror only an allowlisted, immutable image set, verify
  signatures/digests locally, control garbage collection, and plan registry
  recovery. A mirror must not silently fall back to the public internet.

## Proposed modules and profiles

1. `edge/registry-mirror`: configure an existing local registry/mirror endpoint,
   CA trust, digest-only policy and explicit upstream/offline mode. It must not
   embed registry credentials in Terraform state or cloud-init.
2. `edge/observability-lite`: resource-bounded node/application metrics and logs,
   local retention limits, disk-pressure protection and optional remote write.
3. `edge/app-profile`: conservative requests/limits, replicas, update strategy,
   probes, architecture constraints, local dependency declarations and offline
   image validation.
4. `edge/backup-target`: explicit local removable/object target or remote target,
   encryption reference, retention and restore-test metadata. Backups must remain
   useful during management-plane outages.
5. GitOps pull profile: a site-local agent pulls signed desired state, keeps the
   last known good revision, reports last sync, and never requires inbound access
   from a central runner.

These should compose with K3s/RKE2 modules rather than creating a hidden edge
distribution. Server provisioning, disk partitioning and OS patching remain
visible external layers unless separately implemented.

## Proposed CLI surface

- `cf edge init`: scaffold non-secret site inventory, target/profile, architecture,
  resource budget, connectivity mode and artifact sources. It must not provision
  the device or request credentials.
- `cf edge bundle`: resolve a lock file into a deterministic offline directory or
  archive containing manifests, charts, images/provider artifacts and checksums.
  Signing keys remain outside ClusterForge; bundle creation fails on unresolved or
  unverified inputs.
- `cf edge status`: read local/imported status such as version, capacity, last
  GitOps sync, backup evidence and connectivity age. Remote access must be
  explicitly configured and read-only by default.

No command is implemented by this RFC. Detailed command contracts, schemas,
threat model, size limits and acceptance tests are prerequisites.

## Security model

- **Secrets distribution:** bundles and Git contain references only. A site uses
  device/workload identity to retrieve scoped values or receives an encrypted,
  auditable provisioning package through an external process. Rotation must
  tolerate offline sites and expire lost-device access.
- **Local credentials:** kubeconfigs, bootstrap tokens, registry credentials and
  recovery keys stay in root-owned stores, never Terraform output/state, bundle
  indexes, logs, or support archives. Bootstrap credentials are rotated promptly.
- **Device identity:** each site/device has unique, revocable identity rooted in
  an approved PKI/TPM or equivalent hardware-backed mechanism where available.
  Shared fleet credentials are prohibited.
- **Update signing:** verify bundle manifest, component checksums, image digests and
  signatures offline before installation. Protect against rollback to vulnerable
  versions and record signing identity/version.
- **Physical threat:** assume disks/devices can be stolen or tampered with. Use
  secure boot and disk encryption where supported, minimize cached secrets, and
  define decommission/revocation procedures.

## Failure, upgrade, and recovery behavior

Sites retain a tested last-known-good application revision and bounded local
artifacts. Kubernetes version changes are staged by hardware/OS cohort, never the
entire fleet. Health gates must distinguish disconnected, unknown, degraded, and
failed. Automatic rollback must not reverse irreversible data migrations.

Backup validation uses an isolated matching target where practical. Single-node
replacement must document reconstruction of OS, runtime, registry, configuration,
identity, secrets, workloads and local data. Central dashboards cannot be the
only source of recovery information.

## Acceptance criteria

- A documented hardware/OS/architecture matrix passes install, reboot, power-loss,
  disk-pressure, 72-hour disconnection, reconnect and cleanup tests.
- A pinned offline bundle installs without DNS/internet and verifies all artifacts.
- GitOps retains known-good workloads while disconnected and reconciles bounded,
  observable changes after reconnect.
- Credential bootstrap, rotation, lost-device revocation and signed update failure
  are exercised without exposing secrets.
- Upgrade, rollback limits, backup and full single-node rebuild meet published
  resource use and recovery measurements.

## Decision

Keep K3s and RKE2 experimental while designing the K3s-first edge profile and
offline bundle contract. Do not implement `cf edge` commands or claim MicroK8s,
single-node production readiness, or remote Docker security until the acceptance
criteria have reproducible evidence.
