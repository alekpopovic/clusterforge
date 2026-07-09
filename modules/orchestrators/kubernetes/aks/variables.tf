variable "name" {
  description = "AKS cluster name."
  type        = string
}

variable "environment" {
  description = "Environment identifier."
  type        = string
}

variable "location" {
  description = "Azure region."
  type        = string
}

variable "resource_group_name" {
  description = "Resource group for the AKS cluster."
  type        = string
}

variable "subnet_id" {
  description = "Subnet ID for the default node pool."
  type        = string
}

variable "kubernetes_version" {
  description = "Optional Kubernetes version."
  type        = string
  default     = ""
}

variable "dns_prefix" {
  description = "AKS DNS prefix."
  type        = string
}

variable "default_node_pool" {
  description = "Default AKS node pool settings."
  type = object({
    name       = optional(string, "system")
    node_count = optional(number, 2)
    vm_size    = optional(string, "Standard_DS2_v2")
  })
  default = {}
}

variable "tags" {
  description = "Tags to apply to AKS resources."
  type        = map(string)
  default     = {}
}
