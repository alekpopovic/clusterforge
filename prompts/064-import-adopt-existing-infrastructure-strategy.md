## Prompt 64 — Import/adopt existing infrastructure strategy

```text
Create an adoption strategy for existing infrastructure.

Goal:
Help teams use ClusterForge with existing AWS/Kubernetes/ECS resources.

Create docs:
- docs/adopting-existing-infrastructure.md
- docs/import-strategy.md

CLI design document:
- docs/rfcs/001-import-adoption.md

Cover:
1. Existing VPC adoption
2. Existing EKS cluster adoption
3. Existing Kubernetes cluster adoption
4. Existing ECS cluster adoption
5. Existing Route53 zone adoption
6. Existing Terraform state adoption
7. Non-Terraform infrastructure import

Define future commands:
- cf adopt aws-vpc
- cf adopt eks
- cf adopt ecs
- cf adopt route53-zone
- cf import plan
- cf import generate

For now:
- Do not implement actual import commands unless straightforward.
- Add documentation and RFC only.
- Add examples of Terraform import blocks if applicable.
- Explain risks:
  - accidental replacement
  - naming mismatch
  - missing tags
  - state ownership conflict
  - provider diff noise

Rules:
- Be conservative.
- Tell users to start with read-only data sources when possible.
- Never recommend importing production without backup and plan review.

Final response:
- Summarize adoption strategy.
- List future commands.
```

---
