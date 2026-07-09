# cloud/aws/sns

## Purpose

Creates SNS topics and optional subscriptions for container workloads. This
module does not create broad publisher or subscriber IAM permissions.

## Status

Implemented.

## Topic Example

```hcl
module "topics" {
  source = "../../../modules/cloud/aws/sns"

  topics = {
    events = {}
  }
}
```

## Subscription Example

```hcl
module "topics" {
  source = "../../../modules/cloud/aws/sns"

  topics = {
    events = {}
  }

  subscriptions = {
    audit_queue = {
      topic    = "events"
      protocol = "sqs"
      endpoint = module.queues.queue_arns["audit"]
    }
  }
}
```

Grant publishers `sns:Publish` only on the topic ARN they need.

## Generated Terraform Documentation

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->
