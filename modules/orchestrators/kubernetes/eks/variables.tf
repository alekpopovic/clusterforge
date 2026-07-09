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

variable "cluster_log_retention_days" {
  description = "CloudWatch Logs retention in days for EKS control plane logs. Set to 0 to keep logs indefinitely."
  type        = number
  default     = 30

  validation {
    condition     = var.cluster_log_retention_days == 0 || contains([1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, 3653], var.cluster_log_retention_days)
    error_message = "cluster_log_retention_days must be 0 or a valid CloudWatch Logs retention value."
  }
}

variable "enable_cluster_encryption" {
  description = "Whether to enable EKS envelope encryption for Kubernetes secrets."
  type        = bool
  default     = false
}

variable "kms_key_arn" {
  description = "Existing KMS key ARN for EKS secrets encryption. Leave empty to create one when create_kms_key is true."
  type        = string
  default     = ""
}

variable "create_kms_key" {
  description = "Whether to create a KMS key for EKS secrets encryption when enable_cluster_encryption is true and kms_key_arn is empty."
  type        = bool
  default     = false
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

variable "node_group_update_config" {
  description = "Managed node group rolling update settings applied to all node groups."
  type = object({
    max_unavailable            = optional(number)
    max_unavailable_percentage = optional(number)
  })
  default = {
    max_unavailable = 1
  }

  validation {
    condition = (
      try(var.node_group_update_config.max_unavailable, null) != null &&
      try(var.node_group_update_config.max_unavailable_percentage, null) == null
      ) || (
      try(var.node_group_update_config.max_unavailable, null) == null &&
      try(var.node_group_update_config.max_unavailable_percentage, null) != null
    )
    error_message = "Set exactly one of node_group_update_config.max_unavailable or max_unavailable_percentage."
  }
}

variable "node_group_ami_type" {
  description = "AMI type applied to managed node groups. Leave null for AWS defaults."
  type        = string
  default     = null
}

variable "node_group_release_version" {
  description = "AMI release version applied to managed node groups. Leave null for AWS defaults."
  type        = string
  default     = null
}

variable "node_group_force_update_version" {
  description = "Whether to force node group version updates when pods cannot be drained."
  type        = bool
  default     = false
}

variable "node_group_remote_access" {
  description = "Optional SSH remote access settings for managed node groups. Null keeps SSH disabled."
  type = object({
    ec2_ssh_key               = string
    source_security_group_ids = optional(list(string), [])
  })
  default = null
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

variable "enable_irsa" {
  description = "Whether to create an IAM OIDC provider for IAM Roles for Service Accounts."
  type        = bool
  default     = true
}

variable "enable_ebs_csi_driver_addon" {
  description = "Whether to install or manage the AWS EBS CSI Driver EKS add-on."
  type        = bool
  default     = false
}

variable "create_ebs_csi_irsa_role" {
  description = "Whether to create an IRSA role for the EBS CSI controller service account when the EBS CSI add-on is enabled."
  type        = bool
  default     = true

  validation {
    condition     = !var.enable_ebs_csi_driver_addon || !var.create_ebs_csi_irsa_role || var.enable_irsa
    error_message = "create_ebs_csi_irsa_role requires enable_irsa to be true when the EBS CSI add-on is enabled."
  }
}
