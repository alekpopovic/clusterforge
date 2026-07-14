# Prompt 224 — Pull request environment preview RFC

```text
Create preview environment RFC.

Create:
- docs/rfcs/023-preview-environments.md
- docs/preview-environments.md

Goal:
Design ephemeral environments for pull requests.

Use cases:
- app preview
- integration testing
- QA review
- temporary namespace
- temporary ECS service

Targets:
- existing Kubernetes namespace
- EKS shared dev cluster namespace
- ECS temporary service
- local kind not relevant for remote PR

Design:
1. namespace-per-PR
2. app-per-PR
3. TTL-based cleanup
4. DNS naming
5. secrets strategy
6. cost controls
7. policy restrictions
8. approval optional

CLI future:
- cf preview create
- cf preview list
- cf preview delete
- cf preview cleanup

Control Plane future:
- preview environments table
- PR webhook creates preview request
- runner deploys preview
- TTL cleanup job

Risks:
- cost explosion
- leaked secrets
- stale previews
- public exposure
- database migrations per PR
- tenant isolation

Do not implement code yet.
Update roadmap.
```
