## Prompt 32 — External Secrets Operator integracija

```text
Add Kubernetes External Secrets Operator platform module.

Goal:
Allow ClusterForge users to install External Secrets Operator and configure secret store references without putting secret values in Terraform.

Create module:
- modules/platform/kubernetes/external-secrets

Use Helm:
- repository: https://charts.external-secrets.io
- chart: external-secrets

Inputs:
- namespace: string, default "external-secrets"
- create_namespace: bool, default true
- chart_version: string, default ""
- values: list(string), default []
- labels: map(string), default {}

Outputs:
- namespace
- release_name

Update:
- modules/platform/kubernetes/bootstrap
  Add input:
  - enable_external_secrets: bool, default false

  When enabled, call external-secrets child module.

Add optional module:
- modules/platform/kubernetes/external-secrets/aws-cluster-secret-store

Purpose:
Create a ClusterSecretStore or SecretStore manifest for AWS Secrets Manager or SSM Parameter Store.

Inputs:
- name: string
- region: string
- service: string, default "SecretsManager"
- auth_type: string, default "jwt"
- service_account_ref_name: string
- service_account_ref_namespace: string
- kind: string, default "ClusterSecretStore"

Resources:
- Use kubernetes_manifest if available.
- Keep schema clear.

README:
- Explain that secret values must stay outside Terraform.
- Show AWS Secrets Manager example.
- Show how workload modules reference Kubernetes secrets created by External Secrets.
- Add warnings about Terraform state.

Update docs/security.md:
- Add recommended secret strategy:
  - Cloud secret manager stores values.
  - External Secrets Operator syncs to Kubernetes.
  - Workloads reference Kubernetes Secret keys.
  - Terraform only manages references.

Run:
- terraform fmt -recursive

Final response:
- List modules added.
- Explain how this improves secret handling.
- Mention limitations.
```

---
