## Prompt 70 — K3s and RKE2 self-hosted Kubernetes support

```text
Add self-hosted Kubernetes support design and MVP.

Goal:
Support K3s and RKE2 as lightweight/self-hosted Kubernetes targets.

Create RFC:
- docs/rfcs/004-self-hosted-kubernetes.md

Then implement MVP modules if feasible:
- modules/orchestrators/kubernetes/k3s
- modules/orchestrators/kubernetes/rke2

Design options:
1. Provision servers using an existing cloud VM module.
2. Install K3s/RKE2 using cloud-init/user_data.
3. Output kubeconfig.

For MVP:
- Keep cloud-specific server provisioning out of this module unless already implemented.
- Accept list of server connection details or user_data generation mode.
- Prefer generating install scripts/cloud-init snippets instead of SSH provisioning.

Inputs:
- cluster_name
- environment
- server_count
- k3s_version or rke2_version default ""
- install_channel default "stable"
- cluster_cidr optional
- service_cidr optional
- disable_components list(string)
- tls_san list(string)
- labels/tags

Outputs:
- install_command
- server_user_data
- agent_user_data
- kubeconfig_retrieval_notes

Create examples:
- examples/k3s-cloud-init
- examples/rke2-cloud-init

CLI:
- cf env create dev --cloud local --orchestrator k3s
- cf generate dev

Rules:
- Do not SSH from Terraform by default.
- Do not store kubeconfig with credentials in repo.
- Document production caveats.
- Keep self-hosted support experimental.

Run:
- terraform fmt -recursive
- gofmt
- go test ./...
```

---
