locals {
  zone_arns = [
    for zone_id in var.hosted_zone_ids : "arn:aws:route53:::hostedzone/${zone_id}"
  ]

  change_actions = var.policy_mode == "sync" ? [
    "route53:ChangeResourceRecordSets",
    ] : [
    "route53:ChangeResourceRecordSets",
  ]
}

data "aws_iam_policy_document" "external_dns" {
  statement {
    sid = "ListHostedZones"

    actions = [
      "route53:ListHostedZones",
      "route53:ListHostedZonesByName",
      "route53:ListResourceRecordSets",
    ]

    resources = ["*"]
  }

  statement {
    sid       = var.policy_mode == "sync" ? "ChangeSelectedHostedZones" : "UpsertSelectedHostedZones"
    actions   = local.change_actions
    resources = var.allow_all_hosted_zones ? ["arn:aws:route53:::hostedzone/*"] : local.zone_arns

    condition {
      test     = "ForAllValues:StringLike"
      variable = "route53:ChangeResourceRecordSetsActions"
      values   = var.policy_mode == "sync" ? ["CREATE", "UPSERT", "DELETE"] : ["CREATE", "UPSERT"]
    }
  }
}

module "irsa" {
  source = "../irsa-role"

  name                 = var.name
  environment          = var.environment
  oidc_provider_arn    = var.oidc_provider_arn
  oidc_provider_url    = var.oidc_provider_url
  namespace            = var.namespace
  service_account_name = var.service_account_name
  tags                 = var.tags

  inline_policies = {
    external-dns-route53 = data.aws_iam_policy_document.external_dns.json
  }
}

resource "terraform_data" "hosted_zone_guard" {
  lifecycle {
    precondition {
      condition     = var.allow_all_hosted_zones || length(var.hosted_zone_ids) > 0
      error_message = "hosted_zone_ids must not be empty unless allow_all_hosted_zones is true."
    }
  }
}
