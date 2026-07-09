## Prompt 53 — Plan-mode tests for AWS modules

```text
Add plan-mode Terraform tests for AWS modules where possible.

Target modules:
- modules/cloud/aws/network
- modules/cloud/aws/tfstate-backend
- modules/cloud/aws/dns
- modules/cloud/aws/irsa-role

Goal:
Add lightweight tests that validate Terraform plans without requiring real apply.

Create tests:
- modules/cloud/aws/network/tests/basic.tftest.hcl
- modules/cloud/aws/tfstate-backend/tests/basic.tftest.hcl
- modules/cloud/aws/dns/tests/basic.tftest.hcl
- modules/cloud/aws/irsa-role/tests/basic.tftest.hcl

Test strategy:
- Use mocked/safe provider configuration where possible.
- Prefer command = plan.
- Validate that important outputs are present.
- Validate that required resources would be planned.
- Validate input validation failures.

For network module:
- Valid VPC with 2 AZs.
- Invalid subnet length mismatch fails.
- NAT disabled plan works.
- NAT enabled plan works.

For tfstate backend:
- Versioning enabled by default.
- Encryption enabled by default.
- Public access block exists.
- DynamoDB lock table exists.

For DNS:
- Existing zone lookup scenario may be hard to test without real AWS.
- Add test for record object validation.
- Document limitations.

For IRSA:
- Valid trust policy generation.
- Service account subject format is correct.

Rules:
- Do not require real AWS credentials in default CI unless tests are explicitly skipped.
- If provider limitations prevent safe tests, document the limitation.
- Add a test matrix in docs/testing.md.

Run:
- terraform fmt -recursive
- terraform test where possible

Final response:
- List tests added.
- List tests that are skipped and why.
```

---
