output "install_command" {
  description = "RKE2 server install command for review or manual use."
  value       = local.install_command
}

output "server_user_data" {
  description = "Cloud-init user data for an RKE2 server."
  value       = <<-EOT
  #cloud-config
  write_files:
    - path: /etc/rancher/rke2/config.yaml
      permissions: "0600"
      content: |
  ${indent(6, join("\n", local.config_lines))}
  runcmd:
    - ${local.install_command}
    - systemctl enable rke2-server
    - systemctl start rke2-server
  EOT
}

output "agent_user_data" {
  description = "Cloud-init template for RKE2 agents. Replace server URL and token out of band."
  value       = <<-EOT
  #cloud-config
  write_files:
    - path: /etc/rancher/rke2/config.yaml
      permissions: "0600"
      content: |
        server: https://<server>:9345
        token: <token>
  runcmd:
    - curl -sfL https://get.rke2.io | INSTALL_RKE2_TYPE="agent" sh -
    - systemctl enable rke2-agent
    - systemctl start rke2-agent
  EOT
}

output "kubeconfig_retrieval_notes" {
  description = "Instructions for retrieving kubeconfig without storing it in state."
  value       = "Retrieve /etc/rancher/rke2/rke2.yaml from the first server through your approved secure access path and rewrite the server address."
}
