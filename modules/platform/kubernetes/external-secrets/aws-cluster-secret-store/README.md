# platform/kubernetes/external-secrets/aws-cluster-secret-store

## Purpose

Creates an External Secrets Operator `ClusterSecretStore` or `SecretStore`
manifest for AWS Secrets Manager or SSM Parameter Store.

This module manages references only. It must not contain secret values.

## Status

Implemented.

## Usage

```hcl
module "aws_secret_store" {
  source = "../../../../modules/platform/kubernetes/external-secrets/aws-cluster-secret-store"

  name                          = "aws-secrets"
  region                        = "eu-central-1"
  service                       = "SecretsManager"
  service_account_ref_name      = "external-secrets"
  service_account_ref_namespace = "external-secrets"
}
```

## Notes

- Use `SecretsManager` for AWS Secrets Manager.
- Use `ParameterStore` for AWS SSM Parameter Store.
- `auth_type = "jwt"` is intended for IRSA-based authentication.
- Terraform state contains the rendered manifest. Keep state encrypted and
  access-controlled.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
