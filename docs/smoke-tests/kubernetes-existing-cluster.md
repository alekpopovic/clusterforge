---
title: Existing Kubernetes Cluster Smoke Test
permalink: /smoke-tests/kubernetes-existing-cluster/
---

# Existing Kubernetes Cluster Smoke Test

This runbook verifies ClusterForge Kubernetes platform and workload modules
against an existing disposable Kubernetes cluster. It must be run manually and
is not part of default CI because it can modify a real cluster.

Do not use a production cluster. Do not commit kubeconfigs, tokens, generated
state files, plan files, or real endpoint details.

## Required Tools

- `terraform` or `tofu`
- `kubectl`
- `helm`
- ClusterForge `cf` CLI, when using generated environments

## Required Permissions

Use a temporary Kubernetes identity with permissions to create, update, and
delete test resources:

- namespaces
- deployments, services, config maps, service accounts, and RBAC resources
- ingress resources, when testing ingress
- Helm-managed resources for bootstrap components
- cluster-scoped resources only when the selected module requires them

## Test Inputs

| Field | Value |
| --- | --- |
| Tester | `<name or handle>` |
| Date | `<YYYY-MM-DD>` |
| Cluster alias | `<redacted cluster alias>` |
| Kubernetes version | `<version>` |
| Terraform/OpenTofu version | `<version>` |
| ClusterForge version or commit | `<version or commit SHA>` |
| Kubeconfig context | `<context name>` |
| Evidence folder | `<local path outside git or ignored path>` |

## Procedure

1. Confirm the current Kubernetes context points to the disposable test cluster:

   ```bash
   kubectl config current-context
   kubectl cluster-info
   kubectl version
   kubectl auth can-i create namespaces
   ```

2. Create an isolated namespace:

   ```bash
   export KUBE_SMOKE_NS="clusterforge-smoke"
   kubectl create namespace "${KUBE_SMOKE_NS}"
   ```

3. Run a provider connectivity check from a Kubernetes example root:

   ```bash
   cd examples/kubernetes-basic-app
   terraform init
   terraform validate
   terraform plan \
     -out .cf-kubernetes-basic-app.tfplan \
     -var="namespace=${KUBE_SMOKE_NS}"
   terraform apply .cf-kubernetes-basic-app.tfplan
   ```

   Use `tofu` instead of `terraform` when testing OpenTofu.

4. Verify the demo workload:

   ```bash
   kubectl -n "${KUBE_SMOKE_NS}" get deployments
   kubectl -n "${KUBE_SMOKE_NS}" get services
   kubectl -n "${KUBE_SMOKE_NS}" get pods -o wide
   ```

5. Install platform bootstrap components for the selected scope:

   ```bash
   helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
   helm repo update
   helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx \
     --namespace ingress-nginx \
     --create-namespace
   kubectl -n ingress-nginx rollout status deployment/ingress-nginx-controller
   ```

   If testing ClusterForge platform modules directly, use the relevant module
   root and save a plan before applying. Review any cluster-scoped resources
   before approval.

6. Optional ingress check:

   ```bash
   kubectl -n ingress-nginx get service ingress-nginx-controller
   ```

   Record ingress status. Record an app endpoint only when the test creates one
   and it is safe to share in redacted form.

7. Destroy Terraform-managed resources:

   ```bash
   terraform destroy
   ```

8. Delete bootstrap resources and namespace:

   ```bash
   helm uninstall ingress-nginx -n ingress-nginx || true
   kubectl delete namespace ingress-nginx --ignore-not-found
   kubectl delete namespace "${KUBE_SMOKE_NS}" --ignore-not-found
   ```

9. Continue with [cleanup](./cleanup.md) before marking the run complete.

## Expected Outputs

- Kubernetes context alias
- Kubernetes version
- workload deployment status
- service status
- ingress status, if tested
- app endpoint, if available

## Evidence Placeholders

| Evidence | Location |
| --- | --- |
| Tool versions | `<path>` |
| `kubectl get nodes` | `<path>` |
| workload status | `<path>` |
| ingress status | `<path or not tested>` |
| Terraform/OpenTofu plan summary | `<path>` |
| cleanup confirmation | `<path>` |

## Result Record

| Field | Value |
| --- | --- |
| Result | `not run`, `passed`, `failed`, or `blocked` |
| Start time | `<timestamp>` |
| End time | `<timestamp>` |
| Cluster alias | `<redacted alias>` |
| Kubernetes version | `<version>` |
| Ingress status | `<status or not tested>` |
| App endpoint | `<redacted endpoint or none>` |
| Evidence path | `<path>` |
| Cleanup confirmed by | `<name or handle>` |
| Notes | `<findings and follow-ups>` |
