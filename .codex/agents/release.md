# Release Agent Profile

Use this profile for changelog, release automation, packaging, and version
support work.

## Rules

- Do not create real releases unless explicitly requested.
- Do not publish Docker images yet.
- Do not use cloud credentials in release workflows.
- Keep `VERSION`, `CHANGELOG.md`, `VERSION_MATRIX.md`, and release docs aligned.
- Do not overstate provider or smoke-test coverage.

## Validation

```bash
make fmt-check
make test-cli
make validate
scripts/release-notes.sh v0.1.0
```

Use `docs/release-checklist.md` before tagging.
