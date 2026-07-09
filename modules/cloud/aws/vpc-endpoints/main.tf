data "aws_region" "current" {}

locals {
  name        = trimspace(var.name)
  environment = trimspace(var.environment)

  common_tags = merge(var.tags, {
    Name        = local.name
    Environment = local.environment
  })

  gateway_endpoints   = toset(var.gateway_endpoints)
  interface_endpoints = toset(var.interface_endpoints)

  create_interface_security_group = var.create_security_group && length(local.interface_endpoints) > 0
  interface_security_group_ids    = local.create_interface_security_group ? [aws_security_group.interface_endpoints[0].id] : var.security_group_ids
}

resource "aws_security_group" "interface_endpoints" {
  count = local.create_interface_security_group ? 1 : 0

  name        = "${local.name}-vpc-endpoints"
  description = "Interface VPC endpoint access for ${local.name}."
  vpc_id      = var.vpc_id
  tags = merge(local.common_tags, {
    Name = "${local.name}-vpc-endpoints"
  })
}

resource "aws_vpc_security_group_ingress_rule" "interface_https" {
  for_each = local.create_interface_security_group ? toset(var.allowed_security_group_ids) : toset([])

  security_group_id            = aws_security_group.interface_endpoints[0].id
  referenced_security_group_id = each.value
  ip_protocol                  = "tcp"
  from_port                    = 443
  to_port                      = 443
  description                  = "Allow HTTPS from ${each.value}."
}

resource "aws_vpc_security_group_egress_rule" "interface_https" {
  count = local.create_interface_security_group ? 1 : 0

  security_group_id = aws_security_group.interface_endpoints[0].id
  cidr_ipv4         = "0.0.0.0/0"
  ip_protocol       = "tcp"
  from_port         = 443
  to_port           = 443
  description       = "Allow endpoint return traffic."
}

resource "aws_vpc_endpoint" "gateway" {
  for_each = local.gateway_endpoints

  vpc_id            = var.vpc_id
  service_name      = "com.amazonaws.${data.aws_region.current.region}.${each.value}"
  vpc_endpoint_type = "Gateway"
  route_table_ids   = var.route_table_ids
  tags = merge(local.common_tags, {
    Name    = "${local.name}-${replace(each.value, ".", "-")}"
    Service = each.value
  })
}

resource "aws_vpc_endpoint" "interface" {
  for_each = local.interface_endpoints

  vpc_id              = var.vpc_id
  service_name        = "com.amazonaws.${data.aws_region.current.region}.${each.value}"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = var.subnet_ids
  security_group_ids  = local.interface_security_group_ids
  private_dns_enabled = var.private_dns_enabled
  tags = merge(local.common_tags, {
    Name    = "${local.name}-${replace(each.value, ".", "-")}"
    Service = each.value
  })

  lifecycle {
    precondition {
      condition     = length(var.subnet_ids) > 0
      error_message = "subnet_ids must not be empty when interface_endpoints are configured."
    }

    precondition {
      condition     = var.create_security_group || length(var.security_group_ids) > 0
      error_message = "security_group_ids must not be empty when create_security_group is false and interface_endpoints are configured."
    }
  }
}
