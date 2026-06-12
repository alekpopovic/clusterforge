variable "create_zone" {
  description = "Whether to create a Route53 hosted zone for zone_name."
  type        = bool
  default     = false
}

variable "zone_name" {
  description = "Route53 hosted zone name, such as example.com."
  type        = string

  validation {
    condition     = length(trimspace(var.zone_name)) > 0
    error_message = "zone_name must not be empty."
  }
}

variable "zone_id" {
  description = "Existing Route53 hosted zone ID. When empty and create_zone is false, the module looks up zone_name."
  type        = string
  default     = ""
}

variable "records" {
  description = "DNS records to create in the selected hosted zone."
  type = map(object({
    name    = string
    type    = string
    ttl     = optional(number)
    records = optional(list(string))
    alias = optional(object({
      name                   = string
      zone_id                = string
      evaluate_target_health = optional(bool, true)
    }))
  }))
  default = {}

  validation {
    condition = alltrue([
      for record in values(var.records) : length(trimspace(record.name)) > 0
    ])
    error_message = "Each record name must not be empty."
  }

  validation {
    condition = alltrue([
      for record in values(var.records) : contains([
        "A", "AAAA", "CAA", "CNAME", "MX", "NS", "PTR", "SOA", "SPF", "SRV", "TXT"
      ], upper(record.type))
    ])
    error_message = "Each record type must be a supported Route53 record type."
  }

  validation {
    condition = alltrue([
      for record in values(var.records) : (
        (record.records != null && record.alias == null) ||
        (record.records == null && record.alias != null)
      )
    ])
    error_message = "Each record must set exactly one of records or alias."
  }

  validation {
    condition = alltrue([
      for record in values(var.records) : record.alias == null || record.ttl == null
    ])
    error_message = "Alias records must not set ttl."
  }

  validation {
    condition = alltrue([
      for record in values(var.records) : record.alias != null || try(record.ttl > 0, false)
    ])
    error_message = "Non-alias records must set ttl greater than 0."
  }
}

variable "tags" {
  description = "Tags applied to created Route53 hosted zones."
  type        = map(string)
  default     = {}
}
