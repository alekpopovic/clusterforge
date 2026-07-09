locals {
  zone_arns = [
    for zone_id in var.hosted_zone_ids : "arn:aws:route53:::hostedzone/${zone_id}"
  ]
}

data "aws_iam_policy_document" "route53_dns01" {
  statement {
    sid = "ReadRoute53"

    actions = [
      "route53:GetChange",
      "route53:ListHostedZonesByName",
      "route53:ListResourceRecordSets",
    ]

    resources = ["*"]
  }

  statement {
    sid = "ChangeDns01Records"

    actions = [
      "route53:ChangeResourceRecordSets",
    ]

    resources = var.allow_all_hosted_zones ? ["arn:aws:route53:::hostedzone/*"] : local.zone_arns

    condition {
      test     = "ForAllValues:StringLike"
      variable = "route53:ChangeResourceRecordSetsNormalizedRecordNames"
      values   = ["_acme-challenge.*"]
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
    cert-manager-route53-dns01 = data.aws_iam_policy_document.route53_dns01.json
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
