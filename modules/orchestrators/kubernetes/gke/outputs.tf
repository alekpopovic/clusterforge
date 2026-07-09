output "cluster_name" {
  description = "GKE cluster name."
  value       = google_container_cluster.this.name
}

output "endpoint" {
  description = "GKE Kubernetes endpoint."
  value       = google_container_cluster.this.endpoint
  sensitive   = true
}

output "ca_certificate" {
  description = "GKE cluster CA certificate."
  value       = google_container_cluster.this.master_auth[0].cluster_ca_certificate
  sensitive   = true
}

output "location" {
  description = "GKE cluster location."
  value       = google_container_cluster.this.location
}
