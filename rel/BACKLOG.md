# Backlog

| Area | Title | Priority | Status | Complexity | Notes |
| --- | --- | --- | --- | --- | --- |
| CLI | Add full tests for drift/state/cost/template commands | high | todo | medium | Must cover exit codes and JSON output. |
| CLI | Implement executable plugin prototype | medium | planned | high | Requires RFC 006 approval. |
| CLI | Implement import/adopt dry-run commands | medium | planned | high | Start with AWS VPC and Route53. |
| Terraform modules | Harden AKS module | high | experimental | medium | Add node pool variants and tests. |
| Terraform modules | Harden GKE module | high | experimental | medium | Add private cluster options later. |
| Kubernetes | Native sidecars and volumes for app/worker | medium | partial | high | Prefer GitOps for complex pods. |
| Kubernetes | K3s/RKE2 validation | medium | experimental | medium | Requires disposable VMs. |
| ECS | CodeDeploy blue/green smoke test | medium | experimental | medium | Needs ALB service fixture. |
| Nomad | Add real Nomad smoke runbook | low | todo | medium | After Kubernetes/ECS validation. |
| Docker | Add Docker Swarm smoke runbook | low | todo | low | Local-only target. |
| Multi-cloud | Azure smoke test runbook | high | todo | medium | Required before marking AKS beta. |
| Multi-cloud | GCP smoke test runbook | high | todo | medium | Required before marking GKE beta. |
| Security | Add OPA rules for policy packs | high | planned | medium | Convert advisory checks where possible. |
| Docs | Expand tutorials with screenshots/evidence examples | medium | partial | low | Keep secrets redacted. |
| CI/CD | Add release dry-run workflow | medium | planned | medium | Should not publish artifacts. |
| Testing | Real cloud smoke cadence | high | planned | medium | Record dates and testers in matrix. |
| Product | SaaS/API exploration brief | low | planned | high | v0.3 discovery only. |
