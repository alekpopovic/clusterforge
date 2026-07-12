# Prompt 197 — Container signing and provenance plan

```text
Add container signing and provenance plan.

Create:
- docs/container-signing-provenance.md
- docs/supply-chain-security.md update

Cover:
- image tags
- immutable digests
- SBOM
- cosign signing
- GitHub Actions provenance
- SLSA concepts
- verification in Kubernetes admission later
- release process

Optional implementation:
- add release workflow steps for:
  - SBOM generation
  - checksum generation
  - image signing if cosign configured
  - provenance if GitHub supports it in project setup

Rules:
- Do not claim signing is active unless workflow implements it.
- Do not require signing for local development.
- No private keys in repo.
- Prefer keyless signing if supported by environment.

Final response:
- State implemented vs documented.
```
