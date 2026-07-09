## Prompt 100 — AWS messaging modules: SQS and SNS

```text
Implement AWS messaging modules for container workloads.

Create modules:
- modules/cloud/aws/sqs
- modules/cloud/aws/sns

SQS module inputs:
- queues: map(object({
    fifo_queue = optional(bool, false)
    visibility_timeout_seconds = optional(number, 30)
    message_retention_seconds = optional(number, 345600)
    delay_seconds = optional(number, 0)
    max_message_size = optional(number, 262144)
    receive_wait_time_seconds = optional(number, 0)
    dead_letter_queue = optional(bool, false)
    max_receive_count = optional(number, 5)
  }))
- tags map(string)

SQS resources:
- aws_sqs_queue
- optional DLQ
- redrive policy

Outputs:
- queue_urls
- queue_arns
- dead_letter_queue_urls
- dead_letter_queue_arns

SNS module inputs:
- topics: map(object({
    fifo_topic = optional(bool, false)
    content_based_deduplication = optional(bool, false)
  }))
- subscriptions optional
- tags

Outputs:
- topic_arns

Docs:
- Queue worker example.
- ECS worker example.
- EKS worker with IRSA example.
- DLQ explanation.
- IAM permissions guidance.

Examples:
- examples/aws-sqs-worker
- examples/aws-sns-topic

Rules:
- Do not create overly broad IAM by default.
- Do not include real ARNs.
- Keep modules composable.

Run:
- terraform fmt -recursive
```

---
