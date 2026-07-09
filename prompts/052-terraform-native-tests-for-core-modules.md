## Prompt 52 — Terraform native tests for core modules

```text
Add Terraform native tests for ClusterForge core modules.

Target modules:
- modules/core/naming
- modules/core/tags
- modules/core/labels

Goal:
Use Terraform test files to validate module behavior.

Create:
- modules/core/naming/tests/naming.tftest.hcl
- modules/core/tags/tests/tags.tftest.hcl
- modules/core/labels/tests/labels.tftest.hcl

Test cases for naming:
- basic name generation
- lowercase conversion
- max length truncation
- extra parts
- suffix
- DNS-safe name
- labels-safe name
- invalid separator should fail
- empty required values should fail

Test cases for tags:
- required tags exist
- optional tags are omitted when empty
- extra tags merge correctly
- override behavior matches README

Test cases for labels:
- app.kubernetes.io labels exist when inputs provided
- clusterforge.io labels exist
- extra labels merge correctly
- generated labels are Kubernetes-compatible

Rules:
- Use plan-mode tests where possible.
- Do not create cloud resources.
- Keep tests fast.
- Update Makefile:
  - make test-terraform
- Update scripts/validate.sh to optionally run terraform test for safe modules.
- Update docs/testing.md.

Run:
- terraform fmt -recursive
- terraform test in each core module

Final response:
- List test files added.
- List test results.
- Mention any Terraform version requirement.
```

---
