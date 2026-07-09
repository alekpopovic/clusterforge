# ClusterForge Context

ClusterForge is a Terraform/OpenTofu framework and Go CLI for container
platforms. It supports Kubernetes, ECS, Nomad, and Docker targets, with AWS EKS
and ECS as the most mature paths.

## Important Paths

- `cli/`: Go CLI using Cobra.
- `cli/internal/`: CLI business logic.
- `cli/templates/`: generated environment templates.
- `modules/`: reusable Terraform/OpenTofu modules.
- `live/`: real environment roots.
- `examples/`: copy-paste friendly examples.
- `docs/`: user and maintainer documentation.
- `policies/`: static policy packs and scanner guidance.
- `scripts/`: repeatable validation and release helpers.
- `prompts/`: project prompt backlog and implementation history.

## Validation Shortcuts

```bash
make fmt-check
make lint
make test-cli
make validate
make security
```

Use `TERRAFORM_BIN=tofu` when validating with OpenTofu.
