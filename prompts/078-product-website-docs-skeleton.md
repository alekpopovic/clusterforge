## Prompt 78 — Product website docs skeleton

```text
Create a documentation website skeleton for ClusterForge.

Goal:
Prepare docs for publishing with a static site generator.

Choose one:
- Docusaurus
- MkDocs
- Hugo

Prefer MkDocs Material if simple and docs are Markdown-heavy.

Create:
- mkdocs.yml
- docs-site/ or reuse docs/ if appropriate
- docs/index.md
- docs/getting-started.md
- docs/architecture.md
- docs/cli.md
- docs/modules.md
- docs/security.md
- docs/roadmap.md

Add:
- scripts/docs-serve.sh
- scripts/docs-build.sh
- Makefile targets:
  - docs-serve
  - docs-build

Rules:
- Do not remove existing docs.
- Avoid heavy customization.
- Keep docs source of truth in Markdown.
- Do not require publishing credentials.

Final response:
- Explain how to run local docs site.
```

---
