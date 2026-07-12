# Release Process

ClusterForge releases are driven by Git tags matching `v*`.

```bash
git tag v0.1.0
git push origin v0.1.0
```

The release workflow checks formatting, runs CLI tests, builds binaries for
Linux, macOS, and Windows, generates `SHA256SUMS`, extracts release notes from
`CHANGELOG.md` when possible, and attaches artifacts to a GitHub release. Stable
asset names (`cf-linux-amd64`, for example), their individual `.sha256` files,
and `install.sh` support checksum-verified curl installation. Test both `latest`
and a pinned tag URL before announcing a release.

Do not publish Docker images yet. Do not use cloud credentials in the release
workflow. Cosign signing and SBOM upload are tracked as follow-up work.
