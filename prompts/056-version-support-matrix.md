## Prompt 56 — Version support matrix

```text
Add an explicit version support matrix for ClusterForge.

Create:
- docs/version-support.md
- VERSION_MATRIX.md

Cover:
1. ClusterForge CLI version
2. Terraform versions
3. OpenTofu versions
4. Go versions for CLI development
5. Kubernetes supported minor versions
6. AWS provider versions
7. Kubernetes provider versions
8. Helm provider versions
9. Nomad provider versions
10. Docker provider versions
11. Tested orchestrators:
    - EKS
    - ECS
    - existing Kubernetes
    - K3s
    - RKE2
    - Nomad
    - Docker Swarm

Add policy:
- The framework supports the currently maintained Kubernetes minor versions unless a module explicitly documents otherwise.
- Provider versions must be pinned with compatible constraints.
- Examples must state tested versions.
- Release notes must mention version changes.

Update:
- README.md
- docs/roadmap.md
- cli doctor command:
  - warn when Terraform/OpenTofu version is below supported minimum
  - warn when Kubernetes version is outside supported matrix if kubectl is available

Rules:
- Do not hardcode future version promises.
- Distinguish:
  - supported
  - tested
  - experimental
  - deprecated

Final response:
- List version docs updated.
- Mention current support policy.
```

---
