---
title: App Manifest
permalink: /app-manifest/
---

# App Manifest

ClusterForge app manifests live in `apps/<name>.yaml`. They describe a
workload once, then the CLI can render Terraform module calls for Kubernetes or
ECS environments.

Validate manifests before rendering:

```bash
cf app validate api
cf app validate --all
```

## Schema

```yaml
name: api
type: web
image: ghcr.io/company/api:1.0.0
replicas: 2

ports:
  - name: http
    container_port: 8080
    protocol: TCP

env:
  NODE_ENV: production

secret_env:
  DATABASE_URL:
    secret_name: api-secrets
    secret_key: database-url

resources:
  cpu_request: 100m
  memory_request: 128Mi
  cpu_limit: 500m
  memory_limit: 512Mi

ingress:
  enabled: true
  host: api.dev.example.com
  path: /
  tls: true

autoscaling:
  enabled: true
  min_replicas: 2
  max_replicas: 5
  cpu_percent: 70
```

## Validation Rules

- `name` is required and must be DNS/Kubernetes-name compatible.
- `type` must be one of `web`, `worker`, `cronjob`, or `service`.
- `image` is required.
- `replicas` must be greater than or equal to `0`.
- `ports[].container_port` must be between `1` and `65535`.
- `ports[].protocol` must be `TCP` or `UDP`.
- `ingress.host` is required when `ingress.enabled=true`.
- `ingress.path` must start with `/`.
- `autoscaling.min_replicas` must be less than or equal to
  `autoscaling.max_replicas`.
- `autoscaling.max_replicas` must be greater than `0`.
- Provided `resources` values must be non-empty strings.
- `secret_env` entries may only reference `secret_name` and `secret_key`.

## Secret Handling

Do not put secret values in app manifests.

Use references only:

```yaml
secret_env:
  DATABASE_URL:
    secret_name: api-secrets
    secret_key: database-url
```

Do not use:

```yaml
secret_env:
  DATABASE_URL:
    value: postgres://user:password@example
```

ClusterForge validates against this pattern because manifests are committed to
Git and rendered into Terraform. Secret values should live in Kubernetes
Secrets, External Secrets Operator, AWS Secrets Manager, SSM Parameter Store,
Vault, or another approved secret system.

## Kubernetes Example

```bash
cf app add api --image ghcr.io/company/api:1.0.0 --port 8080 --host api.dev.example.com
cf app validate api
cf app render api --env dev
```

## ECS Example

```yaml
name: worker
type: worker
image: ghcr.io/company/worker:1.0.0
replicas: 2
env:
  QUEUE_NAME: jobs
```

```bash
cf app validate worker
cf app render worker --env dev
```
