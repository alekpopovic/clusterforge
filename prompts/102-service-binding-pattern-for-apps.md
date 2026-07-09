## Prompt 102 — Service binding pattern for apps

```text
Add service binding support to app manifests.

Goal:
Allow applications to reference infrastructure dependencies without hardcoding outputs manually.

App manifest extension:
dependencies:
  database:
    type: rds-postgres
    reference: main
    env:
      DATABASE_HOST: endpoint
      DATABASE_PORT: port
      DATABASE_NAME: db_name
  cache:
    type: elasticache-redis
    reference: redis
    env:
      REDIS_HOST: primary_endpoint_address
  queue:
    type: sqs
    reference: jobs
    env:
      QUEUE_URL: queue_url

CLI behavior:
- cf app render should read environment dependency registry.
- Generate Terraform references when dependency modules are in same stack.
- For cross-stack dependencies, generate terraform_remote_state references or documented placeholders.
- Never inject secret values directly.
- For secrets, generate external secret references or comments.

Create:
- docs/service-bindings.md
- cli/internal/bindings/

Tests:
- App with database dependency renders env reference.
- App with queue dependency renders correct reference.
- Unknown dependency fails clearly.
- Secret dependency does not render plaintext.

Rules:
- Keep MVP simple.
- Do not build full dependency graph engine yet.
- Make generated Terraform readable.

Run:
- gofmt
- go test ./...
```

---
