## Prompt 2 — Repo skeleton

```text
Create the initial repository skeleton for ClusterForge.

Before making changes:
- Inspect the current directory.
- Preserve any existing files.
- If a file exists, update it carefully instead of replacing it blindly.

Create this structure:

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

cli/
  cmd/
  internal/
    config/
    terraform/
    generator/
    orchestrators/
    policy/
    ui/
  templates/

policies/
  conftest/
  checkov/

scripts/

.github/
  workflows/

For every Terraform module directory, create:
- main.tf
- variables.tf
- outputs.tf
- versions.tf
- README.md

For now, modules may contain placeholders, but they must be valid enough that future tasks can fill them in.

Also create:
- .gitignore
- .editorconfig
- README.md
- LICENSE placeholder
- scripts/lint.sh
- scripts/validate.sh

.gitignore must exclude:
- .terraform/
- *.tfstate
- *.tfstate.*
- *.tfplan
- crash.log
- .terraform.lock.hcl only if we decide not to commit it; otherwise add a comment explaining root modules may commit it
- kubeconfig files
- .env files
- CLI build artifacts

README.md must include:
- What ClusterForge is
- Architecture layers
- Supported orchestrators
- Planned CLI workflow
- Current status: early development

Do not implement real cloud resources yet. Focus on clean structure.
Run a tree/listing at the end and summarize what you created.
```

---
