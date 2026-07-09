# RFC 004: Self-Hosted Kubernetes

ClusterForge should support K3s and RKE2 as experimental self-hosted targets.

MVP approach:

- do not provision servers inside the K3s/RKE2 modules
- generate install commands and cloud-init snippets
- output kubeconfig retrieval notes, not kubeconfig credentials
- keep SSH/provisioners out of Terraform by default

Future work can add cloud-specific VM modules that consume the generated user
data.

CLI examples:

```bash
cf env create dev --cloud local --orchestrator k3s
cf env create dev --cloud local --orchestrator rke2
cf generate dev
```

Production caveats: HA control planes, backup, upgrades, CNI, storage, and
certificate rotation remain operator responsibilities.
