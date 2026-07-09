variable "queues" {
  description = "SQS queues keyed by queue name."
  type = map(object({
    fifo_queue                 = optional(bool, false)
    visibility_timeout_seconds = optional(number, 30)
    message_retention_seconds  = optional(number, 345600)
    delay_seconds              = optional(number, 0)
    max_message_size           = optional(number, 262144)
    receive_wait_time_seconds  = optional(number, 0)
    dead_letter_queue          = optional(bool, false)
    max_receive_count          = optional(number, 5)
  }))

  validation {
    condition     = length(var.queues) > 0
    error_message = "queues must not be empty."
  }
}

variable "tags" {
  description = "Tags applied to SQS queues."
  type        = map(string)
  default     = {}
}
