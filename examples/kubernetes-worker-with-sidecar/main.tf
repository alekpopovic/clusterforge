# Sidecar support is planned for the generic worker module. Use helm-app,
# Kustomize, or GitOps for complex multi-container pods today.
module "worker" {
  source = "../../modules/workloads/kubernetes/worker"

  name      = "worker"
  namespace = "clusterforge-worker"
  image     = "busybox:1.36"
  command   = ["sh", "-c"]
  args      = ["while true; do date; sleep 30; done"]
}
