# Template Packs

Template packs let teams override generated Terraform templates without
forking the CLI.

```yaml
template_packs:
  - name: company-standard
    path: templates/company-standard
```

Expected layout:

```text
templates/company-standard/
  metadata.yaml
  env/
    aws-eks/
    aws-ecs/
  app/
    kubernetes/
    ecs/
```

Validate and use:

```bash
cf template list
cf template validate
cf generate dev --template-pack company-standard
cf app render web --env dev --template-pack company-standard
```

Template packs use `text/template` files only. They do not execute arbitrary
code.
