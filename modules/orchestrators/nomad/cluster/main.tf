locals {
  server_config = jsonencode({ datacenter = var.datacenter, data_dir = var.data_dir, bind_addr = var.bind_addr, server = true, bootstrap_expect = var.server_count })
  client_config = jsonencode({ datacenter = var.datacenter, data_dir = var.data_dir, bind_addr = var.bind_addr, client = { enabled = true, servers = var.server_addresses } })
  cloud_init    = <<-EOT
    #cloud-config
    packages: ["nomad"]
    write_files:
      - path: /etc/nomad.d/nomad.json
        permissions: "0640"
        content: |
          ${indent(10, local.client_config)}
    runcmd:
      - [systemctl, enable, --now, nomad]
  EOT
}
