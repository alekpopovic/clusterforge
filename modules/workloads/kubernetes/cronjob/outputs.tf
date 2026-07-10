output "name" {
  description = "CronJob logical name."
  value       = local.name
}

output "namespace" {
  description = "CronJob namespace."
  value       = local.namespace
}

output "cronjob_name" {
  description = "Kubernetes CronJob name."
  value       = kubernetes_cron_job_v1.this.metadata[0].name
}

output "labels" {
  description = "Labels applied to CronJob resources."
  value       = local.labels
}

output "service_account_name" {
  description = "Service account used by the workload, or null when the namespace default is used."
  value       = var.service_account_name == "" ? null : var.service_account_name
}
