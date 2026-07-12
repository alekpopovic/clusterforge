resource "aws_ecr_repository" "this" {
  #checkov:skip=CKV_AWS_51:The reported control is an explicit reusable-module input or requires operator-owned integration resources.
  for_each = var.repositories

  name                 = each.key
  image_tag_mutability = each.value.image_tag_mutability
  tags                 = var.tags

  image_scanning_configuration {
    scan_on_push = each.value.scan_on_push
  }

  encryption_configuration {
    encryption_type = each.value.encryption_type
    kms_key         = each.value.encryption_type == "KMS" ? each.value.kms_key_arn : null
  }

  lifecycle {
    precondition {
      condition     = each.value.encryption_type != "KMS" || try(length(trimspace(each.value.kms_key_arn)) > 0, false)
      error_message = "kms_key_arn is required when encryption_type is KMS."
    }
  }
}

resource "aws_ecr_lifecycle_policy" "this" {
  for_each = {
    for name, repository in var.repositories : name => repository
    if try(length(trimspace(repository.lifecycle_policy_json)) > 0, false)
  }

  repository = aws_ecr_repository.this[each.key].name
  policy     = each.value.lifecycle_policy_json
}
