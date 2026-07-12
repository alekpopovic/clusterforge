# Prompt 168 — Runner architecture RFC

```text
Create the remote runner architecture RFC.

Create:
- docs/rfcs/017-remote-runner.md
- docs/remote-runner.md

Goal:
Design how ClusterForge runs plan/apply jobs outside the user laptop.

Runner responsibilities:
- poll control plane for jobs
- clone repository
- checkout branch/commit
- install/use ClusterForge CLI
- run policy checks
- run terraform init/plan/apply
- upload sanitized outputs
- upload plan summary
- upload audit events
- never upload secrets
- cleanup workspace

Job types:
- validate
- policy_check
- plan
- drift_check
- cost_scan
- apply
- destroy only if explicitly allowed

Security model:
- runner token
- least privilege cloud credentials
- repository access
- isolated workspace per job
- no state storage in API
- plan artifacts may contain sensitive data and must be protected
- apply requires approval

Deployment models:
1. local runner
2. Kubernetes runner
3. VM runner
4. CI runner integration

Non-goals:
- arbitrary command execution
- unreviewed apply
- storing cloud credentials in control plane
- multi-tenant SaaS isolation in first version

Include:
- sequence diagrams
- threat model
- runner lifecycle
- job status model
- artifact handling
- cleanup strategy

Do not implement runner code yet.
```
