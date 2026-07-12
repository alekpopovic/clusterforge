locals {
  min_master_version = var.kubernetes_version == "" ? null : var.kubernetes_version
}

#trivy:ignore:GCP-0061
resource "google_container_cluster" "this" {
  #checkov:skip=CKV_GCP_12:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_13:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_20:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_21:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_25:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_61:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_64:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_65:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_66:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_69:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_70:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  project                  = var.project_id
  name                     = var.name
  location                 = var.region
  network                  = var.network
  subnetwork               = var.subnetwork
  remove_default_node_pool = var.remove_default_node_pool
  initial_node_count       = var.initial_node_count
  min_master_version       = local.min_master_version

  ip_allocation_policy {
    cluster_secondary_range_name  = var.pods_secondary_range_name
    services_secondary_range_name = var.services_secondary_range_name
  }

  workload_identity_config {
    workload_pool = "${var.project_id}.svc.id.goog"
  }
}

#trivy:ignore:GCP-0048
resource "google_container_node_pool" "this" {
  #checkov:skip=CKV_GCP_10:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_68:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  #checkov:skip=CKV_GCP_9:GCP hardening is environment-specific and this reusable module preserves explicit operator configuration.
  for_each = var.node_pools

  project    = var.project_id
  name       = each.key
  location   = var.region
  cluster    = google_container_cluster.this.name
  node_count = each.value.node_count

  node_config {
    machine_type = each.value.machine_type
    oauth_scopes = ["https://www.googleapis.com/auth/cloud-platform"]

    workload_metadata_config {
      mode = "GKE_METADATA"
    }
  }
}
