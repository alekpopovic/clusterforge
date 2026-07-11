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
module "collector" {
  source        = "../../modules/platform/kubernetes/opentelemetry-collector"
  chart_version = var.chart_version
  mode          = "deployment"
  presets = {
    kubernetesAttributes = { enabled = true }
  }
}
