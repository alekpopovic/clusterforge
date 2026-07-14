# Prompt 274 — Blueprint registry support

```text
Add blueprint registry support.

Goal:
Allow organizations to maintain reusable environment and cluster blueprints outside the core repository.

Sources:
- local path
- Git source with ref
- archive file

Config:
blueprint_registries:
  - name: company-blueprints
    source: git::https://github.com/example/clusterforge-blueprints.git?ref=v0.1.0
    enabled: true

CLI:
- cf blueprint registry list
- cf blueprint registry fetch <name>
- cf blueprint registry update <name>
- cf blueprint registry validate <name>
- cf blueprint cache clear

Cache:
- .cf/cache/blueprints/<registry>/<version>

Validation:
- metadata.yaml exists
- blueprints are valid
- supported cloud/orchestrator declared
- no secrets in blueprints
- refs pinned
- warn on main/master branch refs

Security:
- no code execution
- templates only
- explicit fetch
- source shown to user

Tests:
- local registry
- Git source parsing
- cache path
- missing metadata fails
- unpinned ref warns
- blueprint validation

Docs:
- docs/blueprint-registry.md

Rules:
- Do not implement marketplace.
- Do not download remote code automatically.
```
