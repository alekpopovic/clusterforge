# AKS production hardening

AKS support is **experimental** and does not have production parity with the
EKS path. The current module creates an AKS cluster with a system-assigned
managed identity, one configurable default node pool, subnet attachment,
optional Kubernetes version, and tags. Provider configuration remains in the
root.

## Implemented today

- System-assigned cluster managed identity.
- Azure VNet/subnet composition through the separate network module.
- Configurable Kubernetes version and default node-pool size/SKU.
- Generated Azure root and experimental example validation.

## Required production work

- Decide private cluster/API DNS topology and authorized operator/CI access.
- Configure Azure CNI mode, NetworkPolicy, egress control, private endpoints,
  and separate system/user node pools with autoscaling and disruption policy.
- Enable Microsoft Entra ID integration, Azure RBAC/Kubernetes RBAC, local
  account restrictions, workload identity and OIDC issuer.
- Add Azure Monitor/diagnostic logs with retention and alert ownership.
- Integrate Key Vault CSI or External Secrets without putting values in state.
- Define Azure DNS ownership, TLS, ingress exposure, and DDoS/WAF requirements.
- Select Azure Backup/Velero patterns and execute restore tests.
- Establish maintenance windows, supported minor-step upgrades, node-image
  upgrades, rollback/forward-fix procedures, and workload compatibility tests.
- Use an encrypted, access-controlled Azure Storage backend with locking and
  separate state per environment; do not commit backend credentials.

## Production checklist

- [ ] Private/public API decision and authorized networks reviewed.
- [ ] Managed/workload identities and least-privilege roles reviewed.
- [ ] Azure CNI, NetworkPolicy, egress, DNS, and ingress exposure tested.
- [ ] System and workload node pools are isolated and upgrade-tested.
- [ ] Monitoring, audit, secrets, backup, restore, RTO/RPO and cleanup evidence exists.
- [ ] Azure Storage state access, encryption and recovery are tested.

These are operator requirements and planned module capabilities, not claims
that ClusterForge currently enforces them.
