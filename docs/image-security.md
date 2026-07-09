# Container Image Security

ClusterForge separates infrastructure generation from image build and release
controls. Use this workflow as a baseline for application repositories that
publish images consumed by ClusterForge app manifests.

## Image Scanning

Use Trivy as the default scanner in CI before pushing or deploying images.
Grype is a reasonable optional second scanner when teams want another advisory
source.

Scanning should fail builds for critical issues only after the team has agreed
on severity thresholds, ignore policy, and fix windows.

## SBOM

Generate an SBOM with Syft when the application delivery process needs package
inventory evidence. Store SBOM artifacts with the build or release record.

## Signing

Cosign signing is recommended for mature production pipelines, but signature
verification is not enforced by ClusterForge by default today. Do not mark an
environment as signature-verified unless admission control or deploy-time
verification is actually configured.

## Provenance

GitHub artifact attestations and SLSA provenance can help connect an image to
the workflow that built it. Treat provenance as evidence to retain with release
artifacts. It is not a replacement for vulnerability scanning or deploy policy.

## Registry

Use ECR immutable tags for production repositories. Avoid `latest` in
production manifests because it is not repeatable. Prefer version tags such as
`1.4.2` or immutable digests such as:

```text
123456789012.dkr.ecr.us-east-1.amazonaws.com/api@sha256:<digest>
```

Do not put real account IDs in shared examples or documentation.

## Workload Policy

ClusterForge warns when an app manifest uses `:latest`. During `cf app render
--env prod`, it also warns when an image is not pinned by a tag or digest.

Current behavior is warning-only. Future production policy packs may block
`latest`, require digests, or verify signatures, but those controls must be
enabled explicitly and tested before maintainers claim enforcement.
