---
title: AWS ECS Smoke Test
permalink: /smoke-tests/aws-ecs/
---

# AWS ECS Smoke Test

This runbook describes a manual smoke test against a real AWS account for the
ClusterForge AWS ECS path. It is not part of default CI because it creates
billable cloud resources.

Do not put real account IDs, credentials, state files, plan files, or private
identifiers in the repository.

## Cost Warning

This test can cost money. NAT Gateway, load balancers, ECS service resources,
CloudWatch logs, and related networking resources can all create charges.
Destroy the environment after evidence is collected.

## Required Tools

- `terraform` or `tofu`
- AWS CLI
- ClusterForge `cf` CLI

Optional tools for workload checks:

- `curl`
- `jq`

## Required AWS Permissions

Use a disposable test identity with permissions for:

- VPC
- IAM, when testing workload execution roles
- ECS
- EC2
- Elastic Load Balancing, when testing ALB-backed services
- CloudWatch
- Route53, optional when testing DNS integrations

## Test Inputs

| Field | Value |
| --- | --- |
| Tester | `<name or handle>` |
| Date | `<YYYY-MM-DD>` |
| AWS account alias | `<redacted test account alias>` |
| AWS region | `<region>` |
| Terraform/OpenTofu version | `<version>` |
| ClusterForge version or commit | `<version or commit SHA>` |
| ECS launch type | `<FARGATE or other>` |
| Test environment name | `<smoke-ecs-YYYYMMDD-initials>` |
| Evidence folder | `<local path outside git or ignored path>` |

## Procedure

1. Confirm the active AWS identity is a disposable test identity:

   ```bash
   aws sts get-caller-identity
   ```

   Save only redacted output in evidence.

2. Set test variables:

   ```bash
   export CF_ENV="smoke-ecs-YYYYMMDD-initials"
   export AWS_REGION="us-east-2"
   export TF_ENGINE="terraform"
   mkdir -p .cf/evidence/${CF_ENV}
   ```

   When testing OpenTofu through `cf`, add `--engine tofu` to each `cf init`,
   `cf plan`, `cf apply`, and `cf destroy` command.

3. Initialize a ClusterForge project if needed:

   ```bash
   cf project init clusterforge-smoke
   ```

4. Create an isolated ECS environment:

   ```bash
   cf env create "${CF_ENV}" \
     --cloud aws \
     --orchestrator ecs \
     --region "${AWS_REGION}"
   ```

5. Generate Terraform files:

   ```bash
   cf generate "${CF_ENV}"
   ```

6. Configure `terraform.tfvars` without secrets:

   ```bash
   cd "live/${CF_ENV}/aws-ecs"
   cp terraform.tfvars.example terraform.tfvars
   ```

   Review and edit:

   ```hcl
   region      = "<region>"
   project     = "clusterforge-smoke"
   environment = "<smoke environment>"
   name        = "<unique ecs cluster name>"

   enable_nat_gateway        = true
   single_nat_gateway        = true
   enable_container_insights = true
   capacity_providers        = ["FARGATE", "FARGATE_SPOT"]
   ```

7. Initialize, plan, and apply with a plan file:

   ```bash
   cd -
   cf init "${CF_ENV}"
   cf plan "${CF_ENV}" \
     --out .cf/plans/${CF_ENV}.tfplan \
     --risk-summary
   cf apply "${CF_ENV}" --plan-file .cf/plans/${CF_ENV}.tfplan
   ```

8. Verify the ECS cluster exists:

   ```bash
   aws ecs describe-clusters \
     --region "${AWS_REGION}" \
     --clusters "<cluster-name>" \
     --query 'clusters[].{name:clusterName,status:status,registeredContainerInstances:registeredContainerInstancesCount,runningTasks:runningTasksCount}'
   ```

9. Verify networking resources were created:

   ```bash
   ${TF_ENGINE} -chdir="live/${CF_ENV}/aws-ecs" output
   aws ec2 describe-vpcs \
     --region "${AWS_REGION}" \
     --filters "Name=tag:Environment,Values=${CF_ENV}" \
     --query 'Vpcs[].{VpcId:VpcId,State:State,CidrBlock:CidrBlock}'
   ```

10. Optional workload smoke test:

    Use `examples/ecs-fargate-app` or `examples/ecs-fargate-app-with-alb` as a
    reviewed workload root. Configure it to use the smoke VPC, subnets, cluster
    name, and temporary task image. Apply only after reviewing the plan.

    Expected workload evidence:

    - ECS service status reaches steady state
    - task starts successfully
    - CloudWatch log stream is created
    - ALB target health is healthy, when ALB is enabled
    - endpoint responds, when an endpoint is created

11. Record evidence:

    - redacted `aws ecs describe-clusters` output
    - Terraform/OpenTofu output
    - ECS service status, if a workload was deployed
    - ALB DNS name or endpoint status, redacted when needed
    - CloudWatch log group status

12. Destroy the test environment:

    Destroy optional workload roots first, then the generated ECS environment:

    ```bash
    cf destroy "${CF_ENV}" --allow-destroy
    ```

    Continue with [cleanup](./cleanup.md) before marking the run complete.

## Expected Outputs

- ECS cluster name
- VPC ID, redacted in shared evidence
- subnet IDs, redacted in shared evidence
- workload service status, if deployed
- app endpoint, if an ALB-backed workload is deployed

## Result Record

| Field | Value |
| --- | --- |
| Result | `not run`, `passed`, `failed`, or `blocked` |
| Start time | `<timestamp>` |
| End time | `<timestamp>` |
| ECS cluster name | `<redacted name>` |
| Workload status | `<status or not tested>` |
| App endpoint | `<redacted endpoint or none>` |
| Evidence path | `<path>` |
| Cleanup confirmed by | `<name or handle>` |
| Notes | `<findings and follow-ups>` |
