## Prompt 9 — Kubernetes platform bootstrap

```text
Implement modules/platform/kubernetes/bootstrap.

Purpose:
A convenience composition module that enables common Kubernetes platform add-ons.

Important:
This module should not create the cluster. It assumes Kubernetes and Helm providers are already configured in the root environment.

Inputs:
- namespace_prefix: string, default ""
- enable_ingress_nginx: bool, default true
- enable_cert_manager: bool, default true
- enable_external_dns: bool, default false
- enable_metrics_server: bool, default true
- enable_prometheus_stack: bool, default false
- enable_loki: bool, default false
- enable_argocd: bool, default false
- common_labels: map(string), default {}

Behavior:
- Call child modules conditionally using count or for_each:
  - platform/kubernetes/ingress-nginx
  - cert-manager
  - external-dns
  - metrics-server
  - prometheus-stack
  - loki
  - argocd

Implement child modules as Helm release wrappers.

For each child module:
- Use helm_release.
- Create namespace if requested.
- Inputs:
  - namespace
  - chart_version
  - values
  - labels
  - create_namespace
- Outputs:
  - namespace
  - release_name
  - release_status if available

Default charts:
- ingress-nginx:
  repository: https://kubernetes.github.io/ingress-nginx
  chart: ingress-nginx
- cert-manager:
  repository: https://charts.jetstack.io
  chart: cert-manager
- external-dns:
  repository: https://kubernetes-sigs.github.io/external-dns/
  chart: external-dns
- metrics-server:
  repository: https://kubernetes-sigs.github.io/metrics-server/
  chart: metrics-server
- prometheus-stack:
  repository: https://prometheus-community.github.io/helm-charts
  chart: kube-prometheus-stack
- loki:
  repository: https://grafana.github.io/helm-charts
  chart: loki
- argocd:
  repository: https://argoproj.github.io/argo-helm
  chart: argo-cd

Do not invent complex values. Provide sane empty/default values.

README:
- Explain Terraform vs GitOps boundary.
- Explain that this module installs platform add-ons, not business applications.
- Include example with ingress-nginx, cert-manager and argocd enabled.
- Include warning about CRDs and Helm lifecycle.

Run terraform fmt -recursive.
```

---
