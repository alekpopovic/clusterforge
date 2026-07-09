## Prompt 116 — Devcontainer and reproducible dev environment

```text
Add reproducible development environment support.

Create:
- .devcontainer/devcontainer.json
- .devcontainer/Dockerfile
- docs/development-environment.md

Devcontainer should include:
- Go
- Terraform
- OpenTofu optional if practical
- terraform-docs
- tflint
- checkov or installation instructions
- trivy or installation instructions
- kubectl
- helm
- make
- git

Rules:
- Keep image reasonably small.
- Do not include cloud credentials.
- Do not bake user-specific config.
- Use versions from VERSION_MATRIX.md where practical.
- If a tool is too heavy, document manual install instead.

Update:
- README development setup
- CONTRIBUTING.md

Final response:
- Explain how to open the devcontainer.
- List included tools.
```

---
