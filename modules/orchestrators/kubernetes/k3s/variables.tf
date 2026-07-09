variable "cluster_name" {
  description = "Logical K3s cluster name."
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

variable "k3s_version" {
  description = "Optional K3s version, such as v1.30.0+k3s1."
  type        = string
  default     = ""
}

variable "install_channel" {
  description = "K3s install channel."
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
  description = "K3s bundled components to disable."
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
