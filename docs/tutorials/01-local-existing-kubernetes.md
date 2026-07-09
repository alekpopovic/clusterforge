# Tutorial 01: Local Existing Kubernetes

## Goal

Deploy a simple app to a disposable namespace on an existing Kubernetes cluster.

## Prerequisites

- `terraform` or `tofu`
- `kubectl`
- kubeconfig for a non-production cluster

## Commands

```bash
kubectl config current-context
kubectl create namespace clusterforge-demo
cd examples/kubernetes-basic-app
terraform init
terraform plan -out demo.tfplan -var='namespace=clusterforge-demo'
terraform apply demo.tfplan
```

## Generated Files

No ClusterForge files are generated in this tutorial.

## What Terraform Creates

- namespace, if configured
- deployment
- service

## Validate

```bash
kubectl -n clusterforge-demo get deployments,services,pods
```

## Cleanup

```bash
terraform destroy -var='namespace=clusterforge-demo'
kubectl delete namespace clusterforge-demo --ignore-not-found
```

## Troubleshooting

Check the active context, RBAC permissions, image pull errors, and provider
configuration. Do not use a production cluster.
