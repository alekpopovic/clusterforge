# RFC 007: Workload cloud identity

Status: accepted for MVP

## Context

Workloads need cloud API access without static access keys. ClusterForge needs a
portable manifest shape while keeping provider-specific trust visible in the
generated Terraform.

## Decision

Applications may opt in with `cloud_identity`:

```yaml
cloud_identity:
  enabled: true
  provider: aws
  policy_arns:
    - arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess
  inline_policies: {}
```

The AWS MVP renders an `irsa-role` module and annotated service account for EKS.
For ECS it configures the service module's task role, or accepts an existing
`task_role_arn`. Policies contain permissions, never credentials. Administrator
policy references are rejected, and identity is disabled unless enabled.

The Kubernetes app, worker, and cronjob modules support annotated service
accounts. Worker and cronjob callers set `service_account_name` and
`service_account_annotations`; the app module uses its `service_account` object.

## Unsupported targets

AKS Workload Identity and GKE Workload Identity use the same conceptual manifest
but are placeholders in this MVP. Rendering an enabled AWS identity for AKS,
GKE, generic Kubernetes, K3s, or RKE2 fails clearly. No credentials are generated
as a fallback.

## Security and operations

- Use workload-specific, least-privilege policies and resource constraints.
- Prefer managed policy references or reviewed JSON policy documents.
- Treat changes to trust relationships and policy documents as security changes.
- Review generated Terraform before apply. ClusterForge cannot prove that a
  custom policy is least privilege.
- Existing ECS roles remain supported through `task_role_arn`; ClusterForge does
  not modify policies on externally managed roles.

## Future work

AKS support should federate a Kubernetes service account with a managed identity.
GKE support should bind a Kubernetes service account to a Google service account.
Both require provider-specific modules and tests before rendering is enabled.
