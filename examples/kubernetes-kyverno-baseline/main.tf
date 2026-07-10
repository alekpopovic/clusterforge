module "kyverno" {
  source = "../../modules/platform/kubernetes/kyverno"

  chart_version            = var.kyverno_chart_version
  enable_baseline_policies = var.enable_baseline_policies
  baseline_failure_action  = "Audit"

  values = [yamlencode({
    admissionController = {
      replicas = 1
    }
    backgroundController = {
      replicas = 1
    }
  })]
}
