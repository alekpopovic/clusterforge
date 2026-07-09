## Prompt 58 — GitHub release automation for v0.1.x

```text
Implement release automation for ClusterForge.

Goal:
Create a repeatable release process for CLI binaries, checksums, changelog and module release notes.

Create/update:
- .github/workflows/release.yml
- scripts/release-notes.sh
- docs/release-process.md
- docs/release-checklist.md

Release workflow:
Trigger:
- tag push matching v*

Steps:
1. Checkout.
2. Run tests.
3. Build CLI binaries:
   - linux amd64
   - linux arm64
   - darwin amd64
   - darwin arm64
   - windows amd64
4. Generate checksums.
5. Attach artifacts to GitHub release.
6. Include CHANGELOG.md section if practical.

Add:
- Cosign/SBOM TODO section, but do not implement unless tools are already configured.

Release checklist:
- Version updated.
- Changelog updated.
- Acceptance report updated.
- Smoke test matrix updated.
- CLI tests pass.
- Terraform fmt passes.
- Security scan reviewed.
- Docs reviewed.
- No secrets committed.

Rules:
- Do not publish Docker images yet.
- Do not require cloud credentials.
- Do not create real releases during this task.
- Workflow should be ready for future tag push.

Final response:
- List release workflow behavior.
- Show release command:
  git tag v0.1.0
  git push origin v0.1.0
```

---
