output "job_id" {
  description = "Nomad job resource ID."
  value       = nomad_job.this.id
}
output "jobspec" {
  description = "Rendered batch jobspec."
  value       = local.jobspec
}
