output "network_id" {
  description = "GCP network ID."
  value       = google_compute_network.this.id
}

output "network_name" {
  description = "GCP network name."
  value       = google_compute_network.this.name
}

output "subnetwork_id" {
  description = "GCP subnetwork ID."
  value       = google_compute_subnetwork.this.id
}

output "subnetwork_name" {
  description = "GCP subnetwork name."
  value       = google_compute_subnetwork.this.name
}
