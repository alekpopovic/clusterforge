output "queue_urls" {
  description = "SQS queue URLs."
  value       = module.queues.queue_urls
}

output "dead_letter_queue_arns" {
  description = "Dead-letter queue ARNs."
  value       = module.queues.dead_letter_queue_arns
}
