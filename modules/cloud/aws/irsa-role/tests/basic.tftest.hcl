provider "aws" {
  region                      = "eu-central-1"
  access_key                  = "test"
  secret_key                  = "test"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}

run "valid_irsa_role_trust_policy" {
  command = plan

  variables {
    name                 = "clusterforge-dev-ebs-csi"
    environment          = "dev"
    oidc_provider_arn    = "arn:aws:iam::123456789012:oidc-provider/oidc.eks.eu-central-1.amazonaws.com/id/EXAMPLE"
    oidc_provider_url    = "https://oidc.eks.eu-central-1.amazonaws.com/id/EXAMPLE"
    namespace            = "kube-system"
    service_account_name = "ebs-csi-controller-sa"
    policy_arns          = ["arn:aws:iam::aws:policy/service-role/AmazonEBSCSIDriverPolicy"]
    inline_policies = {
      example = jsonencode({
        Version = "2012-10-17"
        Statement = [{
          Effect   = "Allow"
          Action   = "sts:GetCallerIdentity"
          Resource = "*"
        }]
      })
    }
  }

  assert {
    condition     = aws_iam_role.this.name == "clusterforge-dev-ebs-csi"
    error_message = "Expected IAM role name to match input."
  }

  assert {
    condition     = length(aws_iam_role_policy_attachment.this) == 1
    error_message = "Expected managed policy attachment to be planned."
  }

  assert {
    condition     = length(aws_iam_role_policy.this) == 1
    error_message = "Expected inline policy to be planned."
  }

  assert {
    condition     = strcontains(data.aws_iam_policy_document.assume_role.json, "system:serviceaccount:kube-system:ebs-csi-controller-sa")
    error_message = "Expected trust policy to include Kubernetes service account subject."
  }

  assert {
    condition     = strcontains(data.aws_iam_policy_document.assume_role.json, "oidc.eks.eu-central-1.amazonaws.com/id/EXAMPLE:aud")
    error_message = "Expected trust policy to include OIDC audience condition."
  }

  assert {
    condition     = output.role_name == "clusterforge-dev-ebs-csi"
    error_message = "Expected role_name output to match role name."
  }
}

run "empty_service_account_name_fails" {
  command = plan

  variables {
    name                 = "clusterforge-dev-ebs-csi"
    environment          = "dev"
    oidc_provider_arn    = "arn:aws:iam::123456789012:oidc-provider/oidc.eks.eu-central-1.amazonaws.com/id/EXAMPLE"
    oidc_provider_url    = "https://oidc.eks.eu-central-1.amazonaws.com/id/EXAMPLE"
    namespace            = "kube-system"
    service_account_name = ""
  }

  expect_failures = [var.service_account_name]
}
