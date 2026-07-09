## Prompt 84 — Provider compatibility matrix CI

```text
Add provider compatibility matrix validation.

Goal:
Catch compatibility issues across supported Terraform/OpenTofu and provider versions.

Create:
- docs/provider-compatibility.md
- .github/workflows/provider-compatibility.yml

Matrix dimensions:
- Terraform versions supported by VERSION_MATRIX.md
- OpenTofu versions supported by VERSION_MATRIX.md
- Core provider sets:
  - AWS stack
  - Kubernetes/Helm stack
  - Azure stack if implemented
  - GCP stack if implemented
  - Nomad stack if implemented
  - Docker stack if implemented

CI behavior:
- Run terraform fmt check.
- Run terraform init -backend=false where safe.
- Run terraform validate where safe.
- Do not require cloud credentials.
- Skip real cloud data source validation if credentials are required.
- Print clear skip reasons.

Docs:
- Explain tested vs supported.
- Explain provider version constraints.
- Explain how to run compatibility tests locally.

Rules:
- Do not make CI flaky.
- Do not require paid cloud resources.
- Do not test every possible version combination if it becomes too slow.
- Prefer focused matrix.

Final response:
- List matrix coverage.
- List known skipped areas.
```

---
