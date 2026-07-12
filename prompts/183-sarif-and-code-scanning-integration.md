# Prompt 183 — SARIF and code scanning integration

```text
Improve SARIF support for policy and security results.

Goal:
Allow ClusterForge policy results to appear in GitHub code scanning or compatible systems.

CLI:
- cf policy check --format sarif --output policy.sarif
- cf security report --format sarif --output security.sarif

SARIF contents:
- rule id
- title
- severity mapping
- message
- location if known
- remediation text
- help URL if available

Inputs:
- app manifests
- Terraform files
- policy engine results
- module conformance results

Tests:
- SARIF validates structurally
- policy result converted to SARIF
- file location included when known
- no secrets in SARIF

Docs:
- docs/sarif.md
- GitHub code scanning workflow example

Rules:
- Do not overstate severity.
- If location unknown, include result without file location.
- No sensitive values.
```
