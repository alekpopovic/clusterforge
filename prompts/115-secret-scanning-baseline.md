## Prompt 115 — Secret scanning baseline

```text
Add secret scanning baseline.

Goal:
Prevent accidental commits of credentials, kubeconfigs, tfstate, and tokens.

Create:
- .gitleaks.toml or equivalent config
- scripts/secret-scan.sh
- docs/secret-scanning.md

Integrate:
- Makefile target: make secret-scan
- CI security workflow
- pre-commit config if available

Scan should catch:
- AWS access keys
- GitHub tokens
- private keys
- kubeconfig client keys/certs
- Terraform state files
- .env files
- generic passwords where confident

Rules:
- Do not commit fake secrets that trigger scanners.
- Allowlist only documented test fixtures.
- Do not broadly suppress findings.
- Secret scan should not require cloud credentials.

Run:
- scripts/secret-scan.sh if tool is installed
- otherwise verify missing-tool behavior

Final response:
- List scanning configuration.
- Mention skipped/allowlisted patterns.
```

---
