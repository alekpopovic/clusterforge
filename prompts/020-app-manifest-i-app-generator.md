## Prompt 20 — App manifest i app generator

```text
Implement app manifest support in ClusterForge CLI.

Goal:
Allow users to define an application once in apps/<name>.yaml and generate Terraform module calls for Kubernetes or ECS.

Add commands:
- cf app add <name>
- cf app list
- cf app render <name> --env <env>
- cf app remove <name>

App manifest schema:
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

cf app add behavior:
- Creates apps/<name>.yaml.
- Accept flags:
  --image
  --port
  --replicas
  --host
  --type
- Do not overwrite existing app without --force.

cf app render behavior:
- Reads clusterforge.yaml.
- Reads apps/<name>.yaml.
- Determines target from environment orchestrator.
- For eks/kubernetes/k3s/rke2/gke/aks:
  - Generate a Terraform module call using modules/workloads/kubernetes/app.
- For ecs:
  - Generate a Terraform module call using modules/workloads/ecs/service.
- Put rendered file in:
  live/<env>/apps/<name>.tf
  or env.path/apps/<name>.tf
- Keep generated HCL readable.

Important:
- Do not put secret values in generated Terraform.
- Use secret references only.
- Add comments where user must provide existing secret names or IAM permissions.

Tests:
- Add app manifest.
- List app manifests.
- Render Kubernetes module call.
- Render ECS module call.
- Refuse unsupported orchestrator with helpful error.

Run:
- gofmt
- go test ./...
```

---
