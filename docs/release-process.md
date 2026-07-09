# Release Process

ClusterForge releases are driven by Git tags matching `v*`.

```bash
git tag v0.1.0
git push origin v0.1.0
```

The release workflow checks formatting, runs CLI tests, builds binaries for
Linux, macOS, and Windows, generates `SHA256SUMS`, extracts release notes from
`CHANGELOG.md` when possible, and attaches artifacts to a GitHub release.

Do not publish Docker images yet. Do not use cloud credentials in the release
workflow. Cosign signing and SBOM upload are tracked as follow-up work.
