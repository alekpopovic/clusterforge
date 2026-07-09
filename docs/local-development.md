# Local Development

ClusterForge can use local Kubernetes clusters for workload development without
cloud infrastructure.

Supported local targets:

- kind
- k3d

## Prerequisites

Install one local cluster tool yourself. ClusterForge does not vendor kind or
k3d.

```bash
kind version
k3d version
```

## Create A Cluster

```bash
cf local create kind
cf local status
cf local kubeconfig kind
```

For k3d:

```bash
cf local create k3d
cf local status
cf local kubeconfig k3d
```

The create command also writes a local environment under `live/local/kind` or
`live/local/k3d`. These roots configure the Kubernetes provider from kubeconfig
and do not install cloud-specific modules.

## Deploy An App

```bash
cf app add hello --image nginx:1.27 --port 80
cf app validate hello
cf app render hello --env local
```

Review the generated Terraform before applying from the local environment root.

## Cleanup

```bash
cf local delete kind
cf local delete k3d
```

Only the named `clusterforge-local` local cluster is deleted.

## Limitations

- Local clusters do not match EKS, AKS, or GKE IAM, load balancer, DNS, or
  storage behavior.
- Platform modules that assume cloud load balancers or DNS should be tested
  carefully.
- ClusterForge does not modify global kubeconfig except through explicit
  kind/k3d create/delete operations.
