locals {
  alias_name = startswith(var.alias_name, "alias/") ? var.alias_name : "alias/${var.alias_name}"

  tags = merge(
    {
      Name        = var.name
      Environment = var.environment
      ManagedBy   = "terraform"
    },
    var.tags
  )
}

resource "aws_kms_key" "this" {
  #checkov:skip=CKV2_AWS_64:Encryption key selection is configurable and may use an approved external or provider-managed key.
  description             = var.description
  deletion_window_in_days = var.deletion_window_in_days
  enable_key_rotation     = var.enable_key_rotation
  policy                  = trimspace(var.policy_json) == "" ? null : var.policy_json
  tags                    = local.tags
}

resource "aws_kms_alias" "this" {
  name          = local.alias_name
  target_key_id = aws_kms_key.this.key_id
}
