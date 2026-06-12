variable "kubeconfig_path" {
  description = "Path to a local kubeconfig for an existing Kubernetes cluster."
  type        = string
  default     = "~/.kube/config"
}

variable "gitops_repo_url" {
  description = "Git repository URL containing app-of-apps definitions. Do not include credentials."
  type        = string
  default     = "https://github.com/example/platform-gitops.git"
}

variable "gitops_path" {
  description = "Path in the GitOps repository for app-of-apps definitions."
  type        = string
  default     = "gitops/apps"
}

variable "gitops_revision" {
  description = "Git revision for app-of-apps definitions."
  type        = string
  default     = "main"
}
