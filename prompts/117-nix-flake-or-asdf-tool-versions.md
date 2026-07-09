## Prompt 117 — Nix flake or asdf tool versions

```text
Add optional tool version management.

Choose one or both:
1. .tool-versions for asdf
2. flake.nix for Nix

Goal:
Make local development tool versions reproducible.

Tools:
- Go
- Terraform
- OpenTofu
- terraform-docs
- tflint
- checkov
- trivy
- kubectl
- helm
- pre-commit

Create docs:
- docs/tool-versions.md

Rules:
- Do not make Nix/asdf mandatory.
- Make Makefile still work without them.
- Keep versions aligned with VERSION_MATRIX.md.
- Document installation.

Final response:
- List version management files added.
- Explain optional usage.
```

---
