## Prompt 87 — Module conformance checker

```text
Implement a module conformance checker.

Goal:
Ensure every Terraform module follows ClusterForge standards.

Create CLI command:
- cf module check
- cf module check --path modules/cloud/aws/network
- cf module check --json

Checks:
- main.tf exists
- variables.tf exists
- outputs.tf exists
- versions.tf exists
- README.md exists
- README has usage section
- README has TF docs markers if terraform-docs is enabled
- variables have descriptions
- outputs have descriptions
- versions.tf has required_version
- required providers are declared when provider resources are used
- no provider configuration inside reusable modules unless allowlisted
- no obvious secrets in examples
- module status is declared in README or MODULE_CATALOG.md

Output:
- pass/warn/fail per module
- JSON support
- non-zero exit when fail

Update:
- Makefile target: make check-modules
- CI workflow to run module check
- docs/module-conformance.md

Rules:
- Do not attempt to parse all Terraform perfectly if simple checks are enough.
- False positives should be warnings where uncertain.
- Keep allowlist documented.

Run:
- gofmt
- go test ./...
```

---
