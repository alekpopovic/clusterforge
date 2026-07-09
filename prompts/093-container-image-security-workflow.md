## Prompt 93 — Container image security workflow

```text
Add container image security workflow documentation and optional checks.

Goal:
Define how ClusterForge users should scan and verify container images.

Create:
- docs/image-security.md
- .github/workflows/image-security-example.yml

Cover:
1. Image scanning:
   - Trivy
   - Grype optional
2. SBOM:
   - Syft optional
3. Signing:
   - Cosign optional
4. Provenance:
   - SLSA/GitHub provenance discussion
5. Registry:
   - ECR immutable tags
   - avoid latest in production
6. Workload policy:
   - reject latest tag in prod
   - require pinned tags or digests
   - optional signature verification later

CLI:
- Add policy check:
  - warn if app manifest image uses :latest
  - warn if prod app image is not pinned by tag or digest
  - optional blocking policy in production pack

Tests:
- latest tag warning
- digest accepted
- normal version tag accepted
- missing image fails

Rules:
- Do not require signing by default.
- Do not claim verification exists unless implemented.
- Keep policy configurable.

Run:
- gofmt
- go test ./...
```

---
