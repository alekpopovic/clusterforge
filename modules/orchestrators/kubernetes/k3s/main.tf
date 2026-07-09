locals {
  version_env = var.k3s_version == "" ? "" : "INSTALL_K3S_VERSION=${var.k3s_version} "
  server_args = compact(concat(
    [for item in var.disable_components : "--disable ${item}"],
    var.cluster_cidr == null ? [] : ["--cluster-cidr ${var.cluster_cidr}"],
    var.service_cidr == null ? [] : ["--service-cidr ${var.service_cidr}"],
    [for san in var.tls_san : "--tls-san ${san}"]
  ))
  install_command = "${local.version_env}curl -sfL https://get.k3s.io | INSTALL_K3S_CHANNEL=${var.install_channel} sh -s - server ${join(" ", local.server_args)}"
}
