# Kubernetes Observability

Example root for installing basic Kubernetes observability components into an
existing cluster:

- kube-prometheus-stack for metrics, alerting, and Grafana
- Loki for logs
- Grafana Alloy as the log/telemetry agent

The Kubernetes and Helm providers use `var.kubeconfig_path`, defaulting to
`~/.kube/config`.

## Usage

```bash
terraform init
terraform validate
terraform plan
```

Grafana ingress is disabled by default. Enable it only after configuring DNS,
TLS, authentication, and network policy:

```bash
terraform plan -var='grafana_host=grafana.example.com'
```

Persistent storage is disabled by default:

```bash
terraform plan \
  -var='enable_persistent_storage=true' \
  -var='storage_class_name=gp3'
```

Alloy needs explicit production values for log sources and destinations. Do not
put credentials in Terraform values; reference Kubernetes Secrets or an external
secret manager instead.
