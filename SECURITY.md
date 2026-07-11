# Security policy

## Reporting a vulnerability

Do not open a public issue for a suspected vulnerability. Use GitHub's private
security advisory reporting for this repository. If private reporting is not
available, contact a repository maintainer privately through their verified
GitHub profile and ask for a secure reporting channel before sending details.

Include the affected version or commit, impact, minimal reproduction, and
suggested mitigation when known. Never attach live credentials, Terraform
state, kubeconfigs, private keys, or customer data. Replace sensitive values
with synthetic examples.

Maintainers should acknowledge a report, assess severity, coordinate a fix and
release, and publish an advisory when appropriate. Response times are
best-effort until a formal security response team and SLA are established.

## Supported versions

ClusterForge is pre-1.0. Security fixes target the current development branch
and the latest tagged release when practical. Older releases may require an
upgrade. See `VERSION_MATRIX.md` for tool and provider support.

## Security posture

ClusterForge does not claim compliance certification or guarantee that a
generated configuration is secure for every environment. Operators remain
responsible for reviewing plans, IAM, network exposure, state storage, provider
and module sources, and Kubernetes permissions.

See the [threat model](docs/security-threat-model.md) and operational
[security checklist](docs/security-checklist.md). Existing controls and planned
work are explicitly separated there.

In particular, release signing and complete build provenance, sandboxed or
signed template packs, continuous entitlement analysis, and automatic security
remediation are not currently implemented. Checksums, SBOM support, secret
scanning, plan/destroy safeguards, and policy modules are useful controls but
do not replace environment-specific review.
