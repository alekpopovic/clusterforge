# Docker and Docker Swarm targets

Status: **experimental**.

Docker Engine is useful for local experiments and a small operator-managed
host. Swarm can support simple legacy/migration deployments, but ClusterForge
does not provide a full production control plane, host provisioning, quorum
operations, node replacement, overlay-network diagnosis, or automated rollback.

Security boundaries are host-wide: daemon/socket access is effectively root,
containers may share kernel risk, and hardening depends on the operator.
Networking, ingress, firewall and service discovery are less standardized than
the supported Kubernetes/ECS paths. Do not put secrets in environment variables
or Terraform state; Swarm secret lifecycle and rotation remain operator-owned.

For larger or strongly isolated production platforms, prefer Kubernetes/EKS,
ECS, or an explicitly operated Nomad deployment. Use Docker only when its
simpler lifecycle and limitations are understood, and maintain tested host
backup, image pinning, upgrade, drain and rollback procedures.
