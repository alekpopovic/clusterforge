variable "namespace" {
  description = "Kubernetes namespace for the argo-cd Helm release."
  type        = string
  default     = "argocd"

  validation {
    condition     = length(trimspace(var.namespace)) > 0
    error_message = "Namespace must not be empty."
  }
}

variable "chart_version" {
  description = "Optional argo-cd chart version. Leave empty to use the latest provider-resolved version."
  type        = string
  default     = ""
}

variable "values" {
  description = "YAML values passed to the Helm release."
  type        = list(string)
  default     = []
}

variable "labels" {
  description = "Labels applied to the namespace when create_namespace is true."
  type        = map(string)
  default     = {}
}

variable "create_namespace" {
  description = "Whether to create the namespace before installing the Helm release."
  type        = bool
  default     = true
}

variable "enable_app_of_apps" {
  description = "Whether to create a bootstrap Argo CD Application that points at an app-of-apps repository path."
  type        = bool
  default     = false
}

variable "app_of_apps_name" {
  description = "Name of the app-of-apps Argo CD Application."
  type        = string
  default     = "cluster-apps"

  validation {
    condition     = length(trimspace(var.app_of_apps_name)) > 0
    error_message = "App-of-apps name must not be empty."
  }
}

variable "app_of_apps_repo_url" {
  description = "Git repository URL containing the app-of-apps definitions. Do not include credentials."
  type        = string
  default     = ""

  validation {
    condition     = !var.enable_app_of_apps || length(trimspace(var.app_of_apps_repo_url)) > 0
    error_message = "app_of_apps_repo_url must not be empty when enable_app_of_apps is true."
  }
}

variable "app_of_apps_path" {
  description = "Path within the Git repository for the app-of-apps definitions."
  type        = string
  default     = "apps"

  validation {
    condition     = length(trimspace(var.app_of_apps_path)) > 0
    error_message = "App-of-apps path must not be empty."
  }
}

variable "app_of_apps_revision" {
  description = "Git revision for the app-of-apps Application."
  type        = string
  default     = "HEAD"

  validation {
    condition     = length(trimspace(var.app_of_apps_revision)) > 0
    error_message = "App-of-apps revision must not be empty."
  }
}

variable "app_of_apps_destination_namespace" {
  description = "Destination namespace used by the app-of-apps Application."
  type        = string
  default     = "argocd"

  validation {
    condition     = length(trimspace(var.app_of_apps_destination_namespace)) > 0
    error_message = "App-of-apps destination namespace must not be empty."
  }
}

variable "app_of_apps_project" {
  description = "Argo CD project for the app-of-apps Application."
  type        = string
  default     = "default"

  validation {
    condition     = length(trimspace(var.app_of_apps_project)) > 0
    error_message = "App-of-apps project must not be empty."
  }
}
