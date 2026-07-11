provider "kubernetes" {}
provider "helm" {
  kubernetes {}
}

module "kyverno" {
  source = "../../modules/platform/kubernetes/kyverno"

  chart_version                       = var.kyverno_chart_version
  enable_baseline_policies            = true
  baseline_failure_action             = var.enforce ? "Enforce" : "Audit"
  enable_pod_security_extended_policy = true
  allowed_registries                  = var.allowed_registries
  require_image_digest                = var.require_image_digest
}
