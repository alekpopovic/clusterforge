## Prompt 27 — Makefile / developer workflow

```text
Add a developer task runner for ClusterForge.

Goal:
Make local development easier with standard commands.

Create:
- Makefile

Optional, only if useful:
- Taskfile.yml

Makefile targets:
- help
- fmt
- fmt-check
- validate
- lint
- test
- test-cli
- build-cli
- security
- docs
- clean
- ci

Behavior:
- make help prints available targets.
- make fmt runs terraform fmt -recursive and gofmt on cli if present.
- make fmt-check checks formatting without modifying files.
- make validate runs scripts/validate.sh.
- make test runs CLI tests and any lightweight Terraform tests.
- make build-cli builds cli/cf.
- make security runs available security scans.
- make clean removes local build artifacts, .cf plans, and temporary files but never removes source files.

Update:
- README.md with local development commands.
- docs/cli.md if needed.
- scripts/lint.sh and scripts/validate.sh if they need to align with Makefile.

Rules:
- Use bash strict mode in scripts.
- Do not require cloud credentials for default make ci.
- Do not run terraform apply.
- Do not run destructive commands.
- Keep commands portable and readable.

Run:
- make help
- make fmt-check if possible
- make test if possible

Final response:
- List created/updated files.
- List commands run.
- Report pass/fail status.
```

---
