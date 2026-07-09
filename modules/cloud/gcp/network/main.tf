locals {
  labels = merge(var.labels, {
    environment = var.environment
    managed_by  = "clusterforge"
  })
}

resource "google_compute_network" "this" {
  project                 = var.project_id
  name                    = "${var.name}-vpc"
  auto_create_subnetworks = var.auto_create_subnetworks
}

resource "google_compute_subnetwork" "this" {
  project       = var.project_id
  name          = "${var.name}-subnet"
  region        = var.region
  network       = google_compute_network.this.id
  ip_cidr_range = var.subnet_cidr

  secondary_ip_range {
    range_name    = "pods"
    ip_cidr_range = var.secondary_pod_range_cidr
  }

  secondary_ip_range {
    range_name    = "services"
    ip_cidr_range = var.secondary_service_range_cidr
  }
}
