## Prompt 68 — GKE module design RFC

```text
Create an RFC for Google Kubernetes Engine support.

Create:
- docs/rfcs/003-gke-support.md

Include:
1. Goals:
   - GCP network
   - GKE cluster
   - node pools
   - Workload Identity
   - optional Cloud DNS later

2. Proposed modules:
   - modules/cloud/gcp/network
   - modules/cloud/gcp/iam
   - modules/cloud/gcp/dns
   - modules/orchestrators/kubernetes/gke

3. Inputs/outputs.

4. Provider strategy:
   - google provider in root
   - kubernetes/helm provider configured from GKE outputs

5. Security:
   - Workload Identity
   - service account permissions
   - no key files committed

6. CLI:
   - cf env create dev --cloud gcp --orchestrator gke --region europe-west1
   - cf generate dev

7. Examples:
   - examples/gcp-gke-minimal
   - live/dev/gcp-gke

8. Risks:
   - regional vs zonal clusters
   - private clusters
   - IAM permissions
   - network/firewall assumptions

Do not implement code yet.
Update docs/roadmap.md.
```

---
