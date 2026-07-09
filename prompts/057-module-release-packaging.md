## Prompt 57 — Module release packaging

```text
Prepare Terraform modules for release and distribution.

Goal:
Make modules usable from Git source and optionally from a Terraform registry later.

Tasks:
1. Review module naming.
2. Ensure each public module has:
   - README.md
   - examples
   - versions.tf
   - inputs/outputs documentation
   - clear status
   - semantic version compatibility notes

3. Create:
   - docs/module-release.md
   - MODULE_CATALOG.md

MODULE_CATALOG.md must include:
- Module path
- Purpose
- Stability:
  - stable
  - beta
  - experimental
  - placeholder
- Providers
- Example path
- Test status
- Registry-ready:
  yes/no

4. Add module source examples:
   - local path
   - Git source with ref tag
   - future registry style

5. Add release tagging rules:
   - v0.1.0 for whole repo
   - module-specific compatibility notes
   - no breaking changes without minor/major version bump depending on pre-1.0 policy

6. Add CI check:
   - all stable modules must have README and example
   - placeholder modules must be clearly marked

Rules:
- Do not publish modules automatically.
- Do not rename directories without updating references.
- Do not mark modules stable unless tests/docs/examples exist.

Final response:
- Summarize release readiness by module.
```

---
