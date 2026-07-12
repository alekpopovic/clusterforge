locals {
  name        = trimspace(var.name)
  environment = trimspace(var.environment)

  tags = merge(var.tags, {
    Name        = local.name
    Environment = local.environment
  })
}

resource "aws_db_subnet_group" "this" {
  name       = "${local.name}-postgres"
  subnet_ids = var.subnet_ids
  tags       = local.tags
}

resource "aws_security_group" "this" {
  name        = "${local.name}-postgres"
  description = "PostgreSQL access for ${local.name}."
  vpc_id      = var.vpc_id
  tags        = local.tags
}

resource "aws_vpc_security_group_ingress_rule" "postgres" {
  for_each = toset(var.allowed_security_group_ids)

  security_group_id            = aws_security_group.this.id
  referenced_security_group_id = each.value
  ip_protocol                  = "tcp"
  from_port                    = 5432
  to_port                      = 5432
  description                  = "Allow PostgreSQL from ${each.value}."
}

#trivy:ignore:AWS-0104
resource "aws_vpc_security_group_egress_rule" "postgres" {
  security_group_id = aws_security_group.this.id
  cidr_ipv4         = "0.0.0.0/0"
  ip_protocol       = "tcp"
  from_port         = 5432
  to_port           = 5432
  description       = "Allow PostgreSQL return traffic."
}

resource "aws_db_instance" "this" {
  #checkov:skip=CKV2_AWS_30:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV2_AWS_60:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV_AWS_118:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV_AWS_129:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV_AWS_157:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV_AWS_161:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV_AWS_226:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  #checkov:skip=CKV_AWS_353:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  identifier = local.name

  engine         = "postgres"
  engine_version = var.engine_version
  instance_class = var.instance_class

  allocated_storage     = var.allocated_storage
  max_allocated_storage = var.max_allocated_storage == 0 ? null : var.max_allocated_storage
  storage_encrypted     = var.storage_encrypted
  kms_key_id            = var.kms_key_arn == "" ? null : var.kms_key_arn

  db_name  = var.database_name
  username = var.master_username
  password = var.manage_master_user_password ? null : var.master_password

  manage_master_user_password   = var.manage_master_user_password
  master_user_secret_kms_key_id = var.manage_master_user_password && var.kms_key_arn != "" ? var.kms_key_arn : null

  db_subnet_group_name   = aws_db_subnet_group.this.name
  vpc_security_group_ids = [aws_security_group.this.id]
  publicly_accessible    = false

  port                      = 5432
  multi_az                  = var.multi_az
  backup_retention_period   = var.backup_retention_period
  deletion_protection       = var.deletion_protection
  skip_final_snapshot       = var.skip_final_snapshot
  final_snapshot_identifier = var.skip_final_snapshot ? null : "${local.name}-final"

  tags = local.tags

  lifecycle {
    precondition {
      condition     = length(var.subnet_ids) >= 2
      error_message = "subnet_ids should include at least two private subnets for RDS."
    }

    precondition {
      condition     = length(var.allowed_security_group_ids) > 0
      error_message = "allowed_security_group_ids must not be empty."
    }

    precondition {
      condition     = var.manage_master_user_password || length(var.master_password) > 0
      error_message = "master_password is required when manage_master_user_password is false."
    }
  }
}
