# platform/kubernetes/external-secrets

## Purpose

Installs External Secrets Operator with Helm. External Secrets Operator syncs
secret values from external secret stores, such as AWS Secrets Manager or SSM
Parameter Store, into Kubernetes Secrets.

This module installs the operator only. Secret values must stay outside
Terraform. Use cloud secret managers for values, SecretStore or
ClusterSecretStore resources for references, and workload modules to reference
the resulting Kubernetes Secret keys.

Provider configuration belongs in the root module. This module declares Helm
and Kubernetes provider requirements but does not configure providers.

## Status

Implemented.

## Usage

```hcl
module "external_secrets" {
  source = "../../../modules/platform/kubernetes/external-secrets"

  namespace = "external-secrets"
  labels = {
    "clusterforge.io/managed-by" = "terraform"
  }
}
```

## AWS Secrets Manager Example

Create an IAM role for the External Secrets Operator service account with IRSA,
then create a ClusterSecretStore using the
`aws-cluster-secret-store` submodule.

```hcl
module "aws_secret_store" {
  source = "../../../modules/platform/kubernetes/external-secrets/aws-cluster-secret-store"

  name                          = "aws-secrets"
  region                        = "eu-central-1"
  service                       = "SecretsManager"
  service_account_ref_name      = "external-secrets"
  service_account_ref_namespace = "external-secrets"
}
```

## Workload Secret References

External Secrets Operator creates Kubernetes Secrets. Workload modules should
reference those Kubernetes Secret keys, not raw secret values:

```hcl
module "api" {
  source = "../../../modules/workloads/kubernetes/app"

  name      = "api"
  namespace = "apps"
  image     = "ghcr.io/company/api:1.0.0"

  secret_env = {
    DATABASE_URL = {
      secret_name = "api-secrets"
      secret_key  = "database-url"
    }
  }
}
```

## Notes

- Do not put secret values in Terraform variables, `tfvars`, Helm values, or
  inline manifests.
- Terraform state may contain rendered manifests and references. Keep state
  encrypted and access-controlled.
- Review External Secrets Operator chart and CRD upgrade notes before changing
  versions.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
