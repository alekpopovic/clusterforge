output "job_id" {
  description = "Nomad job resource ID."
  value       = module.service.job_id
}

output "job_name" {
  description = "Nomad job name."
  value       = module.service.job_name
}
