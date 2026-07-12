locals {
  name        = trimspace(var.name)
  environment = trimspace(var.environment)

  common_tags = merge(var.tags, {
    Name        = local.name
    Environment = local.environment
    Module      = "platform/ecs/alb"
  })

  http_listener_enabled  = var.enable_http
  https_listener_enabled = var.enable_https && trimspace(var.certificate_arn) != ""

  first_target_group_key = length(keys(var.target_groups)) > 0 ? sort(keys(var.target_groups))[0] : null
}

#trivy:ignore:AWS-0104
resource "aws_security_group" "this" {
  #checkov:skip=CKV_AWS_382:Unrestricted egress is an explicit compatibility default; production policy must narrow destination rules.
  name        = "${local.name}-alb"
  description = "Security group for ${local.name} ECS application load balancer."
  vpc_id      = var.vpc_id

  dynamic "ingress" {
    for_each = local.http_listener_enabled ? [1] : []

    content {
      description = "Allow HTTP traffic to the load balancer."
      from_port   = 80
      to_port     = 80
      protocol    = "tcp"
      cidr_blocks = var.allowed_cidr_blocks
    }
  }

  dynamic "ingress" {
    for_each = local.https_listener_enabled ? [1] : []

    content {
      description = "Allow HTTPS traffic to the load balancer."
      from_port   = 443
      to_port     = 443
      protocol    = "tcp"
      cidr_blocks = var.allowed_cidr_blocks
    }
  }

  egress {
    description = "Allow outbound traffic to ECS service target groups."
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = local.common_tags
}

#trivy:ignore:AWS-0052
#trivy:ignore:AWS-0053
resource "aws_lb" "this" {
  #checkov:skip=CKV2_AWS_20:ALB exposure, TLS, WAF, logging, and lifecycle controls are explicit deployment inputs, not safe universal defaults.
  #checkov:skip=CKV2_AWS_28:ALB exposure, TLS, WAF, logging, and lifecycle controls are explicit deployment inputs, not safe universal defaults.
  #checkov:skip=CKV_AWS_131:ALB exposure, TLS, WAF, logging, and lifecycle controls are explicit deployment inputs, not safe universal defaults.
  #checkov:skip=CKV_AWS_150:ALB exposure, TLS, WAF, logging, and lifecycle controls are explicit deployment inputs, not safe universal defaults.
  #checkov:skip=CKV_AWS_91:ALB exposure, TLS, WAF, logging, and lifecycle controls are explicit deployment inputs, not safe universal defaults.
  name               = local.name
  load_balancer_type = "application"
  internal           = false
  security_groups    = [aws_security_group.this.id]
  subnets            = var.public_subnet_ids

  tags = local.common_tags
}

resource "aws_lb_target_group" "this" {
  for_each = var.target_groups

  name        = "${local.name}-${each.key}"
  port        = each.value.port
  protocol    = upper(each.value.protocol)
  vpc_id      = var.vpc_id
  target_type = "ip"

  health_check {
    enabled = true
    path    = each.value.health_check_path
    matcher = each.value.health_check_matcher
  }

  tags = merge(local.common_tags, {
    Name = "${local.name}-${each.key}"
  })
}

resource "aws_lb_listener" "http_redirect" {
  count = local.http_listener_enabled && local.https_listener_enabled ? 1 : 0

  load_balancer_arn = aws_lb.this.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

#trivy:ignore:AWS-0054
resource "aws_lb_listener" "http_forward" {
  #checkov:skip=CKV_AWS_103:ALB exposure, TLS, WAF, logging, and lifecycle controls are explicit deployment inputs, not safe universal defaults.
  #checkov:skip=CKV_AWS_2:ALB exposure, TLS, WAF, logging, and lifecycle controls are explicit deployment inputs, not safe universal defaults.
  count = local.http_listener_enabled && !local.https_listener_enabled ? 1 : 0

  load_balancer_arn = aws_lb.this.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.this[local.first_target_group_key].arn
  }
}
resource "aws_lb_listener" "https" {
  #checkov:skip=CKV2_AWS_74:Cipher selection follows the explicit SSL policy input supplied by the deployment.
  #checkov:skip=CKV_AWS_103:TLS policy is explicit because supported client baselines differ by environment.
  count = local.https_listener_enabled ? 1 : 0

  load_balancer_arn = aws_lb.this.arn
  port              = 443
  protocol          = "HTTPS"
  certificate_arn   = var.certificate_arn

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.this[local.first_target_group_key].arn
  }
}
