variable "location" {
  description = "Azure region."
  type        = string
  default     = "westeurope"
}

variable "environment" {
  description = "Environment name."
  type        = string
  default     = "dev"
}

variable "name" {
  description = "AKS cluster name."
  type        = string
  default     = "clusterforge-dev-aks"
}

variable "resource_group_name" {
  description = "Resource group name."
  type        = string
}

variable "address_space" {
  description = "VNet address space."
  type        = list(string)
  default     = ["10.40.0.0/16"]
}

variable "subnet_prefixes" {
  description = "Subnet prefixes."
  type        = list(string)
  default     = ["10.40.1.0/24"]
}

variable "kubernetes_version" {
  description = "AKS Kubernetes version."
  type        = string
  default     = ""
}

variable "dns_prefix" {
  description = "AKS DNS prefix."
  type        = string
  default     = "clusterforge-dev"
}

variable "default_node_pool" {
  description = "Default node pool."
  type        = any
  default     = {}
}

variable "tags" {
  description = "Azure tags."
  type        = map(string)
  default     = {}
}
