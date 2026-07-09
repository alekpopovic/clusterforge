## Prompt 77 — Template pack support

```text
Implement local template pack support for ClusterForge.

Goal:
Allow organizations to customize generated Terraform and app manifests without forking the CLI.

Config:
clusterforge.yaml:
template_packs:
  - name: company-standard
    path: templates/company-standard

CLI:
- cf template list
- cf template validate
- cf generate <env> --template-pack company-standard
- cf app render <name> --template-pack company-standard

Template pack structure:
templates/company-standard/
  env/
    aws-eks/
    aws-ecs/
  app/
    kubernetes/
    ecs/
  metadata.yaml

metadata.yaml:
- name
- version
- description
- supported_clouds
- supported_orchestrators

Behavior:
- Built-in templates remain default.
- Local template pack overrides built-in templates.
- Validate required template files.
- Do not execute arbitrary code.
- Use text/template only.

Tests:
- Built-in templates still work.
- Local template pack overrides env template.
- Missing template produces useful error.
- Invalid metadata fails validation.

Docs:
- docs/template-packs.md

Run:
- gofmt
- go test ./...
```

---
