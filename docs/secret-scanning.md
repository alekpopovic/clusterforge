# Secret scanning baseline

ClusterForge combines Gitleaks content detection with explicit checks for file
types that must never be tracked.

```bash
make secret-scan
./scripts/secret-scan.sh
./scripts/secret-scan.sh --staged
```

`.gitleaks.toml` extends Gitleaks' built-in rules, which cover common providers
including AWS and GitHub tokens and private keys. ClusterForge adds focused rules
for embedded kubeconfig client key/certificate data and high-entropy password
assignments.

The wrapper independently rejects tracked or staged Terraform state, `.terraform`
content, kubeconfig files, and `.env` variants. `.env.example` is the sole filename
exception because it is an intentionally non-secret template; its contents are
still scanned by Gitleaks when committed. There are no broad path, commit, rule,
or test-fixture allowlists in the Gitleaks configuration.

If Gitleaks is unavailable locally, prohibited-file checks still run and the
script reports that content scanning was skipped. Set `REQUIRE_GITLEAKS=true` to
make a missing binary fatal. CI downloads a pinned Gitleaks release, verifies its
published archive checksum, checks out full history, and requires the content
scan to run.

The pre-commit hook calls staged mode. CI scans repository history because local
hooks can be bypassed and secrets must also be removed from Git history, not only
from the latest tree.

## Responding to a finding

1. Do not paste the detected value into an issue, log, or pull request comment.
2. Revoke or rotate the credential immediately.
3. Remove it from the working tree and Git history using the organization's
   approved history-rewrite process.
4. Review access logs and affected infrastructure.
5. Add a narrow allowlist only after proving the value is a non-secret fixture and
   documenting its exact path and purpose.

Gitleaks redacts findings in normal output, but scanner reports and CI logs should
still be treated as sensitive. Configuration follows the official
[Gitleaks configuration and CLI documentation](https://github.com/gitleaks/gitleaks).
