# Module Conformance

ClusterForge modules are checked with:

```bash
cf module check
```

Check a single module:

```bash
cf module check --path modules/cloud/aws/network
```

Emit machine-readable output:

```bash
cf module check --json
```

The Makefile target used by CI is:

```bash
make check-modules
```

## Checks

Hard failures:

- Required files are missing: `main.tf`, `variables.tf`, `outputs.tf`,
  `versions.tf`, or `README.md`.
- `README.md` has no usage section or module example.
- `versions.tf` has no `required_version`.
- A reusable module configures a provider block.
- Examples contain obvious credential-like values.

Warnings:

- `terraform-docs` markers are missing while `.terraform-docs.yml` exists.
- Variable or output descriptions are missing.
- Provider requirements are uncertain from a simple resource-prefix scan.
- Module status is not declared in `README.md` or `MODULE_CATALOG.md`.

The checker intentionally uses simple scans instead of a full Terraform parser.
Uncertain checks should remain warnings unless they become a clear repository
contract violation.

## Allowlist

Reusable modules should not configure providers. No provider configuration
allowlist is active today. If a future module requires one, document the module
path and reason here before changing the checker behavior.
