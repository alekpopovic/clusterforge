output "alb_arn" {
  description = "ARN of the Application Load Balancer."
  value       = aws_lb.this.arn
}

output "alb_dns_name" {
  description = "DNS name of the Application Load Balancer."
  value       = aws_lb.this.dns_name
}

output "alb_zone_id" {
  description = "Canonical hosted zone ID of the Application Load Balancer."
  value       = aws_lb.this.zone_id
}

output "security_group_id" {
  description = "Security group ID attached to the Application Load Balancer."
  value       = aws_security_group.this.id
}

output "target_group_arns" {
  description = "Target group ARNs keyed by target group name."
  value = {
    for key, target_group in aws_lb_target_group.this : key => target_group.arn
  }
}

output "listener_arns" {
  description = "Listener ARNs keyed by listener purpose."
  value = merge(
    length(aws_lb_listener.http_forward) > 0 ? { http = aws_lb_listener.http_forward[0].arn } : {},
    length(aws_lb_listener.http_redirect) > 0 ? { http_redirect = aws_lb_listener.http_redirect[0].arn } : {},
    length(aws_lb_listener.https) > 0 ? { https = aws_lb_listener.https[0].arn } : {}
  )
}
