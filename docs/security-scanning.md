# Security Scanning

ClusterForge uses static security scanners that can run without cloud
credentials. These checks inspect Terraform and repository configuration; they
do not authenticate to cloud providers, run `terraform plan`, or apply
infrastructure.

## Local Commands

Run all available local security scanners:

```bash
make security
```

or:

```bash
./scripts/security.sh
```

Run formatting, Terraform linting, and Go linting:

```bash
make lint
```

## Tools

### TFLint

Configuration lives in `.tflint.hcl`.

`scripts/lint.sh` runs TFLint recursively when `tflint` is installed. The
configuration enables the base Terraform recommended rules and the AWS ruleset.
These checks are static and should not require AWS credentials.

If TFLint is not installed, linting prints a warning and continues with the
other available checks.

### Checkov

Configuration lives in `.checkov.yml`.

Checkov scans Terraform source while excluding generated/cache paths such as
`.terraform`, `.git`, `.cf`, `vendor`, `node_modules`, `dist`, `bin`, and
`coverage`.

There are no broad rule suppressions in the root Checkov config. If a rule is
intentionally not applicable to an example or placeholder, document the reason
and suppress narrowly near the relevant code or in a specific policy decision.

### Trivy

Configuration lives in `trivy.yaml`.

Trivy runs config scanning with `HIGH` and `CRITICAL` severities configured to
fail the scan. Generated and cache directories are skipped.

## CI

`.github/workflows/security-scan.yml` runs Checkov and Trivy on pull requests
and pushes to `main`. The workflow does not require cloud credentials.

## Limitations

Static scanning cannot prove that live cloud resources are secure. It does not
validate runtime IAM permissions, Kubernetes admission controls, Helm chart
behavior, or provider-side defaults. Treat it as an early quality gate before
reviewed plans and environment-specific checks.
