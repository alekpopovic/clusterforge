variable "topics" {
  description = "SNS topics keyed by topic name."
  type = map(object({
    fifo_topic                  = optional(bool, false)
    content_based_deduplication = optional(bool, false)
  }))
}

variable "subscriptions" {
  description = "SNS subscriptions keyed by subscription name."
  type = map(object({
    topic                = string
    protocol             = string
    endpoint             = string
    raw_message_delivery = optional(bool, false)
  }))
  default = {}
}

variable "tags" {
  description = "Tags applied to SNS topics."
  type        = map(string)
  default     = {}
}
