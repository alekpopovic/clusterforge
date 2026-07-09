# Supply Chain Security

ClusterForge release artifacts must include SHA256 checksums. SBOMs are
recommended for every CLI release, and artifact signing is planned but not yet
implemented.

## Checksums

```bash
cd dist
sha256sum * > SHA256SUMS
```

## SBOM

Preferred manual command when Syft is installed:

```bash
syft packages cli/cf -o spdx-json > dist/cf.spdx.json
```

If Syft is unavailable, document the tool gap in the release notes.

## Go Dependency Hygiene

```bash
cd cli
go mod tidy
go test ./...
```

Run Go vulnerability checks when available:

```bash
govulncheck ./...
```

Missing optional tools should not break normal local development.
