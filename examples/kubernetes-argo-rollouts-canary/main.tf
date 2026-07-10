module "argo_rollouts" {
  source = "../../modules/platform/kubernetes/argo-rollouts"

  chart_version = var.argo_rollouts_chart_version
}

module "demo" {
  count  = var.enable_rollout_app ? 1 : 0
  source = "../../modules/workloads/kubernetes/rollout-app"

  name             = "rollouts-demo"
  namespace        = "rollouts-demo"
  create_namespace = true
  image            = "quay.io/argoproj/rollouts-demo:blue"
  replicas         = 3

  resources = {
    requests = { cpu = "50m", memory = "64Mi" }
    limits   = { cpu = "200m", memory = "128Mi" }
  }

  strategy = {
    type = "canary"
    canary_steps = [
      { set_weight = 20 },
      { pause = true },
      { set_weight = 50 },
      { pause_duration = "30s" },
      { set_weight = 100 }
    ]
  }

  depends_on = [module.argo_rollouts]
}
