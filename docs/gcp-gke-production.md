# GKE production hardening

GKE support is **experimental** and does not have production parity with EKS.
The current module creates a VPC-native cluster using secondary ranges,
Workload Identity, removal of the default node pool, and configurable managed
node pools/version.

## Implemented today

- VPC-native IP allocation using named pod and service secondary ranges.
- GKE Workload Identity and GKE metadata mode on managed node pools.
- Configurable regional/location input, Kubernetes version and node pools.
- Generated GCP root and experimental example validation.

## Required production work

- Choose a regional topology and define zone/node-pool failure behavior.
- Enable private nodes/control-plane access, master authorized networks,
  Private Google Access and Cloud NAT with explicit egress logging.
- Separate system/workload node pools; define autoscaling, surge upgrades,
  maintenance windows, release channels and disruption budgets.
- Review Google groups/RBAC, IAM least privilege, service accounts and
  Workload Identity bindings; avoid broad cloud-platform privileges.
- Configure Cloud Logging/Monitoring, audit retention and alert routing.
- Integrate Secret Manager through External Secrets without storing values in state.
- Define Cloud DNS, TLS, load-balancer exposure, Armor/firewall requirements.
- Select Backup for GKE or Velero and record real restore evidence and RTO/RPO.
- Use an encrypted, versioned, access-controlled GCS backend with state
  separation; credentials must come from short-lived identity federation.

## Production checklist

- [ ] Regional/private topology and authorized control-plane access tested.
- [ ] VPC-native ranges, Cloud NAT, firewall and NetworkPolicy reviewed.
- [ ] Workload Identity, IAM, RBAC and node-pool separation verified.
- [ ] Logging, monitoring, secrets, DNS and ingress controls have owners.
- [ ] Upgrade, backup, restore, rollback/forward-fix and cleanup evidence exists.
- [ ] GCS state permissions, versioning and recovery are tested.

ClusterForge does not automatically replicate application data, guarantee
recovery objectives, or currently enforce all controls above.
