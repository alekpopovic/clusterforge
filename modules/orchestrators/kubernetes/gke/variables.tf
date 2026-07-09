variable "project_id" {
  description = "GCP project ID."
  type        = string
}

variable "name" {
  description = "GKE cluster name."
  type        = string
}

variable "environment" {
  description = "Environment identifier."
  type        = string
}

variable "region" {
  description = "GCP region or location."
  type        = string
}

variable "network" {
  description = "VPC network name or self link."
  type        = string
}

variable "subnetwork" {
  description = "Subnetwork name or self link."
  type        = string
}

variable "pods_secondary_range_name" {
  description = "Secondary range name for pods."
  type        = string
  default     = "pods"
}

variable "services_secondary_range_name" {
  description = "Secondary range name for services."
  type        = string
  default     = "services"
}

variable "kubernetes_version" {
  description = "Optional GKE Kubernetes version."
  type        = string
  default     = ""
}

variable "remove_default_node_pool" {
  description = "Whether to remove the default node pool."
  type        = bool
  default     = true
}

variable "initial_node_count" {
  description = "Initial node count for the temporary default pool."
  type        = number
  default     = 1
}

variable "node_pools" {
  description = "GKE node pools to create."
  type = map(object({
    machine_type = optional(string, "e2-standard-2")
    node_count   = optional(number, 2)
  }))
  default = {
    default = {}
  }
}
