## Prompt 133 — Azure and GCP production hardening docs

```text
Create production hardening documentation for AKS and GKE.

Create:
- docs/azure-aks-production.md
- docs/gcp-gke-production.md

For AKS cover:
- private cluster considerations
- managed identity
- workload identity
- Azure CNI
- network policy
- node pool separation
- Azure Monitor
- Key Vault CSI / External Secrets
- Azure DNS
- backup options
- upgrade process
- RBAC and Azure AD integration
- state backend with Azure Storage

For GKE cover:
- regional clusters
- private clusters
- Workload Identity
- VPC-native clusters
- node pool separation
- Cloud NAT
- Cloud DNS
- Secret Manager and External Secrets
- Cloud Logging/Monitoring
- backup options
- upgrade process
- IAM least privilege
- state backend with GCS

Rules:
- Do not claim feature parity with EKS unless implemented.
- Mark implemented vs planned features.
- Avoid cloud credentials.
- Include production checklist.

Update:
- docs/roadmap.md
- VERSION_MATRIX.md if needed
```


---
