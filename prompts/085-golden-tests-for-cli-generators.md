## Prompt 85 — Golden tests for CLI generators

```text
Add golden/snapshot tests for ClusterForge CLI generators.

Goal:
Ensure generated Terraform files remain stable and readable.

Target:
- cli/internal/generator
- env generation
- app rendering
- backend generation
- template packs

Create:
- cli/testdata/golden/
  aws-eks-simple/
  aws-ecs-simple/
  existing-kubernetes/
  app-kubernetes/
  app-ecs/
  backend-s3/
  template-pack-override/

Test behavior:
- Generate into temp directory.
- Compare generated files against golden files.
- Provide update mechanism:
  go test ./... -update-golden
  or documented env var CLUSTERFORGE_UPDATE_GOLDEN=true

Tests:
- aws+eks generation
- aws+ecs generation
- existing+kubernetes generation
- Kubernetes app render
- ECS app render
- S3 backend generation
- template pack override

Rules:
- Golden files must be readable Terraform.
- Do not include timestamps unless stable.
- Do not include absolute paths.
- Do not include machine-specific data.
- Keep update flow explicit.

Docs:
- docs/testing-generators.md

Run:
- cd cli && go test ./...
```

---
