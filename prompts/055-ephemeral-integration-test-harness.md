## Prompt 55 — Ephemeral integration test harness

```text
Design and implement an optional ephemeral integration test harness.

Goal:
Allow maintainers to run expensive real infrastructure tests explicitly, never by default.

Create:
- tests/integration/
  aws-eks/
  aws-ecs/
  existing-kubernetes/
- scripts/integration-test.sh
- docs/testing-integration.md

Behavior:
- Integration tests are opt-in.
- Require environment variable:
  CLUSTERFORGE_RUN_INTEGRATION_TESTS=true
- Require explicit target:
  scripts/integration-test.sh aws-eks
  scripts/integration-test.sh aws-ecs
  scripts/integration-test.sh existing-kubernetes

AWS EKS test should:
- Create a generated test project in a temp directory.
- Use unique names with timestamp/random suffix.
- Run init, plan, apply.
- Verify outputs.
- Optionally run kubectl get nodes.
- Destroy everything at the end.
- Always attempt cleanup on failure.

AWS ECS test should:
- Generate ECS environment.
- Plan/apply minimal ECS cluster and service if feasible.
- Destroy everything.

Existing Kubernetes test should:
- Require KUBECONFIG.
- Deploy a simple nginx/httpbin workload using workloads/kubernetes/app.
- Verify deployment and service.
- Destroy test resources.

Rules:
- Never run from default CI.
- Never run without explicit env var.
- Print cost warning.
- Use traps for cleanup.
- Do not store tfstate in repo.
- Do not use production account names.

Final response:
- Explain integration harness.
- List commands.
- Mention safety gates.
```

---
