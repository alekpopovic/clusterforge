## Prompt 123 — Template pack registry support

```text
Add template pack registry support.

Goal:
Allow organizations to maintain versioned template packs outside the core repository.

Supported sources:
1. local path
2. Git source with ref
3. archive file path

Do not implement remote marketplace yet.

Config:
clusterforge.yaml:
  template_packs:
    - name: company-standard
      source: git::https://github.com/example/company-clusterforge-templates.git?ref=v0.1.0
      version: v0.1.0
      enabled: true

CLI commands:
- cf template list
- cf template fetch <name>
- cf template update <name>
- cf template validate <name>
- cf template cache clear

Cache:
- Store downloaded template packs under:
  .cf/cache/template-packs/<name>/<version>
- Do not commit cache.
- Add .gitignore entry.

Template pack validation:
- metadata.yaml exists
- declared supported clouds
- declared supported orchestrators
- required template files exist
- no executable files required
- no secrets in templates

Security:
- Print source before fetching.
- Require explicit command to fetch.
- Pin refs strongly in docs.
- Warn when using branch names like main/master.
- Do not execute arbitrary code from template packs.

Tests:
- Local path template pack.
- Git source parsing.
- Cache path generation.
- Missing metadata fails.
- Unpinned source warns.

Docs:
- docs/template-pack-registry.md
- Explain local packs.
- Explain Git packs.
- Explain version pinning.
- Explain security risks.

Run:
- gofmt
- go test ./...
```


---
