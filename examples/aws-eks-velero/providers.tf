provider "aws" {
  region = var.aws_region

  access_key = var.use_fake_credentials_for_plan ? "clusterforge-fake-access-key" : null
  secret_key = var.use_fake_credentials_for_plan ? "clusterforge-fake-secret-key" : null

  skip_credentials_validation = var.use_fake_credentials_for_plan
  skip_metadata_api_check     = var.use_fake_credentials_for_plan
  skip_requesting_account_id  = var.use_fake_credentials_for_plan
}

provider "kubernetes" {
  config_path    = var.kubeconfig_path
  config_context = var.kubeconfig_context
}

provider "helm" {
  kubernetes = {
    config_path    = var.kubeconfig_path
    config_context = var.kubeconfig_context
  }
}
