# aws-sns-topic

Example SNS topic for application events.

Run local validation:

```bash
terraform init
terraform validate
terraform plan -refresh=false
```

Grant publishers `sns:Publish` only to the specific topic ARN.
