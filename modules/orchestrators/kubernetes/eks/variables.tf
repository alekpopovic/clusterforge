variable "name" {
  description = "EKS cluster name."
  type        = string

  validation {
    condition     = length(trimspace(var.name)) > 0
    error_message = "Name must not be empty."
  }
}

variable "environment" {
  description = "Environment name for tagging and resource names."
  type        = string

  validation {
    condition     = length(trimspace(var.environment)) > 0
    error_message = "Environment must not be empty."
  }
}

variable "kubernetes_version" {
  description = "Kubernetes version for the EKS control plane."
  type        = string
  default     = "1.30"
}

variable "vpc_id" {
  description = "VPC ID where EKS will run. Kept as an explicit input for validation and future security group integration."
  type        = string

  validation {
    condition     = length(trimspace(var.vpc_id)) > 0
    error_message = "VPC ID must not be empty."
  }
}

variable "subnet_ids" {
  description = "Subnet IDs used by the EKS control plane and as defaults for node groups."
  type        = list(string)

  validation {
    condition     = length(var.subnet_ids) > 0 && alltrue([for subnet_id in var.subnet_ids : length(trimspace(subnet_id)) > 0])
    error_message = "At least one non-empty subnet ID is required."
  }
}

variable "endpoint_public_access" {
  description = "Whether the EKS API endpoint is publicly reachable."
  type        = bool
  default     = true
}

variable "endpoint_private_access" {
  description = "Whether the EKS API endpoint is reachable from inside the VPC."
  type        = bool
  default     = true
}

variable "public_access_cidrs" {
  description = "CIDR blocks allowed to reach the public EKS API endpoint."
  type        = list(string)
  default     = ["0.0.0.0/0"]

  validation {
    condition     = alltrue([for cidr in var.public_access_cidrs : can(cidrnetmask(cidr))])
    error_message = "Each public access CIDR must be a valid IPv4 CIDR block."
  }
}

variable "enabled_cluster_log_types" {
  description = "EKS control plane log types to enable."
  type        = list(string)
  default     = ["api", "audit", "authenticator"]

  validation {
    condition = alltrue([
      for log_type in var.enabled_cluster_log_types :
      contains(["api", "audit", "authenticator", "controllerManager", "scheduler"], log_type)
    ])
    error_message = "Enabled cluster log types must be one of api, audit, authenticator, controllerManager, or scheduler."
  }
}

variable "tags" {
  description = "Tags applied to supported AWS resources."
  type        = map(string)
  default     = {}
}

variable "node_groups" {
  description = "Managed node groups to create for the EKS cluster."
  type = map(object({
    subnet_ids     = optional(list(string))
    instance_types = optional(list(string), ["t3.medium"])
    capacity_type  = optional(string, "ON_DEMAND")
    min_size       = optional(number, 1)
    desired_size   = optional(number, 2)
    max_size       = optional(number, 4)
    disk_size      = optional(number, 50)
    labels         = optional(map(string), {})
    taints = optional(list(object({
      key    = string
      value  = optional(string)
      effect = string
    })), [])
  }))
  default = {
    default = {}
  }

  validation {
    condition = alltrue([
      for node_group in values(var.node_groups) :
      node_group.min_size <= node_group.desired_size && node_group.desired_size <= node_group.max_size
    ])
    error_message = "Each node group must satisfy min_size <= desired_size <= max_size."
  }

  validation {
    condition = alltrue([
      for node_group in values(var.node_groups) :
      contains(["ON_DEMAND", "SPOT"], node_group.capacity_type)
    ])
    error_message = "Each node group capacity_type must be ON_DEMAND or SPOT."
  }

  validation {
    condition = alltrue(flatten([
      for node_group in values(var.node_groups) : [
        for taint in node_group.taints :
        contains(["NO_SCHEDULE", "NO_EXECUTE", "PREFER_NO_SCHEDULE"], taint.effect)
      ]
    ]))
    error_message = "Node group taint effects must be NO_SCHEDULE, NO_EXECUTE, or PREFER_NO_SCHEDULE."
  }
}

variable "enable_vpc_cni_addon" {
  description = "Whether to install or manage the Amazon VPC CNI EKS add-on."
  type        = bool
  default     = true
}

variable "enable_coredns_addon" {
  description = "Whether to install or manage the CoreDNS EKS add-on."
  type        = bool
  default     = true
}

variable "enable_kube_proxy_addon" {
  description = "Whether to install or manage the kube-proxy EKS add-on."
  type        = bool
  default     = true
}

variable "enable_ebs_csi_driver_addon" {
  description = "Whether to install or manage the AWS EBS CSI Driver EKS add-on. IRSA is a future enhancement."
  type        = bool
  default     = false
}
