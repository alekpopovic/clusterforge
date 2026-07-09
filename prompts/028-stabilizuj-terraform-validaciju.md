## Prompt 28 — Stabilizuj Terraform validaciju

```text
Improve Terraform validation across the repository.

Goal:
Make Terraform formatting and validation reliable for modules, examples, and live templates where possible.

Tasks:
1. Inspect all Terraform files.
2. Fix syntax issues.
3. Ensure every module has:
   - versions.tf
   - variables.tf
   - outputs.tf
   - main.tf
   - README.md

4. Update scripts/validate.sh:
   - Run terraform fmt -check -recursive.
   - Find Terraform root directories.
   - Run terraform init -backend=false where safe.
   - Run terraform validate where safe.
   - Skip directories that require real cloud credentials or remote backend initialization.
   - Print clear skip reasons.

5. Add scripts/list-terraform-roots.sh:
   - Detect directories that contain Terraform root configurations.
   - Exclude .terraform directories.
   - Exclude generated temp directories.

6. Add docs/validation.md:
   - Explain what gets validated.
   - Explain why some examples are skipped.
   - Explain how to validate with real cloud credentials.

Rules:
- Do not add fake providers just to make validation pass.
- Do not hardcode credentials.
- Do not commit .terraform directories.
- Do not modify module behavior unless needed to fix syntax or validation.
- Prefer explicit skip reasons over silent failures.

Run:
- terraform fmt -recursive
- scripts/validate.sh

Final response:
- Summarize validation improvements.
- List directories validated.
- List directories skipped and why.
- Mention remaining validation limitations.
```

---
