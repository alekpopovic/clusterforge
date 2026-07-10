# Pre-commit hooks

ClusterForge's pre-commit configuration runs fast credential-free checks before
each commit:

- trailing whitespace and final newline fixes;
- YAML syntax checks;
- `terraform fmt` for staged Terraform files;
- `gofmt` for staged Go files;
- Gitleaks staged scan when `gitleaks` is installed;
- ShellCheck for staged shell scripts when `shellcheck` is installed.

Install the framework and repository hook, then run all checks once:

```bash
./scripts/install-hooks.sh
pre-commit run --all-files
```

Terraform and Go must be on `PATH` when matching files are committed. Gitleaks
and ShellCheck are optional locally: their wrappers print an explicit skip rather
than downloading tools or requiring cloud credentials. CI remains authoritative
and may install stricter tooling.

The example workflows directory is excluded from the generic YAML hook because
YAML 1.1 parsers can misinterpret GitHub's `on` key; workflow syntax is validated
separately in CI/review.

For a genuine emergency, `git commit --no-verify` bypasses local hooks. Document
why it was necessary in the pull request and run the skipped checks immediately
afterward. Never use `--no-verify` to bypass a secret finding or required CI gate.

The pinned standard hook set comes from the official
[pre-commit-hooks repository](https://github.com/pre-commit/pre-commit-hooks).
