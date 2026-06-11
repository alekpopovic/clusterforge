output "job_id" {
  description = "Nomad job resource ID."
  value       = nomad_job.this.id
}

output "job_name" {
  description = "Nomad job name."
  value       = local.name
}
