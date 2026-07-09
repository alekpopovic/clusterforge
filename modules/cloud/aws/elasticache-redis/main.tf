locals {
  name        = trimspace(var.name)
  environment = trimspace(var.environment)

  tags = merge(var.tags, {
    Name        = local.name
    Environment = local.environment
  })
}

resource "aws_elasticache_subnet_group" "this" {
  name       = "${local.name}-redis"
  subnet_ids = var.subnet_ids
  tags       = local.tags
}

resource "aws_security_group" "this" {
  name        = "${local.name}-redis"
  description = "Redis access for ${local.name}."
  vpc_id      = var.vpc_id
  tags        = local.tags
}

resource "aws_vpc_security_group_ingress_rule" "redis" {
  for_each = toset(var.allowed_security_group_ids)

  security_group_id            = aws_security_group.this.id
  referenced_security_group_id = each.value
  ip_protocol                  = "tcp"
  from_port                    = 6379
  to_port                      = 6379
  description                  = "Allow Redis from ${each.value}."
}

resource "aws_vpc_security_group_egress_rule" "redis" {
  security_group_id = aws_security_group.this.id
  cidr_ipv4         = "0.0.0.0/0"
  ip_protocol       = "tcp"
  from_port         = 6379
  to_port           = 6379
  description       = "Allow Redis return traffic."
}

resource "aws_elasticache_replication_group" "this" {
  replication_group_id = local.name
  description          = "Redis cache for ${local.name}."

  engine         = "redis"
  engine_version = var.engine_version
  node_type      = var.node_type

  num_cache_clusters         = var.num_cache_nodes
  automatic_failover_enabled = var.automatic_failover_enabled
  multi_az_enabled           = var.multi_az_enabled

  subnet_group_name  = aws_elasticache_subnet_group.this.name
  security_group_ids = [aws_security_group.this.id]
  port               = 6379

  at_rest_encryption_enabled = var.at_rest_encryption_enabled
  transit_encryption_enabled = var.transit_encryption_enabled
  parameter_group_name       = var.parameter_group_name == "" ? null : var.parameter_group_name

  tags = local.tags

  lifecycle {
    precondition {
      condition     = length(var.subnet_ids) > 0
      error_message = "subnet_ids must not be empty."
    }

    precondition {
      condition     = length(var.allowed_security_group_ids) > 0
      error_message = "allowed_security_group_ids must not be empty."
    }

    precondition {
      condition     = !var.automatic_failover_enabled || var.num_cache_nodes >= 2
      error_message = "automatic_failover_enabled requires num_cache_nodes >= 2."
    }

    precondition {
      condition     = var.auth_token_secret_arn == "" || var.transit_encryption_enabled
      error_message = "auth_token_secret_arn should only be used when transit_encryption_enabled is true."
    }
  }
}
