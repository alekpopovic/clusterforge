# cloud/aws/sqs

## Purpose

Creates SQS queues for container workers, with optional dead-letter queues.
This module does not create IAM roles by default; grant producers and consumers
only the actions they need.

## Status

Implemented.

## Queue Worker Example

```hcl
module "queues" {
  source = "../../../modules/cloud/aws/sqs"

  queues = {
    jobs = {
      dead_letter_queue = true
      max_receive_count = 5
    }
  }
}
```

## EKS Worker With IRSA

Create an IRSA role for the worker service account and attach a policy scoped
to the queue ARN. Do not grant wildcard queue access unless it is explicitly
reviewed.

## ECS Worker

Attach scoped `sqs:ReceiveMessage`, `sqs:DeleteMessage`, and
`sqs:GetQueueAttributes` permissions to the ECS task role for consumers.

## DLQ Notes

Dead-letter queues capture messages that fail repeatedly. Monitor DLQ depth and
create a replay procedure before production use.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
