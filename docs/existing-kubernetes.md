# Existing Kubernetes

Use this target when a Kubernetes cluster already exists and ClusterForge should
manage only platform add-ons and workloads.

## Create Environment

```bash
cf env create dev --cloud existing --orchestrator kubernetes --region local
cf generate dev
```

The generated root is:

```text
live/dev/existing-kubernetes/
  versions.tf
  providers.tf
  main.tf
  variables.tf
  outputs.tf
  terraform.tfvars.example
  README.md
```

## Provider Configuration

The generated Kubernetes and Helm providers use the same kubeconfig settings:

```hcl
kubeconfig_path    = "~/.kube/config"
kubeconfig_context = "replace-with-context"
```

Do not commit kubeconfigs, tokens, generated state, or local `terraform.tfvars`.

## Platform Bootstrap

The generated `main.tf` includes a commented bootstrap module. Enable only the
components that are appropriate for the target cluster. Cloud-specific DNS,
load balancer, and identity behavior should be reviewed before applying.

## Workloads

Render app manifests into the environment:

```bash
cf app add hello --image nginx:1.27 --port 80
cf app validate hello
cf app render hello --env dev
```

## Limitations

- ClusterForge does not create the cluster.
- In-cluster Terraform execution is a future option.
- Provider permissions depend on the selected kubeconfig context.
- No secrets are generated.
