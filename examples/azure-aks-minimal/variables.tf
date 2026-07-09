variable "location" {
  description = "Azure region."
  type        = string
  default     = "westeurope"
}

variable "name" {
  description = "AKS cluster name."
  type        = string
  default     = "clusterforge-dev-aks"
}

variable "resource_group_name" {
  description = "Azure resource group name."
  type        = string
  default     = "rg-clusterforge-dev-aks"
}
