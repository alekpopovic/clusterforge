---
title: AWS EKS Smoke Test
permalink: /smoke-tests/aws-eks/
---

# AWS EKS Smoke Test

This runbook describes a manual smoke test against a real AWS account. It is
intentionally not part of default CI because it creates billable cloud
resources.

Do not use production accounts, shared state, real credentials in files, or
real account IDs in test notes. Record evidence with redacted identifiers.

## Cost Warning

This test can cost money. EKS control planes, NAT Gateway, load balancers,
managed node groups, EBS volumes, and CloudWatch logs can all create charges.
Destroy the environment as soon as evidence is collected.

## Required Tools

- `terraform` or `tofu`
- AWS CLI
- `kubectl`
- `helm`
- ClusterForge `cf` CLI

## Required AWS Permissions

Use a temporary role or disposable test account with permissions for:

- VPC
- IAM
- EKS
- EC2
- CloudWatch
- Route53, optional when testing DNS integrations

## Test Inputs

Fill these values before starting:

| Field | Value |
| --- | --- |
| Tester | `<name or handle>` |
| Date | `<YYYY-MM-DD>` |
| AWS account alias | `<redacted test account alias>` |
| AWS region | `<region>` |
| Terraform/OpenTofu version | `<version>` |
| ClusterForge version or commit | `<version or commit SHA>` |
| Kubernetes version | `<version>` |
| Test environment name | `<smoke-eks-YYYYMMDD-initials>` |
| Evidence folder | `<local path outside git or ignored path>` |

## Preflight

1. Confirm the active AWS identity is a disposable test identity:

   ```bash
   aws sts get-caller-identity
   ```

   Save only redacted output in evidence.

2. Confirm tool versions:

   ```bash
   terraform version || tofu version
   aws --version
   kubectl version --client=true
   helm version
   cf version
   ```

3. Create an isolated test workspace from a clean checkout:

   ```bash
   export CF_ENV="smoke-eks-YYYYMMDD-initials"
   export AWS_REGION="us-east-2"
   export TF_ENGINE="terraform"
   mkdir -p .cf/evidence/${CF_ENV}
   ```

   When testing OpenTofu through `cf`, add `--engine tofu` to each `cf init`,
   `cf plan`, `cf apply`, and `cf destroy` command.

## Procedure

1. Initialize a ClusterForge project if the checkout does not already have one:

   ```bash
   cf project init clusterforge-smoke
   ```

2. Create an isolated EKS environment:

   ```bash
   cf env create "${CF_ENV}" \
     --cloud aws \
     --orchestrator eks \
     --region "${AWS_REGION}"
   ```

3. Generate readable Terraform files:

   ```bash
   cf generate "${CF_ENV}"
   ```

4. Configure `terraform.tfvars` without secrets:

   ```bash
   cd "live/${CF_ENV}/aws-eks"
   cp terraform.tfvars.example terraform.tfvars
   ```

   Review and edit these values:

   ```hcl
   region      = "<region>"
   project     = "clusterforge-smoke"
   environment = "<smoke environment>"
   name        = "<unique cluster name>"

   enable_nat_gateway = true
   single_nat_gateway = true

   kubernetes_version = "<supported EKS version>"

   default_node_instance_types = ["t3.medium"]
   default_node_min_size       = 1
   default_node_desired_size   = 2
   default_node_max_size       = 2
   ```

   Keep `terraform.tfvars`, plans, state, and kubeconfigs out of git.

5. Return to the repository root and initialize Terraform/OpenTofu:

   ```bash
   cd -
   cf init "${CF_ENV}"
   ```

6. Create and review a saved plan:

   ```bash
   cf plan "${CF_ENV}" \
     --out .cf/plans/${CF_ENV}.tfplan \
     --risk-summary
   ```

   Review the plan for expected VPC, IAM, EKS, EC2, and CloudWatch resources.
   Save the plan summary in the evidence folder.

7. Apply using the reviewed plan file:

   ```bash
   cf apply "${CF_ENV}" --plan-file .cf/plans/${CF_ENV}.tfplan
   ```

8. Capture expected Terraform outputs:

   ```bash
   ${TF_ENGINE} -chdir="live/${CF_ENV}/aws-eks" output
   ```

   Expected outputs include:

   - cluster name
   - kubeconfig command
   - ingress status, after platform bootstrap is installed
   - app endpoint, if a load balancer or ingress address is available

9. Verify the cluster exists:

   ```bash
   aws eks describe-cluster \
     --region "${AWS_REGION}" \
     --name "<cluster-name>" \
     --query 'cluster.{name:name,status:status,version:version,endpoint:endpoint}'
   ```

10. Verify the managed node group is ready:

    ```bash
    aws eks list-nodegroups \
      --region "${AWS_REGION}" \
      --cluster-name "<cluster-name>"

    aws eks describe-nodegroup \
      --region "${AWS_REGION}" \
      --cluster-name "<cluster-name>" \
      --nodegroup-name "<nodegroup-name>" \
      --query 'nodegroup.{name:nodegroupName,status:status,desired:scalingConfig.desiredSize}'
    ```

11. Configure and verify `kubectl` access:

    ```bash
    aws eks update-kubeconfig \
      --region "${AWS_REGION}" \
      --name "<cluster-name>" \
      --alias "${CF_ENV}"

    kubectl config use-context "${CF_ENV}"
    kubectl get nodes -o wide
    kubectl get pods -A
    ```

12. Install platform bootstrap:

    If the generated environment includes the optional platform bootstrap
    module, uncomment and configure it, then run another saved plan and apply.
    Otherwise install the minimum bootstrap components manually for the smoke
    test:

    ```bash
    helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
    helm repo update
    helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx \
      --namespace ingress-nginx \
      --create-namespace

    kubectl -n ingress-nginx rollout status deployment/ingress-nginx-controller
    kubectl -n ingress-nginx get service ingress-nginx-controller
    ```

13. Deploy a demo app:

    ```bash
    kubectl create namespace clusterforge-smoke
    kubectl -n clusterforge-smoke create deployment hello \
      --image=nginx:1.27
    kubectl -n clusterforge-smoke expose deployment hello \
      --port 80 \
      --target-port 80
    kubectl -n clusterforge-smoke rollout status deployment/hello
    kubectl -n clusterforge-smoke get service hello
    ```

    If ingress is enabled for the test, create a temporary ingress with a
    placeholder host under a domain you control. Do not record real hosted zone
    IDs in the repository.

14. Record evidence:

    - redacted `aws eks describe-cluster` output
    - redacted node group status
    - `kubectl get nodes`
    - `kubectl get pods -A`
    - ingress controller service status
    - app service or endpoint status
    - screenshots or logs with account IDs and external identifiers redacted

15. Destroy the test environment:

    ```bash
    kubectl delete namespace clusterforge-smoke --ignore-not-found
    helm uninstall ingress-nginx -n ingress-nginx || true
    kubectl delete namespace ingress-nginx --ignore-not-found
    cf destroy "${CF_ENV}" --allow-destroy
    ```

    Continue with [cleanup](./cleanup.md) before marking the run complete.

## Cleanup Checklist

- Delete the demo app namespace.
- Delete platform bootstrap Helm releases and namespaces.
- Delete the EKS cluster and managed node groups through Terraform/OpenTofu.
- Delete the VPC through Terraform/OpenTofu.
- Check for stuck AWS Load Balancers, target groups, security groups, ENIs,
  EBS volumes, NAT gateways, Elastic IPs, and CloudWatch log groups.
- Remove temporary kubeconfig contexts.
- Remove local plan files, state backups, and evidence files that should not be
  retained.

## Result Record

| Field | Value |
| --- | --- |
| Result | `not run`, `passed`, `failed`, or `blocked` |
| Start time | `<timestamp>` |
| End time | `<timestamp>` |
| Cluster name | `<redacted name>` |
| Kubeconfig command | `<redacted command>` |
| Ingress status | `<status>` |
| App endpoint | `<redacted endpoint or none>` |
| Evidence path | `<path>` |
| Cleanup confirmed by | `<name or handle>` |
| Notes | `<findings and follow-ups>` |
