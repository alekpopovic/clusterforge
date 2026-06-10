variable "app" {
  description = "Application label value."
  type        = string
}

variable "component" {
  description = "Component label value."
  type        = string
  default     = "app"
}

variable "part_of" {
  description = "Higher-level system this application belongs to."
  type        = string
  default     = "clusterforge"
}

variable "managed_by" {
  description = "Controller or tool managing the resource."
  type        = string
  default     = "terraform"
}

variable "extra_labels" {
  description = "Additional labels to merge with the standard Kubernetes labels."
  type        = map(string)
  default     = {}
}
