# aws-sqs-worker

Example SQS queue with a dead-letter queue for worker workloads.

Run local validation:

```bash
terraform init
terraform validate
terraform plan -refresh=false
```

Attach narrowly scoped IAM permissions to EKS IRSA roles or ECS task roles.
