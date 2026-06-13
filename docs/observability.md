---
title: Observability
permalink: /observability/
---

# Observability

ClusterForge separates metrics, logs, and agents into focused Kubernetes
platform modules.

## Metrics

Use `modules/platform/kubernetes/prometheus-stack` for metrics, alerting, and
Grafana dashboards through kube-prometheus-stack.

Grafana ingress is disabled by default. Enable it only after reviewing:

- DNS and TLS
- authentication
- allowed source networks
- ingress controller behavior
- whether Grafana should be private-only

Do not place Grafana admin passwords or OAuth secrets in Terraform values.
Reference existing Kubernetes Secrets or use External Secrets Operator.

## Logs

Use `modules/platform/kubernetes/loki` for log storage and querying.

Loki is installed without persistent storage by default. That is safer for
development and avoids surprising storage costs, but it is not a production
retention strategy.

For production, decide on:

- retention period
- persistent volumes or object storage
- resource requests and limits
- tenant and label strategy
- backup and restore behavior

## Log Agent

ClusterForge uses `modules/platform/kubernetes/alloy` as the simple log and
telemetry agent wrapper. Grafana Alloy is the preferred direction for new
installations instead of starting new work on Promtail.

The Alloy module installs the chart but does not invent collection rules.
Provide explicit values for your cluster, labels, processors, and Loki endpoint.
Keep credentials out of Terraform state by referencing Kubernetes Secrets or an
external secret store.

## Storage

Persistent storage is disabled by default for Prometheus and Loki. Enable it
only after choosing a StorageClass and sizing strategy:

```hcl
storage_enabled    = true
storage_class_name = "gp3"
```

The module defaults are starting points. Production values should be tuned for
cluster size, retention, scrape volume, log volume, and upgrade process.

## Example

See `examples/kubernetes-observability` for a root module that installs
kube-prometheus-stack, Loki, and Alloy against an existing Kubernetes cluster.
