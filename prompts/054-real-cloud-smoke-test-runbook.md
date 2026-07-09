## Prompt 54 — Real cloud smoke test runbook

```text
Create a real cloud smoke test runbook.

Goal:
Document the exact manual process for testing ClusterForge against a real AWS account without making it part of default CI.

Create:
- docs/smoke-tests/aws-eks.md
- docs/smoke-tests/aws-ecs.md
- docs/smoke-tests/kubernetes-existing-cluster.md
- docs/smoke-tests/cleanup.md
- SMOKE_TEST_MATRIX.md

AWS EKS smoke test must include:
1. Required tools:
   - terraform or tofu
   - aws cli
   - kubectl
   - helm
   - cf CLI

2. Required permissions:
   - VPC
   - IAM
   - EKS
   - EC2
   - CloudWatch
   - Route53 optional

3. Steps:
   - create isolated test environment
   - run cf project init
   - run cf env create
   - run cf generate
   - configure tfvars
   - run cf init
   - run cf plan
   - run cf apply with plan file
   - verify cluster exists
   - verify node group ready
   - verify kubectl access
   - install platform bootstrap
   - deploy demo app
   - destroy test environment

4. Expected outputs:
   - cluster name
   - kubeconfig command
   - ingress status
   - app endpoint if available

5. Cleanup:
   - app deletion
   - platform deletion
   - cluster deletion
   - VPC deletion
   - manual cleanup checklist for stuck load balancers or ENIs

6. Cost warning:
   - EKS, NAT Gateway, load balancers and logs can cost money.

SMOKE_TEST_MATRIX.md:
- Provider
- Orchestrator
- Status
- Last tested date
- Tester
- Terraform/OpenTofu version
- Kubernetes version
- Result
- Notes

Rules:
- Do not put real account IDs.
- Do not put real credentials.
- Do not claim a smoke test passed unless it was actually run.
- Add placeholders for test evidence.

Final response:
- List docs created.
- Explain how maintainers should run smoke tests.
```

---
