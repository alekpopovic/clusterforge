provider "aws" {
  region = var.aws_region

  access_key = var.use_fake_credentials_for_plan ? "clusterforge-fake-access-key" : null
  secret_key = var.use_fake_credentials_for_plan ? "clusterforge-fake-secret-key" : null

  skip_credentials_validation = var.use_fake_credentials_for_plan
  skip_metadata_api_check     = var.use_fake_credentials_for_plan
  skip_requesting_account_id  = var.use_fake_credentials_for_plan
}

provider "kubernetes" {
  host                   = module.eks.cluster_endpoint
  cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
  token                  = data.aws_eks_cluster_auth.this.token
}

provider "helm" {
  kubernetes = {
    host                   = module.eks.cluster_endpoint
    cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
    token                  = data.aws_eks_cluster_auth.this.token
  }
}
