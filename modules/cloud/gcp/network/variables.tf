variable "project_id" {
  description = "GCP project ID."
  type        = string
}

variable "name" {
  description = "Base name for GCP network resources."
  type        = string
}

variable "environment" {
  description = "Environment identifier."
  type        = string
}

variable "region" {
  description = "GCP region."
  type        = string
}

variable "auto_create_subnetworks" {
  description = "Whether the VPC should use auto-created subnetworks."
  type        = bool
  default     = false
}

variable "subnet_cidr" {
  description = "Primary subnet CIDR range."
  type        = string
  default     = "10.50.0.0/20"
}

variable "secondary_pod_range_cidr" {
  description = "Secondary range CIDR for GKE pods."
  type        = string
  default     = "10.52.0.0/16"
}

variable "secondary_service_range_cidr" {
  description = "Secondary range CIDR for GKE services."
  type        = string
  default     = "10.53.0.0/20"
}

variable "labels" {
  description = "Labels to apply to supported resources."
  type        = map(string)
  default     = {}
}
