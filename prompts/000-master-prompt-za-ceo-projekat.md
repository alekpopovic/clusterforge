## Prompt 0 — Master prompt za ceo projekat

```text
You are helping me build a production-quality Terraform framework for container orchestrators.

Project name: ClusterForge.

Goal:
Build an opinionated but readable Infrastructure-as-Code framework that can provision and manage container platforms using Terraform or OpenTofu.

Supported targets:
1. Kubernetes:
   - AWS EKS first
   - later AKS, GKE, K3s, RKE2, generic Kubernetes
2. AWS ECS/Fargate
3. Nomad
4. Docker Swarm / Docker Engine as a simple self-hosted target

Architecture:
The framework must be split into four layers:

1. Foundation layer:
   - cloud networking
   - IAM
   - DNS
   - storage
   - registry
   - firewall/security groups

2. Orchestrator layer:
   - EKS
   - ECS
   - Nomad
   - Docker Swarm
   - later AKS/GKE/K3s/RKE2

3. Platform layer:
   - ingress
   - TLS/cert-manager
   - external-dns
   - observability
   - logging
   - secrets
   - Argo CD / GitOps

4. Workload layer:
   - web app
   - worker
   - cronjob
   - service
   - scheduled task
   - Nomad job
   - Docker service/container

Important design rules:
- Do not create one giant "everything" module.
- Each Terraform module must do one thing.
- Provider configuration must stay in root/live environments, not hidden deep inside child modules.
- Each module must include:
  - main.tf
  - variables.tf
  - outputs.tf
  - versions.tf
  - README.md
  - examples if useful
- Use typed variables with validation where appropriate.
- Use clear outputs.
- Use consistent naming, labels and tags.
- Do not put secrets in plain text tfvars.
- Do not auto-apply production changes.
- Terraform/OpenTofu must remain visible and readable. The CLI can generate files, but it must not hide infrastructure logic.

Repository structure target:

clusterforge/
  README.md
  LICENSE
  .gitignore
  .editorconfig
  AGENTS.md

  modules/
    core/
      naming/
      labels/
      tags/

    cloud/
      aws/
        network/
        iam/
        dns/
        storage/

    orchestrators/
      kubernetes/
        eks/
        generic/
      ecs/
        cluster/
      nomad/
        cluster/
      docker/
        engine/
        swarm-service/

    platform/
      kubernetes/
        bootstrap/
        ingress-nginx/
        cert-manager/
        external-dns/
        metrics-server/
        prometheus-stack/
        loki/
        argocd/
      ecs/
        alb/
        cloudwatch/
      nomad/
        consul/
        ingress/

    workloads/
      kubernetes/
        app/
        cronjob/
        helm-app/
      ecs/
        service/
        scheduled-task/
      nomad/
        service/
        batch/
      docker/
        container/
        swarm-service/

  cli/
    go.mod
    main.go
    cmd/
    internal/
    templates/

  live/
    dev/
      aws-eks/
      aws-ecs/
    staging/
      aws-eks/
    prod/
      aws-eks/

  examples/
    kubernetes-basic-app/
    kubernetes-with-ingress/
    ecs-fargate-app/
    nomad-service/
    docker-swarm-service/

  policies/
    conftest/
    checkov/

  scripts/
    lint.sh
    validate.sh
    docs.sh

  .github/
    workflows/
      terraform-validate.yml
      cli-test.yml
      security-scan.yml

Expected behavior:
- Start by inspecting the existing repository.
- If files already exist, do not overwrite them blindly.
- Create small, reviewable changes.
- After each major change, run formatting, validation, and tests where possible.
- Explain what was changed and why.
- Prefer simple maintainable code over clever abstractions.
```

---
