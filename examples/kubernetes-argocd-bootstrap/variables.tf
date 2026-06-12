variable "kubeconfig_path" {
  description = "Path to a local kubeconfig for an existing Kubernetes cluster."
  type        = string
  default     = "~/.kube/config"
}

variable "namespace" {
  description = "Namespace where Argo CD is installed."
  type        = string
  default     = "argocd"

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "namespace must not be empty."
  }
}

variable "labels" {
  description = "Labels applied to the Argo CD namespace and app-of-apps manifest."
  type        = map(string)
  default = {
    "clusterforge.io/example" = "kubernetes-argocd-bootstrap"
  }
}

variable "enable_app_of_apps" {
  description = "Whether to create the app-of-apps Argo CD Application. Enable after reviewing the GitOps repo URL and path."
  type        = bool
  default     = false
}

variable "gitops_repo_url" {
  description = "Git repository URL containing app-of-apps definitions. Do not include credentials."
  type        = string
  default     = "https://github.com/example/platform-gitops.git"

  validation {
    condition     = !var.enable_app_of_apps || length(trimspace(var.gitops_repo_url)) > 0
    error_message = "gitops_repo_url must not be empty when enable_app_of_apps is true."
  }
}

variable "gitops_path" {
  description = "Path in the GitOps repository for app-of-apps definitions."
  type        = string
  default     = "gitops/apps"

  validation {
    condition     = length(trimspace(var.gitops_path)) > 0
    error_message = "gitops_path must not be empty."
  }
}
