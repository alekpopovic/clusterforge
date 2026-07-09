## Prompt 119 — Security threat model

```text
Create a security threat model for ClusterForge.

Create:
- docs/security-threat-model.md

Cover assets:
- Terraform state
- cloud credentials
- kubeconfig files
- generated Terraform
- app manifests
- CLI audit logs
- release artifacts
- CI credentials
- module sources
- template packs
- policy packs

Threats:
- leaked tfstate
- committed secrets
- malicious template pack
- compromised CLI binary
- provider supply chain risk
- over-permissive IAM
- accidental production destroy
- drift hiding malicious changes
- Kubernetes privilege escalation
- public ingress exposure
- DNS takeover/misconfiguration

Controls:
- .gitignore
- secret scanning
- policy packs
- plan file requirement
- destroy block
- remote state encryption
- least privilege IAM
- workload identity
- release checksums/SBOM
- CI isolation
- audit log

Create:
- docs/security-checklist.md

Rules:
- Be honest about controls that are not implemented.
- Mark planned controls separately.
- Do not overclaim compliance.

Update:
- README security section
- SECURITY.md
```

---
