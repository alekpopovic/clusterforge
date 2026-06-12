locals {
  zone_name = trimsuffix(trimspace(var.zone_name), ".")

  effective_zone_id = var.create_zone ? aws_route53_zone.this[0].zone_id : (
    trimspace(var.zone_id) != "" ? trimspace(var.zone_id) : data.aws_route53_zone.this[0].zone_id
  )

  effective_zone_name = var.create_zone ? aws_route53_zone.this[0].name : (
    trimspace(var.zone_id) != "" ? local.zone_name : data.aws_route53_zone.this[0].name
  )
}

data "aws_route53_zone" "this" {
  count = !var.create_zone && trimspace(var.zone_id) == "" ? 1 : 0

  name         = local.zone_name
  private_zone = false
}

resource "aws_route53_zone" "this" {
  count = var.create_zone ? 1 : 0

  name = local.zone_name
  tags = var.tags
}

resource "aws_route53_record" "this" {
  for_each = var.records

  zone_id = local.effective_zone_id
  name    = each.value.name
  type    = upper(each.value.type)
  ttl     = each.value.alias == null ? each.value.ttl : null
  records = each.value.alias == null ? each.value.records : null

  dynamic "alias" {
    for_each = each.value.alias == null ? [] : [each.value.alias]

    content {
      name                   = alias.value.name
      zone_id                = alias.value.zone_id
      evaluate_target_health = alias.value.evaluate_target_health
    }
  }
}
