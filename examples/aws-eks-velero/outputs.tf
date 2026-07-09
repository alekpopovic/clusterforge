output "velero_bucket_name" {
  description = "S3 bucket name for Velero backups."
  value       = module.velero_bucket.bucket_name
}

output "velero_namespace" {
  description = "Namespace where Velero is installed."
  value       = module.velero.namespace
}
