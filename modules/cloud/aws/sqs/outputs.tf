output "queue_urls" {
  description = "SQS queue URLs keyed by queue name."
  value       = { for name, queue in aws_sqs_queue.this : name => queue.url }
}

output "queue_arns" {
  description = "SQS queue ARNs keyed by queue name."
  value       = { for name, queue in aws_sqs_queue.this : name => queue.arn }
}

output "dead_letter_queue_urls" {
  description = "Dead-letter queue URLs keyed by source queue name."
  value       = { for name, queue in aws_sqs_queue.dlq : name => queue.url }
}

output "dead_letter_queue_arns" {
  description = "Dead-letter queue ARNs keyed by source queue name."
  value       = { for name, queue in aws_sqs_queue.dlq : name => queue.arn }
}
