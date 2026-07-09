## Prompt 16 — CLI osnova u Go + Cobra

```text
Implement the initial ClusterForge CLI in cli/.

Language:
- Go

CLI framework:
- Cobra

Binary name:
- cf

Commands to implement initially:
- cf version
- cf project init <name>
- cf env create <name>
- cf env list
- cf generate <env>
- cf init <env>
- cf plan <env>
- cf apply <env>
- cf destroy <env>
- cf doctor

Repository layout under cli/:
cli/
  go.mod
  main.go
  cmd/
    root.go
    version.go
    project.go
    env.go
    generate.go
    init.go
    plan.go
    apply.go
    destroy.go
    doctor.go
  internal/
    config/
      schema.go
      loader.go
      writer.go
    terraform/
      runner.go
    generator/
      generator.go
    policy/
      checks.go
    ui/
      printer.go

Behavior:
- cf project init <name> creates:
  - clusterforge.yaml
  - apps/
  - live/
  - .cf/
- Do not overwrite existing project files unless a --force flag is used.
- cf env create creates an environment entry in clusterforge.yaml and a matching live directory.
- cf env list prints environments from config.
- cf init/plan/apply/destroy run Terraform/OpenTofu in the environment path.
- Default engine is terraform.
- Support --engine terraform|tofu.
- Support --config clusterforge.yaml.
- Support --verbose.

clusterforge.yaml schema:
project:
  name: string
  default_engine: terraform

engines:
  terraform:
    binary: terraform
  opentofu:
    binary: tofu

defaults:
  cloud: aws
  region: eu-central-1
  orchestrator: eks

environments:
  dev:
    cloud: aws
    region: eu-central-1
    orchestrator: eks
    path: live/dev/aws-eks

policies:
  require_plan_file_for_apply: true
  block_destroy_in_prod: true
  require_manual_approval_for_prod: true

Safety:
- apply against prod must require --plan-file.
- destroy against prod must be blocked unless --allow-destroy is passed.
- Never add --auto-approve by default.
- Print clear warnings for destructive commands.

Implementation quality:
- Add useful errors.
- Add unit tests for config loader and policy checks.
- Use gofmt.
- Use go test ./...

At the end:
- Show example commands.
- Summarize implemented files.
```

---
