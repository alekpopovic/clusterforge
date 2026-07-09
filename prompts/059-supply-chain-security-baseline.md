## Prompt 59 — Supply chain security baseline

```text
Add supply chain security baseline for ClusterForge.

Goal:
Improve trust in CLI releases and repository artifacts.

Tasks:
1. Add SBOM generation support for CLI.
   Preferred:
   - syft if available
   - otherwise document manual command

2. Add checksum generation:
   - sha256sum for release artifacts

3. Add dependency review guidance:
   - docs/supply-chain-security.md

4. Add Go dependency hygiene:
   - go mod tidy check
   - govulncheck integration if available
   - document how to run govulncheck

5. Add GitHub Actions:
   - dependency review if available
   - Go vulnerability check if practical

6. Add release artifact policy:
   - every release artifact must have checksum
   - SBOM recommended
   - signing planned or implemented if straightforward

Rules:
- Do not introduce mandatory tools that break local development if missing.
- Scripts should print clear messages when tools are missing.
- Do not claim artifacts are signed unless signing is implemented.

Run:
- cd cli && go mod tidy
- cd cli && go test ./...
- security scripts if available

Final response:
- List supply chain controls added.
- Mention what remains manual.
```

---
