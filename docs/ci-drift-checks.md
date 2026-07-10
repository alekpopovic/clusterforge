# Scheduled drift checks in GitHub Actions

ClusterForge includes disabled workflow templates:

- `.github/workflows/examples/drift-check-aws-eks.yml`
- `.github/workflows/examples/drift-check-aws-ecs.yml`

Files below `.github/workflows/examples/` are not discovered as runnable
workflows. Copy the required template directly into `.github/workflows/`, update
its environment name and schedule, review every action version, and merge it
through normal code review to enable it.

## AWS authentication with GitHub OIDC

Create a dedicated, read-only IAM role trusted by the repository's GitHub OIDC
subject. Store its ARN as the repository or environment variable
`AWS_DRIFT_ROLE_ARN`, and store the region as `AWS_REGION`. The workflow grants
`id-token: write` only so GitHub can mint a short-lived token; no long-lived AWS
access keys are required.

Restrict the role trust policy to the exact repository, branch or GitHub
environment. Restrict AWS permissions to Terraform refresh/read operations and
backend access. S3 state and lock-table permissions still need careful review
because Terraform state can contain sensitive data.

## Schedule and outcome

The examples run on weekday cron schedules and support `workflow_dispatch`.
GitHub cron uses UTC. Offset EKS and ECS jobs to reduce API and state-lock
contention, and tune frequency to operational need and cloud API cost.

`cf drift check` uses Terraform's detailed exit code: zero means no drift and two
means drift. The templates upload only `drift-summary.json`, never the plan or
state file. With `FAIL_ON_DRIFT=false`, drift produces a workflow warning. Change
it to `true` to fail the job after the artifact is uploaded.

Operational errors such as authentication, backend, initialization, or provider
failures fail the job regardless of drift policy. Investigate them separately
from actual resource drift.

## Why there is no automatic apply

Drift may be intentional, emergency-created, or evidence that configuration is
wrong. Automatically applying after a scheduled check can delete or replace
resources without review. These templates contain no apply, destroy, or
remediation command. Review the summary, create a normal plan, obtain approval,
and use the existing guarded apply workflow.

The templates use GitHub OIDC as documented by the official
[AWS credentials action](https://github.com/aws-actions/configure-aws-credentials)
and upload summaries using the official
[artifact action](https://github.com/actions/upload-artifact).
