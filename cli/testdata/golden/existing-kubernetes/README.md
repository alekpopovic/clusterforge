# dev existing Kubernetes

Generated ClusterForge environment for an existing Kubernetes cluster.

This root configures Kubernetes and Helm providers from kubeconfig. It does not
create cloud networking, IAM, or a Kubernetes control plane.

Never commit kubeconfig files, tokens, secrets, state, or local tfvars.
