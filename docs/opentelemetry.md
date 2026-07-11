# OpenTelemetry Collector

The optional module installs the upstream collector Helm chart for trace,
metric and log pipelines. It is disabled by default in platform bootstrap.

Deployment mode suits a central gateway, daemonset suits node-local log/host
collection, and statefulset is reserved for pipelines that require stable
identity/storage semantics. Choose receivers, processors, exporters, queues,
memory limits and replicas based on measured volume; telemetry can consume
substantial CPU, memory, network and backend quota.

Vendor endpoints and API keys belong in external secret systems and workload
identity, never literal Terraform values. Minimal built-in values only set mode,
optional upstream presets and service-account annotations. Production users
must pin the chart, restrict egress, test backpressure/data loss behavior,
configure sampling and retention, monitor collector health, and validate
upgrade/rollback behavior for traces, metrics and logs independently.
