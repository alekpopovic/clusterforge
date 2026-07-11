# RFC 015: Windows container workloads

Status: experimental assessment

## Decision summary

ClusterForge should prepare for Windows containers, but must not claim Windows
support until disposable real-cluster and real-ECS tests pass. The initial change
is documentation, a declarative app `platform` field, and an ECS task
`runtime_platform` input. Windows node lifecycle automation and Kubernetes
scheduling generation remain out of scope for the first version.

## Kubernetes Windows nodes

Windows containers require Windows worker nodes compatible with the container
base image and the Kubernetes/cloud support matrix. Linux control-plane and
system workloads remain necessary; many DaemonSets, CSI/CNI components, security
agents, admission assumptions, probes, volume types and privileged features are
Linux-specific.

Workloads need an explicit OS node selector such as `kubernetes.io/os: windows`.
Dedicated node pools should also use a reviewed taint and matching toleration so
Linux pods do not land there and Windows pods do not schedule accidentally.
Architecture and Windows build compatibility must be explicit. Windows images
are typically much larger, pull/start more slowly, and must not be built or
assumed interchangeable with Linux images under the same tag.

Networking behavior depends on the cluster distribution, cloud CNI, kube-proxy
mode and Windows host version. NetworkPolicy, service routing, DNS, host process
containers, ingress, load balancers and source IP behavior need platform-specific
tests. Pod Security and Linux capability/seccomp controls do not translate
directly; policy packs need OS-aware exclusions without broadly bypassing Windows
workloads.

## ECS Windows tasks

The ECS service module is Fargate-only and now passes an explicit task-definition
runtime platform, defaulting to `LINUX`/`X86_64`. Windows Fargate availability,
regions, supported Windows families, CPU/memory combinations and platform versions
must be checked against AWS before deployment. This interface is experimental and
has not been validated by ClusterForge against a live Windows task.

Windows task definitions still use `awsvpc`, execution/task IAM roles and secret
references, but image entrypoints, path syntax, environment behavior, health
checks and storage differ. The `awslogs` driver requires an application/logging
design that handles Windows event/application output; creating a log group alone
does not guarantee useful logs. Task IAM remains workload-scoped and does not
replace host/domain identity.

Windows images and task sizes commonly increase registry storage, transfer,
startup time, memory/CPU needs and Fargate cost. Capacity, deployment circuit
breakers, ALB health, autoscaling signals, patch cadence and rollback must be
measured rather than copied from Linux defaults. ECS on EC2 Windows would add AMI,
host patching, capacity provider and node lifecycle concerns and is not included.

## Module and schema impact

- `workloads/kubernetes/app`: needs optional node selectors, taints/tolerations,
  OS-aware security context and probe/path behavior before Windows rendering can
  be enabled. Current rendering remains Linux-oriented.
- `workloads/ecs/service`: exposes `runtime_platform`; Linux/X86_64 stays the
  backward-compatible default. Windows/ARM is rejected by the initial interface.
- Kubernetes node-group modules: future work must create separate Windows pools,
  validate supported versions/images, apply taints/labels, and document upgrades,
  bootstrap, logging and CNI dependencies.
- App manifest schema accepts:

  ```yaml
  platform:
    os: windows
    architecture: amd64
  ```

  Allowed values are `linux|windows` and `amd64|arm64`; Windows is limited to
  amd64. This field describes intent only. It does not currently cause the
  Kubernetes renderer to schedule Windows pods or select a Windows Server version
  for ECS.

## Proposed MVP and acceptance criteria

1. Keep Windows explicitly experimental in docs and module descriptions.
2. Retain the ECS runtime input and validate Terraform plans for Linux default and
   reviewed Windows Server variants.
3. Design Kubernetes `node_selector` and `tolerations` schema as general workload
   scheduling features, then add OS-aware rendering with unit tests.
4. Add separate Windows examples only after testing on disposable EKS/AKS and ECS
   environments with pinned image/host versions.
5. Exercise image pull, DNS/service networking, ALB health, secrets, IAM, logs,
   restart, scale, rolling deployment, node/task patch, rollback and cleanup.
6. Record cost/startup/resource differences and unsupported features.

## Non-goals

- Full Windows node provisioning, bootstrap, patching or lifecycle automation in
  the first version.
- Active Directory/domain-join automation or credential custody.
- HostProcess workloads, mixed-OS image construction, or Windows build pipelines.
- Claiming parity with Linux Kubernetes/ECS modules before acceptance evidence.

## Recommendation

Merge only the low-risk schema and ECS task-definition preparation. Keep generated
Kubernetes applications on the existing Linux path and require users to manage
Windows manifests outside ClusterForge until scheduling and live validation are
implemented. Label all Windows examples and inputs experimental.
