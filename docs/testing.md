# Testing

ClusterForge uses a mix of Go unit tests, Terraform validation, and Terraform
native tests.

## CLI Tests

Run CLI tests and a build check:

```bash
make test-cli
```

This runs:

- `go mod download`
- `gofmt` check
- `go test ./...`
- `go build`

## Terraform Native Tests

Provider-free core modules use Terraform native tests under each module's
`tests/` directory:

- `modules/core/naming/tests/naming.tftest.hcl`
- `modules/core/tags/tests/tags.tftest.hcl`
- `modules/core/labels/tests/labels.tftest.hcl`

Run all safe Terraform native tests:

```bash
make test-terraform
```

Run one module directly:

```bash
terraform -chdir=modules/core/naming test -no-color
terraform -chdir=modules/core/tags test -no-color
terraform -chdir=modules/core/labels test -no-color
```

These tests use plan mode and create no cloud resources.

## Terraform Test Matrix

| Area | Modules | Command | Default CI | Notes |
| --- | --- | --- | --- | --- |
| Core metadata | `modules/core/naming`, `modules/core/tags`, `modules/core/labels` | `make test-terraform` | Yes, through `make validate` | Provider-free, fast, and safe. |
| AWS foundation | `modules/cloud/aws/network`, `modules/cloud/aws/tfstate-backend`, `modules/cloud/aws/dns`, `modules/cloud/aws/irsa-role` | `make test-terraform-aws` | No | Plan-mode tests only. Requires working AWS provider plugins but does not run `apply` or require real AWS credentials. |

AWS module tests use fake provider configuration or Terraform mock providers
where practical. They are intentionally not part of default CI because provider
plugin installation and startup can be unavailable in restricted environments.

The Route53 DNS tests cover created-zone and record validation paths. Existing
hosted-zone lookup is not covered by default because `data.aws_route53_zone`
requires real AWS API access.

## Repository Validation

Run repository validation:

```bash
make validate
```

This runs Terraform formatting checks, module file-contract checks, safe root
validation where provider plugins are available, and Terraform native tests for
the provider-free core modules.

Disable Terraform native tests during validation:

```bash
RUN_TERRAFORM_TESTS=false make validate
```

## Full Local Test Pass

Run the default local test target:

```bash
make test
```

This combines CLI tests with repository Terraform validation.

## Limitations

Some modules and examples depend on Terraform providers such as AWS,
Kubernetes, Helm, Docker, or Nomad. Validation for those roots may be skipped in
restricted or offline environments when provider plugins cannot be installed or
started. Skips should be explicit in the validation output.

No default test target runs `terraform apply` or requires cloud credentials.
