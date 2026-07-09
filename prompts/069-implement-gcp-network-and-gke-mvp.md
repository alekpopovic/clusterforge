## Prompt 69 — Implement GCP network and GKE MVP

```text
Implement GCP GKE MVP based on docs/rfcs/003-gke-support.md.

Create modules:
- modules/cloud/gcp/network
- modules/orchestrators/kubernetes/gke

GCP network module:
Inputs:
- project_id
- name
- environment
- region
- auto_create_subnetworks default false
- subnet_cidr
- secondary_pod_range_cidr
- secondary_service_range_cidr
- labels

Resources:
- google_compute_network
- google_compute_subnetwork with secondary ranges

Outputs:
- network_id
- network_name
- subnetwork_id
- subnetwork_name

GKE module:
Inputs:
- project_id
- name
- environment
- region
- network
- subnetwork
- pods_secondary_range_name
- services_secondary_range_name
- kubernetes_version default ""
- remove_default_node_pool default true
- initial_node_count default 1
- node_pools map(object)

Resources:
- google_container_cluster
- google_container_node_pool

Outputs:
- cluster_name
- endpoint
- ca_certificate sensitive
- location

Create:
- examples/gcp-gke-minimal
- live/dev/gcp-gke template
- CLI generator support:
  - cloud gcp + orchestrator gke

Rules:
- Provider configured in root.
- No service account keys in repo.
- Keep private cluster advanced options TODO unless simple.
- Mark sensitive outputs.

Run:
- terraform fmt -recursive
- gofmt
- go test ./...
```

---
