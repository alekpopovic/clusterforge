locals {
  version_env = var.rke2_version == "" ? "" : "INSTALL_RKE2_VERSION=${var.rke2_version} "
  config_lines = compact(concat(
    ["write-kubeconfig-mode: \"0640\""],
    var.cluster_cidr == null ? [] : ["cluster-cidr: ${var.cluster_cidr}"],
    var.service_cidr == null ? [] : ["service-cidr: ${var.service_cidr}"],
    [for item in var.disable_components : "disable: ${item}"],
    [for san in var.tls_san : "tls-san: ${san}"]
  ))
  install_command = "${local.version_env}curl -sfL https://get.rke2.io | INSTALL_RKE2_CHANNEL=${var.install_channel} sh -"
}
