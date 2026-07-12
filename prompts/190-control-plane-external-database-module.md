# Prompt 190 — Control Plane external database module

```text
Add production database deployment guidance and modules for Control Plane.

For AWS:
- Use existing modules/cloud/aws/rds-postgres if implemented.
- Create example:
  examples/control-plane-aws-rds

For Kubernetes/local:
- Provide docs for external PostgreSQL.
- Optional dev-only PostgreSQL Helm example:
  examples/control-plane-local-postgres

Docs:
- docs/control-plane-database.md

Cover:
- PostgreSQL required for production
- backups
- encryption
- network access
- migrations
- connection secrets
- high availability
- restore process

Rules:
- Do not recommend embedded database for production.
- Do not store DB password in Terraform outputs.
- Use external secrets or Kubernetes Secret references.
```
