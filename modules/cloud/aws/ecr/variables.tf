variable "repositories" {
  description = "ECR repositories keyed by repository name."
  type = map(object({
    image_tag_mutability  = optional(string, "IMMUTABLE")
    scan_on_push          = optional(bool, true)
    encryption_type       = optional(string, "AES256")
    kms_key_arn           = optional(string)
    lifecycle_policy_json = optional(string)
  }))

  validation {
    condition     = length(var.repositories) > 0
    error_message = "repositories must not be empty."
  }

  validation {
    condition = alltrue([
      for repository in values(var.repositories) :
      contains(["MUTABLE", "IMMUTABLE"], repository.image_tag_mutability)
    ])
    error_message = "image_tag_mutability must be MUTABLE or IMMUTABLE."
  }

  validation {
    condition = alltrue([
      for repository in values(var.repositories) :
      contains(["AES256", "KMS"], repository.encryption_type)
    ])
    error_message = "encryption_type must be AES256 or KMS."
  }
}

variable "tags" {
  description = "Tags applied to ECR repositories."
  type        = map(string)
  default     = {}
}
