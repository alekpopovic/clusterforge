## Prompt 50 — Final MVP packaging

```text
Prepare ClusterForge MVP release candidate.

Goal:
Make the repository ready for an initial v0.1.0 release.

Tasks:
1. Review README:
   - Clear quickstart.
   - Current supported targets.
   - Clear warnings for production.
   - CLI usage.
   - Module examples.

2. Review docs:
   Required docs:
   - docs/architecture.md
   - docs/cli.md
   - docs/app-manifest.md
   - docs/environments.md
   - docs/backends.md
   - docs/security.md
   - docs/gitops.md
   - docs/roadmap.md

3. Review CLI:
   - cf version works.
   - cf project init works.
   - cf env create works.
   - cf generate works for aws+eks and aws+ecs.
   - cf app add/list/validate/render works.
   - cf doctor works.
   - cf plan/apply safety rules are documented.

4. Review Terraform:
   - Core modules implemented.
   - AWS network implemented.
   - EKS implemented enough for MVP.
   - ECS cluster/service implemented enough for MVP.
   - Kubernetes app/worker/cronjob implemented.
   - Platform bootstrap implemented.
   - Placeholder modules clearly marked.

5. Add CHANGELOG.md:
   - v0.1.0 unreleased section.
   - Added/Changed/Security/Known limitations.

6. Add VERSION file:
   - 0.1.0

7. Add CONTRIBUTING.md:
   - Development setup.
   - Testing.
   - Module rules.
   - CLI rules.
   - Security rules.

8. Add release checklist:
   - docs/release-checklist.md

9. Run:
   - make fmt
   - make lint
   - make test
   - make validate
   - make security where tools are available
   - cd cli && go build -o cf .

10. Create FINAL_MVP_REPORT.md:
   Include:
   - What is included in v0.1.0
   - What is not included
   - Known limitations
   - Commands run
   - Pass/fail status
   - Recommended next version goals

Rules:
- Do not pretend cloud apply was tested unless it actually was.
- Be explicit about validation limitations.
- Do not add credentials.
- Do not remove useful TODOs; organize them.

Final response:
- Summarize release candidate status.
- State whether v0.1.0 is ready or what blocks it.
```

---

## Preporučeni redosled za prvih 50

```text
0.  Master prompt
1.  AGENTS.md
2.  Repo skeleton
3.  Terraform standards
4.  core/naming
5.  core/tags + core/labels
6.  AWS network
7.  EKS
8.  live/dev/aws-eks
9.  Kubernetes platform bootstrap
10. Kubernetes app workload
11. Kubernetes cronjob
12. ECS cluster
13. ECS service
14. Nomad service
15. Docker modules
16. CLI base
17. CLI config
18. CLI Terraform runner
19. CLI generator
20. App manifest generator
21. Policy/risk summary
22. CI
23. Docs
24. Tests/examples
25. Final review
26. Audit
27. Makefile
28. Terraform validation
29. terraform-docs
30. security scanning
31. EKS OIDC/IRSA
32. External Secrets
33. Argo CD GitOps
34. Node autoscaling
35. ECS ALB
36. Route53 DNS
37. cert-manager issuer
38. CLI release/install
39. Interactive wizard
40. App manifest validation
41. Multi-stack layout
42. Remote backend generator
43. AWS tfstate backend
44. Observability
45. Kubernetes worker
46. Helm app wrapper
47. CLI doctor
48. CLI JSON output
49. Pod Security / NetworkPolicy
50. Final MVP packaging
```
