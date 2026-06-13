module "podinfo" {
  source = "../../modules/workloads/kubernetes/helm-app"

  name       = "podinfo"
  namespace  = var.namespace
  repository = "https://stefanprodan.github.io/podinfo"
  chart      = "podinfo"

  chart_version = "6.7.1"

  values = [
    yamlencode({
      replicaCount = 2
      service = {
        type = "ClusterIP"
      }
    })
  ]

  labels = {
    "clusterforge.io/example" = "kubernetes-helm-app"
  }
}
