output "install_command" {
  description = "K3s server install command for review or manual use."
  value       = local.install_command
}

output "server_user_data" {
  description = "Cloud-init user data for a K3s server."
  value       = <<-EOT
  #cloud-config
  runcmd:
    - ${local.install_command}
  EOT
}

output "agent_user_data" {
  description = "Cloud-init template for K3s agents. Replace server URL and token out of band."
  value       = <<-EOT
  #cloud-config
  runcmd:
    - curl -sfL https://get.k3s.io | K3S_URL=https://<server>:6443 K3S_TOKEN=<token> sh -
  EOT
}

output "kubeconfig_retrieval_notes" {
  description = "Instructions for retrieving kubeconfig without storing it in state."
  value       = "Retrieve /etc/rancher/k3s/k3s.yaml from the first server through your approved secure access path and rewrite the server address."
}
