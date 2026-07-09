## Prompt 101 — Workload cloud identity abstraction

```text
Design and implement workload cloud identity abstraction.

Goal:
Provide a consistent way for app manifests and workload modules to reference cloud IAM identity across orchestrators.

Create RFC:
- docs/rfcs/007-workload-identity.md

Then implement MVP:
- Kubernetes EKS IRSA support in app/worker/cronjob modules.
- ECS task role support in ECS service module.
- Placeholder docs for AKS Workload Identity and GKE Workload Identity.

App manifest extension:
cloud_identity:
  enabled: true
  provider: aws
  policy_arns:
    - arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess
  inline_policies: {}

CLI rendering:
- For Kubernetes on EKS:
  - Generate IRSA role module call.
  - Annotate service account.
- For ECS:
  - Generate task role or reference task_role_arn.
- For unsupported targets:
  - fail with clear message or warn.

Rules:
- Do not create admin policies.
- Do not store credentials.
- Least privilege examples only.
- Identity must be opt-in.

Tests:
- Kubernetes app manifest with cloud_identity renders IRSA.
- ECS app manifest renders task role reference.
- Unsupported provider fails clearly.

Run:
- gofmt
- go test ./...
- terraform fmt -recursive
```

---
