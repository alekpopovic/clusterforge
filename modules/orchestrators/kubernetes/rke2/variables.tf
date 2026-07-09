variable "cluster_name" {
  description = "Logical RKE2 cluster name."
  type        = string
}

variable "environment" {
  description = "Environment identifier."
  type        = string
}

variable "server_count" {
  description = "Number of server nodes expected."
  type        = number
  default     = 1
}

variable "rke2_version" {
  description = "Optional RKE2 version."
  type        = string
  default     = ""
}

variable "install_channel" {
  description = "RKE2 install channel."
  type        = string
  default     = "stable"
}

variable "cluster_cidr" {
  description = "Optional pod CIDR."
  type        = string
  default     = null
}

variable "service_cidr" {
  description = "Optional service CIDR."
  type        = string
  default     = null
}

variable "disable_components" {
  description = "RKE2 packaged components to disable."
  type        = list(string)
  default     = []
}

variable "tls_san" {
  description = "Additional TLS SANs for the Kubernetes API."
  type        = list(string)
  default     = []
}

variable "labels" {
  description = "Metadata labels for generated scripts."
  type        = map(string)
  default     = {}
}
