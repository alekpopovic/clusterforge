# Code Change Workflow

1. Read `AGENTS.md` and the relevant `.codex/agents/*.md` profile.
2. Check `git status --short --branch`.
3. Inspect nearby files and existing patterns before editing.
4. Make the smallest coherent change.
5. Run targeted validation.
6. Run broader validation when shared behavior changes.
7. Stage, commit, and push from `main`.

Suggested validation:

```bash
make fmt-check
make lint
make test-cli
make validate
```
