output "prometheus_stack_release" {
  description = "kube-prometheus-stack Helm release name."
  value       = module.prometheus_stack.release_name
}

output "loki_release" {
  description = "Loki Helm release name."
  value       = module.loki.release_name
}

output "alloy_release" {
  description = "Grafana Alloy Helm release name."
  value       = module.alloy.release_name
}
