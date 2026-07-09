output "topic_arns" {
  description = "SNS topic ARNs keyed by topic name."
  value       = { for name, topic in aws_sns_topic.this : name => topic.arn }
}
