variable "project_id" {
  description = "GCP project ID."
  type        = string
}

variable "region" {
  description = "GCP region."
  type        = string
  default     = "europe-west1"
}

variable "name" {
  description = "GKE cluster name."
  type        = string
  default     = "clusterforge-dev-gke"
}
