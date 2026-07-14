# Prompt 215 — Plan artifact encryption strategy

```text
Create and implement plan artifact encryption strategy.

Goal:
Protect sensitive plan artifacts when they are stored or transferred.

Docs:
- docs/rfcs/021-plan-artifact-encryption.md
- docs/control-plane/plan-artifact-security.md

Policy:
- raw Terraform plan files disabled by default
- if enabled, encrypt at rest
- restrict access by RBAC
- audit all downloads
- short retention
- optional client-side encryption later

Implementation MVP:
- support encrypted artifact storage for filesystem backend using envelope-style local key from env var
- config:
  artifacts:
    encryption:
      enabled: true
      key_env: CLUSTERFORGE_ARTIFACT_ENCRYPTION_KEY

Requirements:
- do not store encryption key in database
- fail startup if encryption enabled and key missing
- encrypt artifact content before writing
- decrypt only on authorized download
- verify checksum
- audit decrypt/download

Tests:
- encrypted content not stored plaintext
- missing key fails
- download decrypts correctly
- wrong key fails
- audit event created

Rules:
- Do not invent production-grade KMS if not implemented.
- Document local key limitations.
- For S3, document KMS as preferred.
```
