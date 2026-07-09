## Prompt 40 — App manifest schema validation

```text
Add strong validation for ClusterForge app manifests.

Target:
- CLI app manifest support.

Goal:
Validate apps/*.yaml before rendering Terraform.

Validation rules:
- name required and DNS/Kubernetes-name compatible.
- type must be one of:
  - web
  - worker
  - cronjob
  - service
- image required.
- replicas >= 0.
- ports container_port must be 1-65535.
- protocol must be TCP or UDP.
- ingress.host required when ingress.enabled=true.
- ingress.path must start with /.
- autoscaling min_replicas <= max_replicas.
- autoscaling max_replicas > 0.
- resources values must be non-empty strings if provided.
- secret_env must reference secret_name and secret_key only; no secret values.

Add command:
- cf app validate <name>
- cf app validate --all

Behavior:
- Print clear validation errors with field paths.
- Example:
  apps/api.yaml: ingress.host is required when ingress.enabled=true

Tests:
- Valid manifest passes.
- Missing image fails.
- Bad port fails.
- Ingress enabled without host fails.
- Secret value accidentally provided fails.

Docs:
- Update docs/cli.md.
- Add docs/app-manifest.md with full schema and examples.

Run:
- gofmt
- go test ./...
```

---
