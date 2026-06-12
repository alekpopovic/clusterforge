# Kubernetes Argo CD Bootstrap Example

Installs Argo CD into an existing Kubernetes cluster. The app-of-apps
Application is opt-in so the default example does not point a cluster at a
placeholder GitOps repository.

This example assumes the Kubernetes and Helm providers can use a local
kubeconfig. It does not include Git credentials.

## Usage

```bash
terraform init
terraform validate
terraform plan
```

Override the GitOps repository for your own environment:

```bash
terraform plan \
  -var='enable_app_of_apps=true' \
  -var='gitops_repo_url=https://github.com/example/platform-gitops.git' \
  -var='gitops_path=gitops/apps'
```

For private repositories, configure Argo CD repository credentials through an
approved secret management path. Do not put Git credentials in Terraform
variables or `tfvars`.
