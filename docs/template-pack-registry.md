# Template pack registry

Organizations can keep versioned Terraform template packs outside the core
repository and fetch them explicitly into a local cache. This is a source
registry, not a marketplace: ClusterForge never searches for, installs, or
executes remote code.

## Configuration

```yaml
template_packs:
  - name: company-standard
    source: git::https://github.com/example/company-clusterforge-templates.git?ref=v0.1.0
    version: v0.1.0
    enabled: true
```

Supported sources are a local directory (`path::templates/company` or a plain
path), a local `.zip`, `.tar.gz`, or `.tgz` archive (`archive::...`), and a Git
repository with an explicit `?ref=`. Fetched packs are stored under
`.cf/cache/template-packs/<name>/<version>` and the cache is ignored by Git.

Existing configurations using `path:` remain supported and are read directly;
they do not need a fetch step.

## Workflow

```bash
cf template list
cf template fetch company-standard
cf template validate company-standard
cf template update company-standard
cf template cache clear
```

`fetch` prints the configured source before reading it and refuses to replace
an existing cached version. `update` explicitly replaces that version. Neither
command runs hooks, scripts, binaries, or Terraform from the pack.

Each pack must contain `metadata.yaml`, at least one file below both `env/` and
`app/`, and declarations for supported clouds and orchestrators. Validation
rejects executable files, symlinks in fetched local packs, oversized files, and
common literal-secret patterns. These checks reduce risk but do not replace
human review or secret scanning.

## Pinning and security

Use an immutable commit SHA or a protected release tag for Git sources. Branch
refs such as `main` and `master` are mutable and produce a warning. The
configuration `version` selects the cache location; it does not prove that a
mutable source still has the same contents.

Review the source owner, exact ref, metadata, templates, and generated
Terraform before use. Do not place credentials in source URLs or archives.
Archive extraction rejects entries that escape the cache directory, and pack
files are copied without executable permissions.
